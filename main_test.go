package main

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseUrl(t *testing.T) {
	t.Run("TCP:https and '//'", func(t *testing.T) {
		picurl := "//inews.gtimg.com/newsapp_ls/0/13753802147_640330/0"
		website := Website{
			Website:        "qq.com",
			TitlePattern:   "title",
			ImgPattern:     ".pic img",
			ImgAddrPattern: "src",
			TCPProtocol:    "https",
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
func TestReloadParser(t *testing.T) {
	jsonFileAddr := fmt.Sprint("./testdata/test.json")
	got, err := reloadParser(jsonFileAddr)
	if err != nil {
		t.Fatalf("did not expect an error but got one %v", err)
	}
	jsonString := []byte(`{
  "support_websites":[
    {
      "website": "cnblogs.com",
      "title_pattern": "title",
      "img_pattern" : ".navbar-branding img",
      "img_addr_pattern": "src",
      "tcp_protocol":"https"
    },
    {
      "website": "qq.com",
      "title_pattern": "title",
      "img_pattern" : ".pic img",
      "img_addr_pattern": "src",
      "tcp_protocol":"https"
    },
    {
      "website": "haicoder.net",
      "title_pattern": "title",
      "img_pattern" : ".logo img",
      "img_addr_pattern": "src",
      "tcp_protocol":"https"
    },
    {
      "website": "juejin.cn",
      "title_pattern": "title",
      "img_pattern" : ".mobile img",
      "img_addr_pattern": "src",
      "tcp_protocol":"https"
    },
    {
      "website": "www.yili.com",
      "title_pattern": "title",
      "img_pattern" : ".js-index-focus img",
      "img_addr_pattern": "src",
      "tcp_protocol":"https"
    },
    {
      "website": "bookstack.cn",
      "title_pattern": "title",
      "img_pattern" : ".pull-left img",
      "img_addr_pattern": "src",
      "tcp_protocol":"https"
    }
  ]
}`)
	var file Parser
	err = json.Unmarshal(jsonString, &file)
	want := &file
	require.Equal(t, want, got, "the two json should be same")
}
func TestFindDomainByUrl(t *testing.T) {
	t.Run("parser is in list", func(t *testing.T) {
		posturl := "https://www.qq.com"
		jsonfile, err := reloadParser("./testdata/test.json")
		got, err := findDomainByUrl(posturl, jsonfile)
		want := &Website{
			"qq.com",
			"title",
			".pic img",
			"src",
			"https",
		}
		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}
		require.Equal(t, want, got, "the two json should be same")
	})
}
