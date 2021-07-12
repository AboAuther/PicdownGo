# PicdownGo

#### 项目利用命令行实现了特定网站下载静态图片

### 安装方法
```bash
go get -u -x github.com/AboAuther/PicdownGo
```

### 使用方法
```bash
PicdownGo -u URL
```

##加载Json配置
json是一个文本文件，将对应的数据读取到后，赋值给全局变量。
```go
if checkFileIsExits(jsonFileAddr) {
    file, _ = os.ReadFile(jsonFileAddr)
    fmt.Println("Loading local json file...")
} else {
    err := errors.New("json file is not Existed")
    return err
}
if len(file) == 0 {
    fmt.Println("Json file is empty,Loading default parser...")
    file = DefaultJson
}
err = json.Unmarshal(file, &configuration)
```
## 解析域名
通过输入的URL与json文件中支持的网站列表对比，获取域名准备下载。
```go
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
```