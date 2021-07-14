package main

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

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
