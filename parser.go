package main

type Parser struct {
	SupportWebsites []Website `json:"support_websites"`
}
type Website struct {
	Website        string `json:"website"`
	TitlePattern   string `json:"title_pattern"`
	ImgPattern     string `json:"img_pattern"`
	ImgAddrPattern string `json:"img_addr_pattern"`
}

var DefaultJson = []byte(`
{
	"support_websites":[
		{
			"website": "cnblogs.com",
			"title_pattern": "title",
			"img_pattern" : ".navbar-branding img",
			"img_addr_pattern": "src"
		},
		{
			"website": "qq.com",
			"title_pattern": "title",
			"img_pattern" : ".pic img",
			"img_addr_pattern": "src"
		},
		{
			"website": "haicoder.net",
			"title_pattern": "title",
			"img_pattern" : ".logo img",
			"img_addr_pattern": "src"
		},
		{
			"website": "juejin.cn",
			"title_pattern": "title",
			"img_pattern" : ".mobile img",
			"img_addr_pattern": "src"
		},
		{
			"website": "www.yili.com",
			"title_pattern": "title",
			"img_pattern" : ".index-focus-item img",
			"img_addr_pattern": "src"
		}
	]
}`)
