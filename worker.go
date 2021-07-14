package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
)

func urlFormat(picUrl, protocol, website string) (string, error) {
	u, err := url.Parse(picUrl)
	if err != nil {
		return "", fmt.Errorf("picture's Url is not correct,%w", err)
	}
	if u.Scheme == "" && u.Host != "" {
		picUrl = protocol + ":" + picUrl
	} else if u.Scheme == protocol {
		return picUrl, nil
	} else if u.Host == "" && u.Scheme == "" {
		picUrl = protocol + "://" + website + picUrl
	}
	return picUrl, nil
}

func downloading(destDir, picUrl string, targetWeb *Website, client http.Client) (err error) {
	picurl, err := urlFormat(picUrl, targetWeb.TCPProtocol, targetWeb.Website)
	if err != nil {
		return fmt.Errorf("parse picture's url failed,%w", err)
	}
	resp, err := client.Get(picurl)
	if err != nil {
		return fmt.Errorf("picUrl:%s Get picture URL failed,%w", picurl, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Status code error:%d %s \n", resp.StatusCode, resp.Status)
	}
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
