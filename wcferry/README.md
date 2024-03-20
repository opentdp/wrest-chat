# Wrest Chat SDK

这是 Wrest Chat 协议适配层，通过预定义的 [Nanomsg 协议](proto/wcferry.proto) 与聊天软件实现互操作。只要遵循该协议实现聊天软件的适配即可接入其他助手功能。

## 调用方式

```go
package main

import (
    "fmt"
    "github.com/opentdp/wrest-chat/wcferry"
)

func main() {
    wc := &wcferry.Client{
        WcfBinary: "wcf.exe",
    }
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
