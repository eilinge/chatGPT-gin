## mychatgpt

>使用框架<br>

go-gpt3
gin
wire
go-assets
logrus
swaggo

>检测pkg<br>

go mod tidy<br>

>启动demo<br>

go run main.go

>测试访问<br>

浏览器localhost:9090/index<br>

>修改脚本<br>

1. wire_gen.go 文件NewService()函数下填入token
2. 开启魔法, 调用openai接口时添加代理, 本地pkg `github.com/sashabaranov/go-gpt3/api.go`newTransport 方法

```go
func newTransport() *http.Client {
	proxyUrl, _ := url.Parse("socks5://127.0.0.1:10886") // 你的魔法地址
	return &http.Client{Transport: &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	},
		Timeout: time.Duration(30 * time.Second),
	}
}
```