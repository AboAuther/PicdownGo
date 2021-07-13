package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

var baseDir string

func urlFormat(picUrl, protocol, website string) (string, error) {
	u, err := url.Parse(picUrl)
	if err != nil {
		return "", fmt.Errorf("picture's Url is not correct,%w", err)
	}
	if u.Scheme == "" {
		picUrl = protocol + ":" + picUrl
	} else if u.Scheme == protocol {
		return picUrl, nil
	} else if u.Host == "" && u.Scheme == "" {
		picUrl = protocol + ":" + website + picUrl
	}
	return picUrl, nil
}
func downloading(destDir, picUrl string, targetWeb *Website) (err error) {
	picurl, err := urlFormat(picUrl, targetWeb.TCPProtocol, targetWeb.Website)
	if err != nil {
		return fmt.Errorf("parse picture's url failed,%w", err)
	}
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(picurl)
	if err != nil {
		return fmt.Errorf("picUrl:%s Get picture URL failed,%w", picurl, err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Status code error:%d %s \n", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()
	out, err := os.Create(destDir + "/" + fmt.Sprint(rand.Int()) + ".png")
	if err != nil {
		return fmt.Errorf("create file failed,%w", err)
	}
	defer func() {
		err := out.Close()
		if err != nil {
			return
		}
	}()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("down picture failed,%w", err)
	}
	return nil
}
func worker(destDir string, targetWeb *Website, linkChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for picUrl := range linkChan {
		err := downloading(destDir, picUrl, targetWeb)
		if err != nil {
			log.Printf("%s", err)
		}
	}
}
func findDomainByUrl(postUrl string, configuration *Parser) (*Website, error) {
	var targetDomain string
	u, err := url.Parse(postUrl)
	if err != nil {
		return nil, fmt.Errorf("URL may be wrong,%w", err)
	}
	targetDomain = u.Host
	for index, website := range configuration.SupportWebsites {
		if strings.Contains(targetDomain, website.Website) {
			fmt.Println("use ", configuration.SupportWebsites[index].Website, "parser")
			return &configuration.SupportWebsites[index], nil
		}
	}
	err = errors.New("url input is not in support list")
	return nil, err
}

func crawler(postUrl string, workNum int, jsonfile *Parser) (err error) {
	client := http.Client{Timeout: 5 * time.Second}
	res, err := client.Get(postUrl)
	if err != nil {
		return fmt.Errorf("url input is unvalued,%w", err)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Status code error:%d %s \n", res.StatusCode, res.Status)
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return fmt.Errorf("create new documents failed,%w", err)
	}
	targetSiteConf, err := findDomainByUrl(postUrl, jsonfile)
	if err != nil {
		return fmt.Errorf("find domian failed,%w", err)
	}
	if targetSiteConf == nil {
		return nil
	}
	title := doc.Find(targetSiteConf.TitlePattern).Text()
	fmt.Println("[", targetSiteConf.Website, "]:", title, "start downloading...")
	dir := fmt.Sprintf("%v/%v - %v", baseDir, targetSiteConf.Website, title)
	_ = os.MkdirAll(dir, 0755)

	linkChan := make(chan string)
	wg := new(sync.WaitGroup)

	for i := 0; i < workNum; i++ {
		wg.Add(1)
		go worker(dir, targetSiteConf, linkChan, wg)
	}
	doc.Find(targetSiteConf.ImgPattern).Each(func(i int, img *goquery.Selection) {
		imgUrl, _ := img.Attr(targetSiteConf.ImgAddrPattern)
		linkChan <- imgUrl
	})
	close(linkChan)
	wg.Wait()
	return nil
}

func checkFileIsExits(jsonFileAddr string) bool {
	if _, err := os.Stat(jsonFileAddr); err == nil {
		return true
	}
	return false
}
func reloadParser(jsonFileAddr string) (*Parser, error) {
	var file []byte
	var configuration Parser
	if checkFileIsExits(jsonFileAddr) {
		file, _ = os.ReadFile(jsonFileAddr)
		fmt.Println("Loading local json file...")
	} else {
		err := errors.New("json file is not Existed")
		return nil, err
	}
	if len(file) == 0 {
		fmt.Println("Json file is empty,Loading default parser...")
		file = DefaultJson
	}
	err := json.Unmarshal(file, &configuration)
	if err != nil {
		return nil, fmt.Errorf("json file unmarshal failed,%w", err)
	}
	return &configuration, nil
}
func main() {
	log.SetFlags(log.Lshortfile)
	var workNum int
	var postUrl string
	var command = &cobra.Command{
		Use:   "PicdownGo -u URL",
		Short: "Download pictures",
		Long:  `Download given url pictures by cmdline`,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if postUrl == "" {
				return fmt.Errorf("please use 'PicdownGo -u URL'")
			}
			baseDir = "Pic"
			err = os.MkdirAll(baseDir, 0755)
			if err != nil {
				return fmt.Errorf("create dictionary failed,%w", err)
			}
			jsonFileAddr := fmt.Sprintf("./parser.json")
			jsonfile, err := reloadParser(jsonFileAddr)
			if err != nil {
				return fmt.Errorf("loading jsonfile failed,%w", err)
			}
			err = crawler(postUrl, workNum, jsonfile)
			if err != nil {
				return fmt.Errorf("%w", err)
			}
			fmt.Println("pictures have downloaded")
			return nil

		}}
	command.Flags().StringVarP(&postUrl, "URL", "u", "", "URL of post")
	command.Flags().IntVarP(&workNum, "workerNum", "w", 20, "number of workers")
	err := command.Execute()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Fprintf:%v\n", err.Error())
		os.Exit(1)
	}

}
