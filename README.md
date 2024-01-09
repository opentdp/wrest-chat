# 微信 REST API

基于 [WeChatFerry RPC](https://github.com/lich0821/WeChatFerry/tree/master/WeChatFerry) 实现的电脑版微信 REST API，主要特性如下：

- 使用 Go 语言编写，无运行时依赖
- 基于 HTTP 提供操作接口，无缝对接大多数编程语言
- 支持作为标准 SDK 使用，参见 [wcferry/README.md](./wcferry/README.md)
- 内置互动机器人，参见 [wclient/README.md](./wclient/README.md)
- 内置 OpenApi 文档，参见 `http://localhost:7600`
- 支持 HTTP 接口授权，参见 [配置说明](#配置说明)
- 消息中的 Xml 尽可能转为 Object

## 使用方法

1、下载并安装 [WeChatSetup-3.9.2.23](https://github.com/opentdp/wechat-rest/releases/download/v0.0.1/WeChatSetup-3.9.2.23.exe) 和 [Wechat-rest](https://github.com/opentdp/wechat-rest/releases)

2、双击 `wrest.exe` 将自动启动微信和接口服务，扫码登录。

> 初始化时出现 **Attempt to access invalid address** 错误信息可以忽略

3、浏览器打开 `http://localhost:7600` 查看支持的接口

> 接口使用范例请参考 <https://github.com/opentdp/wechat-robot>

## 配置说明

启动 `wrest` 时将自动创建一个默认配置文件，完整配置说明可参考开源仓库中的 [config.yml](./config.yml)

> 若设置了 `token`，请求时需携带 **header** 信息: `Authorization: Bearer $token`

## 功能清单

- 检查登录状态
- 获取登录账号 wxid
- 获取登录账号个人信息
- 获取所有消息类型
- 获取完整通讯录
- 获取好友列表
- 获取所有数据库
- 获取数据库中所有表
- 执行 SQL 查询
- 发送文本消息（可 @）
- 发送图片
- 发送文件
- 发送卡片消息
- 保存图片
- 保存语音
- 图片 OCR
- 接受好友申请
- 接收转账
- 刷新朋友圈
- 添加群成员
- 删除群成员
- 获取群列表
- 获取群成员列表
- 获取群成员昵称
- 邀请群成员
- 拍一拍群友
- 转发消息给好友
- 转发收到的消息到URL

## 开发说明

### 编译须知

由于微信和WCF均为32位应用，所以`go`也必须以`32`位模式编译，务必设置 `GOARCH` 环境变量为 `386`

### 生成 OpenApi 文档

```shell
go get github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag

swag init --parseDependency -g httpd/server.go -o public/swag -ot json
```
