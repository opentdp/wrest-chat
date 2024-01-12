# WeChat Rest Bot

基于 [wcferry](https://github.com/opentdp/wechat-rest/tree/master/wcferry) 实现的微信机器人，已实现如下功能：

- 自动添加好友
- 支持自动拉群
- 自动回应拍一拍
- 支持 Google AI 聊天
- 支持 OpenAi 聊天
- 支持用户选择模型
- 支持用户定义唤醒词

## 菜单示例

```
/ai 提问或交谈
/ban 禁止用户使用助手
/help 查看帮助信息
/m:gemini 切换为 gemini-pro 模型
/m:gpt 切换为 gpt-3.5-turbo 模型
/mr 随机选择模型
/new 重置上下文内容
/room:1 加入群聊 OpenTDP 开发群
/room:2 加入群聊 OpenTDP 闲聊群
/wake 设置或禁用唤醒词
对话模型 gemini，上下文长度 4/20
```
