# WeChat Rest

基于 [WeChatFerry RPC](https://github.com/lich0821/WeChatFerry/tree/master/WeChatFerry) 实现，主要特性如下：

- 使用 Go 语言编写，无运行时依赖
- 提供 HTTP 接口，便于对接各类编程语言
- 提供 Websocket 接口，接收推送的新消息
- 支持 HTTP/WS 接口授权，参见 [配置说明](#配置说明)
- 支持作为 SDK 使用，参见 [wcferry/README.md](./wcferry/README.md)
- 内置 AI 机器人，参见 [wclient/README.md](./wclient/README.md)
- 内置 Web 管理界面，参见 `http://localhost:7600/`
- 内置 Api 调试工具，参见 `http://localhost:7600/swagger/`
- 尽可能将消息中的 Xml 转为 Object，便于前端解析

## 快速开始

1、下载并安装 [WeChatSetup-3.9.2.23.exe](https://github.com/opentdp/wechat-rest/releases/download/v0.0.1/WeChatSetup-3.9.2.23.exe) 和 [wechat-rest.zip](https://github.com/opentdp/wechat-rest/releases)

- 非开发者请直接下载编译好的二进制文件，不要下载源码

2、双击 `wrest.exe` 将自动启动微信和接口服务，扫码登录微信

- 初始化时若出现 *Attempt to access invalid address* 信息可忽略

3、修改 [config.yml](./config.yml) 配置机器人参数，重启 **wrest.exe** 后生效

- 请使用 `Ctrl + C` 终止 **wrest.exe**，切勿直接关闭任务窗口
- 重启时，提示端口被占用，请退出微信后重试

## 配置文件
  
- 如需使用智能机器人，请配置 `LLM.Models` 参数，并设置正确的模型密钥

- 如设置了 `Web.Token`，请求接口时需携带 **header** 信息: `Authorization: Bearer $token`

## 开发指南

- 查看和调试*HTTP*接口文档，请使用浏览器打开 `http://localhost:7600`

### API 模块

实现了 HTTP/WS 接口，详情查看 [httpd/README.md](./httpd/README.md)

### BOT 模块

实现了群聊机器人，详情查看 [wclient/README.md](./wclient/README.md)

### SDK 模块

实现了 WCF 客户端，详情查看 [wcferry/README.md](./wcferry/README.md)

### WEB 模块

实现了 WEB 控制台，详情查看 [webview/README.md](./webview/README.md)
