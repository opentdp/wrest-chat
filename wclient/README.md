# WeChat Rest Bot

基于 [wcferry-gosdk](https://github.com/opentdp/wechat-rest/tree/master/wcferry) 实现的微信机器人，已开放如下功能：

- 自动回应拍一拍
- 自动接受好友请求
- 接受好友请求后主动回复
- 自动拉人进群聊天
- 有人进群后主动回复
- 支持配置黑/白名单
- 支持禁止用户使用助手
- 支持用户定义唤醒词
- 支持用户选择对话模型
- 支持 Google gemini 模型（支持图片识别）
- 支持 OpenAI gpt 模型
- 支持 讯飞 Spark 模型
- 支持 百度千帆（文心一言）模型
- 支持 腾讯混元 模型
- 从网络获取图片或文件发送到群里
- 通过API查询天气、新闻、图片、视频等
- 统计群聊活跃信息（水王、图王）

## 菜单示例

加群指令（`jr:xxx`）在后台添加群聊时指定，模型切换指令（`cm:xxx`）在后台添加模型时指定，外部接口指令（`/api xxx`）在后台全局设置中指定。

### 私聊菜单

```text
【/ai】 提问或交谈
【/ai:new】 重置上下文内容
【/ai:rand】 随机选择一个模型
【/cm:gpt】 换模型：GPT [gpt-3.5-turbo]
【/cm:gem】 换模型：Gemini [gemini-pro]
【/cm:spa】 换模型：Spark [v3]
【/api】 调用远程接口
【/bad】 添加违禁词
【/unbad】 删除违禁词
【/jr:chat】 加群聊：OpenTDP 聊天
【/jr:dev】 加群聊：OpenTDP 开发
【/wget】 获取图片或文件
【/help】 查看帮助信息
----------------
级别 9；唤醒词 ai；对话模型 GPT；上下文长度 0/20；祝你好运！
```

### 群聊菜单

```text
【/ai】 提问或交谈
【/ai:new】 重置上下文内容
【/ai:rand】 随机选择一个模型
【/cm:gpt】 换模型：GPT [gpt-3.5-turbo]
【/cm:gem】 换模型：Gemini [gemini-pro]
【/cm:spa】 换模型：Spark [v3]
【/api】 调用远程接口
【/bad】 添加违禁词
【/unbad】 删除违禁词
【/ban】 拉黑指定的用户
【/unban】 解封拉黑的用户
【/wget】 获取图片或文件
【/top】 获取群聊统计信息
【/help】 查看帮助信息
----------------
级别 9；对话模型 Gemini；上下文长度 0/20；祝你好运！
```

### 外部接口菜单

```text
【/api icp qq.com】 查询域名备案信息 <kapi.9kr.cc>
【/api img 大山】 按关键字返回图片
【/api ip 1.2.3.4】 查询IP地址信息
【/api iptv】 获取 IPTV 源数据
【/api lbs 南山】 地址解析，地址转坐标
【/api lbs 纬度,经度】 逆地址解析，坐标位置描述
【/api news】 获取今日热点
【/api paper 北京】 生成咸鱼日报
【/api port 1.2.3.4】 查询IP/域名端口
【/api price 商品URL】 查询商品历史价格
【/api site qq.com】 查询站点信息
【/api spam 文本】 检测违禁词
【/api video 北极】 按关键字返回视频
【/api weather 上海】 根据地址查询天气
【/api weather 经度,纬度】 根据经纬度查询天气
【/api whois qq.com】 获取域名 Whois 信息 <whois.ddnsip.cn>
```

## 消息类型

```go
switch msg.Type {
    case 0: //朋友圈消息
    case 1: //文字
        receiver1(msg)
    case 3: //图片
    case 34: //语音
    case 37: //好友确认
        receiver37(msg)
    case 40: //POSSIBLEFRIEND_MSG
    case 42: //名片
    case 43: //视频
    case 47: //石头剪刀布 | 表情图片
    case 48: //位置
    case 49: //共享实时位置、文件、转账、链接、群邀请
    case 50: //VOIPMSG
    case 51: //微信初始化
    case 52: //VOIPNOTIFY
    case 53: //VOIPINVITE
    case 62: //小视频
    case 66: //微信红包
    case 9999: //SYSNOTICE
    case 10000: //红包、系统消息
        receiver10000(msg)
    case 10002: //撤回消息
        receiver10002(msg)
    case 1048625: //搜狗表情
    case 16777265: //链接
    case 436207665: //微信红包
    case 536936497: //红包封面
    case 754974769: //视频号视频
    case 771751985: //视频号名片
    case 822083633: //引用消息
    case 922746929: //拍一拍
    case 973078577: //视频号直播
    case 974127153: //商品链接
    case 975175729: //视频号直播
    case 1040187441: //音乐链接
    case 1090519089: //文件case
}
```
