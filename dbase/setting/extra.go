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
	// 开启后只有白名单内的群或好友可以使用机器人
	WhiteLimit = false
	// 用户的默认模型代码
	ModelDefault = ""
	// 定义模型扮演的身份
	ModelContext = "你是由OpenTDP开发的群助手，必须使用尽可能少的字数回答接下来的所有问题"
	// 历史消息数量
	ModelHistory = 20
	// API 指令请求的网址
	ApiEndpoint = "https://api.rehi.org/format=yaml/"
	// API 指令卡片使用的图标
	ApiEndpointIcon = "https://api.rehi.org/assets/icon.png"
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
		case "WhiteLimit":
			WhiteLimit = item.Value == "true"
		case "ModelDefault":
			ModelDefault = item.Value
		case "ModelContext":
			ModelContext = item.Value
		case "ModelHistory":
			ModelHistory, _ = strconv.Atoi(item.Value)
		case "ApiEndpoint":
			ApiEndpoint = item.Value
		case "ApiEndpointIcon":
			ApiEndpointIcon = item.Value
		}
	}

}

// 迁移

func DataMigrate() {

	settings := []*CreateParam{
		{"BotEnable", "bool", "bot", strconv.FormatBool(BotEnable), "机器人", "是否启用机器人，重启后生效"},
		{"FriendAccept", "bool", "bot", strconv.FormatBool(FriendAccept), "确认好友", "是否自动同意新的好友请求"},
		{"FriendHello", "string", "bot", FriendHello, "好友打招呼", "添加好友后的响应内容，“-”表示关闭"},
		{"PatReturn", "bool", "bot", strconv.FormatBool(PatReturn), "回应拍拍", "私聊是否自动回应拍一拍"},
		{"RevokeMsg", "string", "bot", RevokeMsg, "撤回提醒", "私聊撤回消息时响应的内容，“-”表示关闭"},
		{"WhiteLimit", "bool", "bot", strconv.FormatBool(WhiteLimit), "白名单模式", "开启后仅白名单内的群或好友可以使用机器人"},
		{"ModelDefault", "lmodel", "bot", ModelDefault, "默认模型", "用户的默认模型代码"},
		{"ModelContext", "text", "bot", ModelContext, "模型预定义", "定义模型扮演的身份"},
		{"ModelHistory", "number", "bot", strconv.Itoa(ModelHistory), "上下文总量", "历史消息最大数量"},
		{"ApiEndpoint", "string", "bot", ApiEndpoint, "API 地址", "API 指令请求的网址，“-”表示关闭"},
		{"ApiEndpointIcon", "string", "bot", ApiEndpointIcon, "API 图标", "API 指令卡片使用的图标"},
	}

	for _, item := range settings {
		if _, err := Fetch(&FetchParam{Name: item.Name}); err != nil {
			Create(item)
		}
	}

}
