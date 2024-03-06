# WeChat Rest

基于 [WeChatFerry RPC](https://github.com/lich0821/WeChatFerry/tree/master/WeChatFerry) 实现，主要特性如下：

- 使用 Go 语言编写，无运行时依赖
- 提供 HTTP 接口，便于对接各类编程语言
- 提供 Websocket 接口，接收推送的新消息
- 支持 HTTP/WS 接口授权，参见 [配置说明](#配置说明)
- 支持作为 SDK 使用，参见 [wcferry/README.md](./wcferry/README.md)
- 内置 AI 机器人，参见 [wclient/README.md](./wclient/README.md)
- 内置 Web 管理界面，可以管理机器人各项配置
- 内置 Api 调试工具，所有接口都可以在线调试
- 尽可能将消息中的 Xml 转为 Object，便于前端解析

## 快速开始

请仔细阅读本文档和[常见问题](#常见问题)后再开始使用；首次使用可参照下面的步骤开始：

- 下载并安装 [WeChatSetup-3.9.2.23.exe](https://github.com/opentdp/wechat-rest/releases/download/v0.0.1/WeChatSetup-3.9.2.23.exe) 和 [wechat-rest.zip](https://github.com/opentdp/wechat-rest/releases)

  - 非开发者请直接下载编译好的二进制文件，不要下载源码

- 双击 `wrest.exe` 将自动启动微信和接口服务，扫码登录微信

  - 初始化时若出现 *Attempt to access invalid address* 信息可忽略
  - 启动成功后，浏览器访问 `http://localhost:7600` 配置机器人

## 配置文件

机器人相关参数均已支持从 WEB 界面管理，[config.yml](./config.yml) 用来配置一些核心能力，一般情况下保持默认即可。

- 修改 `config.yml` 中的参数，需重启 **wrest.exe** 才能生效

  - 请使用 `Ctrl + C` 终止 **wrest.exe**，切勿直接关闭任务窗口
  - 重启时，提示端口被占用，请退出微信后重试

- 设置 `Web.Token` 后，请求接口时必须携带 **header** 信息: `Authorization: Bearer $token`

## 开发指南

查看和调试 *HTTP/WS* 接口，请使用浏览器访问 `http://localhost:7600/swagger/`

### API 模块

实现了 HTTP/WS 接口，详情查看 [httpd/README.md](./httpd/README.md)

### BOT 模块

实现了群聊机器人，详情查看 [wclient/README.md](./wclient/README.md)

### SDK 模块

实现了 WCF 客户端，详情查看 [wcferry/README.md](./wcferry/README.md)

### WEB 模块

实现了 WEB 控制台，详情查看 [webview/README.md](./webview/README.md)

## 代码提交

提交代码时请使用 `feat: something` 作为说明，支持的标识如下

- `feat` 新功能（feature）
- `fix` 错误修复
- `docs` 文档更改（documentation）
- `style` 格式（不影响代码含义的更改，空格、格式、缺少分号等）
- `refactor` 重构（即不是新功能，也不是修补bug的代码变动）
- `perf` 优化（提高性能的代码更改）
- `test` 测试（添加缺失的测试或更正现有测试）
- `chore` 构建过程或辅助工具的变动
- `revert` 还原以前的提交

## 常见问题

### Q1 注入失败

当前分支兼容的 PC 微信版本是 `3.9.2.23`，请在  [快速开始](#快速开始) 中点击下载

### Q2 如何在群内 `@` 其他人

首先要在消息中添加 `@昵称`，然后在 `aters` 参数添加此人的 `wxid`。相关接口 `/wcf/send_txt`

### Q3 如何更新机器人，并保留配置信息

从 [快速开始](#快速开始) 中下载新版本。关闭机器人后，将解压出来的 `wrest.exe` 和 `wcferry` 覆盖过去即可
