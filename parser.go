package main

type Parser struct {
	SupportWebsites []Website `json:"support_websites"`
}
type Website struct {
	Website        string `json:"website"`
	TitlePattern   string `json:"title_pattern"`
	ImgPattern     string `json:"img_pattern"`
	ImgAddrPattern string `json:"img_addr_pattern"`
	TCPProtocol    string `json:"tcp_protocol"`
}

var DefaultJson = []byte(`
{
	"support_websites":[
		{
			"website": "qq.com",
			"title_pattern": "title",
			"img_pattern" : ".pic img",
			"img_addr_pattern": "src",
			"tcp_protocol":"https"
		},
		{
			"website": "www.yili.com/",
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
		},
		{
			"website": "sina.com.cn",
			"title_pattern": "title",
			"img_pattern" : ".uni-blk-pic img",
			"img_addr_pattern": "src",
			"tcp_protocol":"https"
		}
	]
}`)
