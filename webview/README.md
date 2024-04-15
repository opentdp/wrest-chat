# Wrest Chat Web

基于 [wcferry-gohttp](https://github.com/opentdp/wrest-chat/tree/master/httpd) 接口服务实现的 Web 界面。

## 使用方法

将编译产物 `public/browser` 文件夹内的所有文件覆盖到项目根目录的 `public` 文件夹内。

## 功能菜单

- 首页
- 全局配置
- 模型配置
- 群聊配置
- 用户配置
- 词库管理
- 计划任务
- Webhook
- Wcferry
  - 群聊
  - 通讯录
  - 数据库
  - 消息监听
- 杂项
  - 可用指令
  - 插件状态
    - 关键词
    - 计划任务

## 常用命令

- 开发模式 `npm run start`
- 编译前端项目 `npm run build`
- 生成 jssdk 文件 `npm run build:jssdk`
- 临时忽略修改 `git update-index --assume-unchanged webview/proxy.conf.json`
