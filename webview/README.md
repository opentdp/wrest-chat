# Wrest Chat Web

基于 [wcferry-gohttp](https://github.com/opentdp/wrest-chat/tree/master/httpd) 接口服务实现的 Web 界面。

## 使用方法

将编译产物 `public/browser` 文件夹内的所有文件覆盖到项目根目录的 `public` 文件夹内。

## 功能列表

- 机器人
  - 扫码登录
  - 全局配置
  - 模型配置
  - 群聊配置
  - 用户配置
  - 词库管理
  - 计划任务

- Wcferry
  - 群聊
  - 通讯录
  - 数据库
  - 消息监听

## 常用命令

- 开发模式 `npm run start`
- 编译前端项目 `npm run build`
- 生成 jssdk 文件 `npm run build:jssdk`
