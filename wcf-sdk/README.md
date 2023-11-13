# WeChatFerry go-sdk

WCF客户端go语言版，用于访问RCP服务，实现和微信进程的远程交互。

## 生成pb文件

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go
protoc --proto_path=. --go_out=. proto/*.proto
```

## 参考信息

https://github.com/lich0821/WeChatFerry/tree/master/clients/python

https://github.com/lich0821/WeChatFerry/tree/master/clients/go
