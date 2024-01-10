# 微信 REST API

基于 [WeChatFerry RPC](https://github.com/lich0821/WeChatFerry/tree/master/WeChatFerry) 实现，主要特性如下：

- 使用 Go 语言编写，无运行时依赖
- 基于 HTTP 提供操作接口，无缝对接大多数编程语言
- 支持作为标准 SDK 使用，参见 [wcferry/README.md](./wcferry/README.md)
- 内置互动机器人，参见 [wclient/README.md](./wclient/README.md)
- 内置 OpenApi 文档，参见 `http://localhost:7600`
- 支持 HTTP 接口授权，参见 [配置说明](#配置说明)
- 消息中的 Xml 尽可能转为 Object

## 使用方法

1、下载并安装 [WeChatSetup-3.9.2.23](https://github.com/opentdp/wechat-rest/releases/download/v0.0.1/WeChatSetup-3.9.2.23.exe) 和 [Wechat-rest](https://github.com/opentdp/wechat-rest/releases)

2、双击 `wrest.exe` 将自动启动微信和接口服务，扫码登录

> 初始化时出现 **Attempt to access invalid address** 错误信息可以忽略

3、浏览器打开 `http://localhost:7600` 查看支持的接口

4、修改 `config.yml` 配置机器人参数，重启 wrest 和 wechat 后生效

## 配置说明

启动 `wrest` 时将自动创建一个默认配置文件，完整配置说明可参考开源仓库中的 [config.yml](./config.yml)

- 应使用 `Ctrl + C` 终止 **wrest**，而非直接关闭 **wrest** 窗口
- 若设置了 `token`，请求时需携带 **header** 信息: `Authorization: Bearer $token`
- 免费申请 `Google AI API` 请登录 <https://makersuite.google.com>

## API 模块

实现了 HTTP 接口，详情查看 [httpd/README.md](./httpd/README.md)

## BOT 模块

实现了群聊机器人，详情查看 [wclient/README.md](./wclient/README.md)

## SDK 模块

实现了 WCF 客户端，详情查看 [wcferry/README.md](./wcferry/README.md)

## 开发说明

### 编译须知

由于微信和WCF均为32位应用，所以`go`也必须以`32`位模式编译，务必设置 `GOARCH` 环境变量为 `386`

### 生成 OpenApi 文档

```shell
go get github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag

swag init --parseDependency -g httpd/server.go -o public/swag -ot json
```
