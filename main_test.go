package main

import (
	"testing"
)

func TestParseUrl(t *testing.T) {
	t.Run("TCP:https and '//'", func(t *testing.T) {
		picurl := "//inews.gtimg.com/newsapp_ls/0/13753802147_640330/0"
		website := Website{
			"qq.com",
			"title",
			".pic img",
			"src",
			"https",
		}
		got, err := parseUrl(picurl, &website)
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
		got, err := parseUrl(picurl, &website)
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
		got, err := parseUrl(picurl, &website)
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
		got, err := parseUrl(picurl, &website)
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
		got, err := parseUrl(picurl, &website)
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
		got, err := parseUrl(picurl, &website)
		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}
		want := "http://inews.gtimg.com/newsapp_ls/0/13753802147_640330/0"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

//func TestFindDomainByUrl(t *testing.T){
//	t.Run("parser is in list", func(t *testing.T) {
//		posturl:="https://www.qq.com"
//		got,err:=findDomainByUrl(posturl)
//		want:=&Website{
//			"qq.com",
//			"title",
//			".pic img",
//			"src",
//			"https",
//		}
//		if err!=nil{
//			t.Fatalf("did not expect an error but got one %v", err)
//		}
//		if got!=want{
//			t.Errorf("got %q, want %q", got, want)
//		}
//	})
//}
