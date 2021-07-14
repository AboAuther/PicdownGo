package main

import (
	"testing"
)

func TestUrlFormat(t *testing.T) {
	t.Run("TCP:https and '//'", func(t *testing.T) {
		picurl := "//inews.gtimg.com/newsapp_ls/0/13753802147_640330/0"
		website := Website{
			Website:        "qq.com",
			TitlePattern:   "title",
			ImgPattern:     ".pic img",
			ImgAddrPattern: "src",
			TCPProtocol:    "https",
		}
		got, err := urlFormat(picurl, website.TCPProtocol, website.Website)
		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}
		want := "https://inews.gtimg.com/newsapp_ls/0/13753802147_640330/0"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("TCP:http and '//'", func(t *testing.T) {
		picurl := "//inews.gtimg.com/newsapp_ls/0/13753802147_640330/0"
		website := Website{
			"qq.com",
			"title",
			".pic img",
			"src",
			"http",
		}
		got, err := urlFormat(picurl, website.TCPProtocol, website.Website)
		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}
		want := "http://inews.gtimg.com/newsapp_ls/0/13753802147_640330/0"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("TCP:https and no '//'", func(t *testing.T) {
		picurl := "/inews.gtimg.com/newsapp_ls/0/13753802147_640330/0"
		website := Website{
			"qq.com",
			"title",
			".pic img",
			"src",
			"https",
		}
		got, err := urlFormat(picurl, website.TCPProtocol, website.Website)
		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}
		want := "https://qq.com/inews.gtimg.com/newsapp_ls/0/13753802147_640330/0"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("TCP:http and no '//'", func(t *testing.T) {
		picurl := "/inews.gtimg.com/newsapp_ls/0/13753802147_640330/0"
		website := Website{
			"qq.com",
			"title",
			".pic img",
			"src",
			"http",
		}
		got, err := urlFormat(picurl, website.TCPProtocol, website.Website)
		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}
		want := "http://qq.com/inews.gtimg.com/newsapp_ls/0/13753802147_640330/0"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("TCP:https and 'https://'", func(t *testing.T) {
		picurl := "https://inews.gtimg.com/newsapp_ls/0/13753802147_640330/0"
		website := Website{
			"qq.com",
			"title",
			".pic img",
			"src",
			"https",
		}
		got, err := urlFormat(picurl, website.TCPProtocol, website.Website)
		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}
		want := "https://inews.gtimg.com/newsapp_ls/0/13753802147_640330/0"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("TCP:http and 'http://'", func(t *testing.T) {
		picurl := "http://inews.gtimg.com/newsapp_ls/0/13753802147_640330/0"
		website := Website{
			"qq.com",
			"title",
			".pic img",
			"src",
			"https",
		}
		got, err := urlFormat(picurl, website.TCPProtocol, website.Website)
		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}
		want := "http://inews.gtimg.com/newsapp_ls/0/13753802147_640330/0"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
