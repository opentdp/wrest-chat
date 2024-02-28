# WeChat Rest Bot

基于 [wcferry](https://github.com/opentdp/wechat-rest/tree/master/wcferry) 实现的微信机器人，已实现如下功能：

- 自动回应拍一拍
- 自动接受好友请求
- 接受好友请求后主动回复
- 自动拉人进群聊天
- 有人进群后主动回复
- 支持 Google gemini 模型
- 支持 OpenAI gpt 模型
- 支持 讯飞 Spark 模型
- 支持用户选择对话模型
- 支持用户定义唤醒词
- 支持禁止用户使用助手
- 支持配置黑/白名单
- 支持用户设置唤醒词
- 从网络获取图片或文件发送到群里
- 通过API查询天气、新闻、图片、视频等

## 菜单示例

```text
/ai 提问或交谈
/api 调用远程接口
/bad 添加违规关键词
/ban 拉黑指定的用户
/g:chat 加入群聊 OpenTDP 闲聊群
/g:dev 加入群聊 OpenTDP 开发群
/help 查看帮助信息
/m:gem 切换为 Gemini [gemini-pro]
/m:gpt 切换为 GPT [gpt-3.5-turbo]
/m:spa 切换为 Spark [v3]
/new 重置上下文内容
/save 保存配置信息
/unbad 添加违规关键词
/unban 解封拉黑的用户
/wake 设置或禁用唤醒词
/wget 获取图片或文件
级别 9；对话模型 Gemini，上下文长度 0/20；祝你好运！
```

```text
/api icp qq.com: 查询域名备案信息
/api img 大山: 按关键字返回图片
/api ip 1.2.3.4: 查询IP地址信息
/api lbs 南山: 获取地理编码
/api news weibo: 获取热点
/api news ithome: 获取热点
/api port 1.2.3.4: 查询IP/域名端口
/api site qq.com: 查询站点信息
/api spam 文本: 检测违规内容
/api video 北极: 按关键字返回视频
/api weather 上海: 查询城市天气
/api whois qq.com: 查询域名 Whois
```

## 其他

- `Google AI` 免费申请入口 <https://makersuite.google.com>
