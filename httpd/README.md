# WeChat Rest Api

基于 [wcferry](https://github.com/opentdp/wechat-rest/tree/master/wcferry) 实现的 HTTP 接口服务，已实现如下功能：

## 支持的接口列表

### JOB::计划任务

- 添加计划任务
- 删除计划任务
- 获取计划任务
- 计划任务列表
- 计划任务状态
- 修改计划任务

### BOT::群聊配置

- 添加群聊配置
- 删除群聊配置
- 获取群聊配置
- 群聊配置列表
- 修改群聊配置

### BOT::关键字

- 添加关键字
- 删除关键字
- 获取关键字
- 关键字列表
- 修改关键字

### BOT::大语言模型

- 添加模型
- 删除模型
- 获取模型
- 模型列表
- 修改模型

### BOT::用户配置

- 添加用户配置
- 删除用户配置
- 获取用户配置
- 用户配置列表
- 修改用户配置

### BOT::全局配置

- 添加全局配置
- 删除全局配置
- 获取全局配置
- 全局配置列表
- 修改全局配置

### WCF::联系人管理

- 接受好友请求
- 获取头像列表
- 获取完整通讯录
- 获取好友列表
- 根据wxid获取个人信息

### WCF::群聊管理

- 添加群成员
- 获取群成员昵称
- 获取群成员列表
- 获取群列表
- 删除群成员
- 邀请群成员

### WCF::数据库查询

- 获取数据库列表
- 执行数据库查询
- 获取数据库表列表

### WCF::消息推送

- 关闭推送消息到URL
- 开启推送消息到URL
- GET
- 推送消息到Socket

### WCF::消息收取

- 下载附件
- 下载图片
- 获取语音消息
- 获取OCR识别结果
- 接受转账

### WCF::消息发送

- 转发消息
- 撤回消息
- 发送文件消息
- 发送图片消息
- 拍一拍群友
- 发送卡片消息
- 发送文本消息

### WCF::其他

- 检查登录状态
- 登录二维码
- 获取所有消息类型
- 刷新朋友圈
- 获取登录账号个人信息
- 获取登录账号wxid

## 生成 OpenApi 文档

```shell
go get github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag

swag init --parseDependency -g httpd/server.go -o public/swagger -ot json
```

## 生成 OpenApi 客户端

将生成的 `swagger.json` 上传至 `https://editor.swagger.io` 生成对应的客户端
