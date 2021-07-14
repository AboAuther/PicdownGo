package main

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

func worker(destDir string, targetWeb *Website, linkChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	client := http.Client{Timeout: timeOut}
	for picUrl := range linkChan {
		err := downloading(destDir, picUrl, targetWeb, client)
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
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Status code error:%d %s \n", res.StatusCode, res.Status)
	}
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
