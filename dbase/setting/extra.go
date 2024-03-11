package setting

import (
	"strconv"
)

var (
	// 是否启用内置机器人
	BotEnable = true
	// 是否自动同意新的好友请求
	FriendAccept = true
	// 添加好友后的响应内容，留空则忽略
	FriendHello = "群主去修仙了，请留言"
	// 是否自动回应拍一拍
	PatReturn = true
	// 私聊撤回消息时响应的内容，留空则忽略
	RevokeMsg = "撤回了寂寞？"
	// 自动下载消息中的图片
	AutoSaveImage = true
	// 开启后只有白名单内的群或好友可以使用机器人
	WhiteLimit = false
	// 用户的默认模型代码
	ModelDefault = ""
	// 定义模型扮演的身份
	ModelContext = "你是由OpenTDP开发的群助手，必须使用尽可能少的字数回答接下来的所有问题"
	// 历史消息数量
	ModelHistory = 20
	// 未注册指令时响应的内容，仅对"/"开头的指令有效
	InvalidHandler = "-"
	// API 指令请求的网址
	ApiEndpoint = "https://api.rehi.org/format=yaml/"
	// API 指令卡片使用的图标
	ApiEndpointIcon = "https://api.rehi.org/assets/icon.png"
	// HELP 指令扩展内容
	HelpAdditive = ""
)

// 从数据库加载配置

func Laod() {

	items, err := FetchAll(&FetchAllParam{})
	if err != nil {
		return
	}

	for _, item := range items {
		switch item.Name {
		case "BotEnable":
			BotEnable = item.Value == "true"
		case "FriendAccept":
			FriendAccept = item.Value == "true"
		case "FriendHello":
			FriendHello = item.Value
		case "PatReturn":
			PatReturn = item.Value == "true"
		case "RevokeMsg":
			RevokeMsg = item.Value
		case "AutoSaveImage":
			AutoSaveImage = item.Value == "true"
		case "WhiteLimit":
			WhiteLimit = item.Value == "true"
		case "ModelDefault":
			ModelDefault = item.Value
		case "ModelContext":
			ModelContext = item.Value
		case "ModelHistory":
			ModelHistory, _ = strconv.Atoi(item.Value)
		case "InvalidHandler":
			InvalidHandler = item.Value
		case "ApiEndpoint":
			ApiEndpoint = item.Value
		case "ApiEndpointIcon":
			ApiEndpointIcon = item.Value
		case "HelpAdditive":
			HelpAdditive = item.Value
		}
	}

}

// 迁移

func DataMigrate() {

	settings := []*CreateParam{
		{0, "BotEnable", "bool", "bot", strconv.FormatBool(BotEnable), "机器人", "是否启用机器人，重启生效"},
		{0, "FriendAccept", "bool", "bot", strconv.FormatBool(FriendAccept), "确认好友", "是否自动同意新的好友请求"},
		{0, "FriendHello", "string", "bot", FriendHello, "好友打招呼", "添加好友后的响应内容"},
		{0, "PatReturn", "bool", "bot", strconv.FormatBool(PatReturn), "回应拍拍", "私聊是否自动回应拍一拍"},
		{0, "RevokeMsg", "string", "bot", RevokeMsg, "撤回提醒", "私聊撤回消息时响应的内容"},
		{0, "AutoSaveImage", "bool", "bot", strconv.FormatBool(AutoSaveImage), "自动保存图片", "是否自动下载消息中的图片"},
		{0, "WhiteLimit", "bool", "bot", strconv.FormatBool(WhiteLimit), "白名单模式", "开启后仅白名单内的群或好友可以使用机器人"},
		{0, "ModelDefault", "lmodel", "bot", ModelDefault, "默认模型", "用户的默认大模型代码"},
		{0, "ModelContext", "text", "bot", ModelContext, "模型预定义", "大模型扮演的身份定义"},
		{0, "ModelHistory", "number", "bot", strconv.Itoa(ModelHistory), "上下文总数", "AI 聊天最大记录数，值越大消耗的 Token 越多"},
		{0, "InvalidHandler", "text", "bot", InvalidHandler, "指令未注册", "未注册指令时响应的内容，仅对\"/\"开头的指令有效"},
		{0, "ApiEndpoint", "string", "bot", ApiEndpoint, "API 指令地址", "/api 指令请求的网址"},
		{0, "ApiEndpointIcon", "string", "bot", ApiEndpointIcon, "API 指令图标", "/api 卡片消息使用的图标"},
		{0, "HelpAdditive", "text", "bot", HelpAdditive, "HELP 指令扩展", "/help 指令扩展内容，可添加自定义菜单等"},
	}

	for _, item := range settings {
		if _, err := Fetch(&FetchParam{Name: item.Name}); err != nil {
			Create(item)
		}
	}

}
