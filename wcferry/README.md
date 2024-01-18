# WeChatFerry go-sdk

`WeChatFerry` 客户端 **Go** 语言版，用于管理 `WeChatFerry RPC` 服务，并实现与微信进程的互操作。接口文档请查阅 <https://pkg.go.dev/github.com/opentdp/wechat-rest>

## 调用方式

```go
package main

import (
    "fmt"
    "github.com/opentdp/wechat-rest/wcferry"
)

func main() {
    wc := &wcferry.Client{}
    if err := wc.Connect(); err != nil {
        panic(err)
    }
    // 打印登录状态
    fmt.Println(wc.CmdClient.IsLogin())
    // 打印收到的消息
    wc.EnrollReceiver(true, wcferry.WxMsgPrinter)
    // 阻止程序退出
    select{}
}
```

## 生成pb文件

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go
protoc --proto_path=. --go_out=. proto/*.proto
```

## 参考信息

https://github.com/lich0821/WeChatFerry/tree/master/clients/python

https://github.com/lich0821/WeChatFerry/tree/master/clients/go
