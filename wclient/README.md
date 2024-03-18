# WeChat Rest Bot

基于 [wcferry-gosdk](https://github.com/opentdp/wechat-rest/tree/master/wcferry) 实现的微信机器人，已支持如下功能：

- 自动回应拍一拍
- 自动接受好友请求
- 接受好友请求后主动回复
- 自动拉人进群聊天
- 有人进群后主动回复
- 支持配置黑/白名单
- 支持禁止用户使用助手
- 支持用户定义唤醒词
- 支持用户选择对话模型
- 支持 Google gemini 模型（含图片识别）
- 支持 OpenAI gpt 模型
- 支持 讯飞 Spark 模型
- 支持 百度千帆（文心一言）模型
- 支持 腾讯混元 模型
- 支持 通义千问 模型
- 从网络获取图片或文件发送到群里
- 通过API查询天气、新闻、图片、视频等
- 统计群聊活跃信息（聊天总数、水王、图王）

## 菜单示例

加群指令（`jr:xxx`）在后台添加群聊时指定，模型切换指令（`cm:xxx`）在后台添加模型时指定，外部接口指令（`/api xxx`）在后台全局设置中指定。

### 私聊菜单

```text
【/ai】 提问或交谈
【/ai:new】 重置上下文内容
【/ai:rand】 随机选择一个模型
【/cm:gpt】 换模型 GPT [gpt-3.5-turbo]
【/cm:gem】 换模型 Gemini [gemini-pro]
【/cm:spa】 换模型 Spark [v3]
【/api】 调用远程接口
【/bad】 添加违禁词
【/unbad】 删除违禁词
【/jr:chat】 加群聊 OpenTDP 聊天
【/jr:dev】 加群聊 OpenTDP 开发
【/help】 查看帮助信息
----------------
级别 9；唤醒词 ai；对话模型 GPT；上下文长度 0/20；祝你好运！
```

### 群聊菜单

```text
【/ai】 提问或交谈
【/ai:new】 重置上下文内容
【/ai:rand】 随机选择一个模型
【/cm:gpt】 换模型 GPT [gpt-3.5-turbo]
【/cm:gem】 换模型 Gemini [gemini-pro]
【/cm:spa】 换模型 Spark [v3]
【/api】 调用远程接口
【/bad】 添加违禁词
【/unbad】 删除违禁词
【/ban】 拉黑指定的用户
【/unban】 解封拉黑的用户
【/top】 获取群聊统计信息
【/help】 查看帮助信息
----------------
级别 9；对话模型 Gemini；上下文长度 0/20；祝你好运！
```

### 远程接口菜单

```text
【/api img 大山】 按关键字返回图片
【/api ip 1.2.3.4】 查询IP地址信息
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

## 自定义 API 使用说明

`/api` 命令会将用户输入的参数和后台设置的 `ApiEndpoint` 一起构造为一个完整的 URL，并通过 `http.Get` 发送请求。为方便前端用户输入，指令和参数由空格分隔，且第三部分及之后的任意字符都将整体被转义，示例：

- `/api` 组装为 `https://example.com/help`
- `/api icp qq.com` 组装为 `https://example.com/icp/qq.com`
- `/api img one two` 组装为 `https://example.com/img/one%20two`
- `/api img one=two` 组装为 `https://example.com/img/one=two`

### 回调数据结构

API 回调支持返回文本和结构化两种数据。下面的 GO 结构体描述了返回的结构化数据，**转为 JSON 后对应的字段均为小写**。

```go
type ApiCallbackData struct {
    Type string
    Card struct { 
        Name    string
        Account string
        Title   string
        Digest  string
        Link    string
        Icon    string
    }
    Link string
    Text string
}
```

### 字段说明

| 字段   | 类型   | 说明                                                        |
| ------ | ------ | ----------------------------------------------------------- |
| `Type` | string | 数据类型，可选值为 `card`、`file`、`image`、`text`、`error` |
| `Card` | struct | 当 `Type` 为 `card` 时有效，详见下表                        |
| `Link` | string | 当 `Type` 为 `file` 或 `image` 时有效，指向文件的链接       |
| `Text` | string | 当 `Type` 为 `text` 或 `error` 时有效，文本内容             |

### `Card` 字段说明

| 字段      | 类型   | 说明                              |
| --------- | ------ | --------------------------------- |
| `Name`    | string | 左下显示的名字，可选              |
| `Account` | string | 公众号 id，可显示对应的头像，可选 |
| `Title`   | string | 标题，最多显示为两行              |
| `Digest`  | string | 摘要，最多显示为三行              |
| `Link`    | string | 点击后跳转的链接                  |
| `Icon`    | string | 右侧缩略图的链接，可选            |

### 服务端返回消息示例

```json
{
  "type": "card",
  "card": {
    "name": "公众号名称",
    "account": "公众号 id",
    "title": "标题",
    "digest": "摘要",
    "link": "链接",
    "icon": "缩略图链接"
  }
}
```

```json
{
  "type": "file",
  "link": "文件链接"
}
```

```json
{
  "type": "text",
  "text": "文本内容"
}
```

```json
{
  "type": "error",
  "text": "错误信息"
}
```

## 微信消息类型

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
