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
	// 默认模型
	ModelDefault = ""
	// 定义模型扮演的身份
	ModelContext = "你是由OpenTDP开发的群助手，必须使用尽可能少的字数回答接下来的所有问题"
	// 历史消息数量
	ModelHistory = 20
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
		}
	}

}

// 将配置保存到数据库

func Save() {

	Update(&UpdateParam{Name: "BotEnable", Value: strconv.FormatBool(BotEnable)})
	Update(&UpdateParam{Name: "FriendAccept", Value: strconv.FormatBool(FriendAccept)})
	Update(&UpdateParam{Name: "FriendHello", Value: FriendHello})
	Update(&UpdateParam{Name: "PatReturn", Value: strconv.FormatBool(PatReturn)})
	Update(&UpdateParam{Name: "RevokeMsg", Value: RevokeMsg})
	Update(&UpdateParam{Name: "WhiteLimit", Value: strconv.FormatBool(WhiteLimit)})
	Update(&UpdateParam{Name: "ModelDefault", Value: ModelDefault})
	Update(&UpdateParam{Name: "ModelContext", Value: ModelContext})
	Update(&UpdateParam{Name: "ModelHistory", Value: strconv.Itoa(ModelHistory)})

}

// 迁移

func DataMigrate() {

	if c, _ := Count(&CountParam{}); c == 0 {
		return
	}

	Create(&CreateParam{"BotEnable", "bool", "bot", "", "机器人", "是否启用机器人，重启后生效"})
	Create(&CreateParam{"FriendAccept", "bool", "bot", "", "确认好友", "是否自动同意新的好友请求"})
	Create(&CreateParam{"FriendHello", "string", "bot", "", "好友打招呼", "添加好友后的响应内容，留空则忽略"})
	Create(&CreateParam{"PatReturn", "bool", "bot", "", "回应拍拍", "私聊是否自动回应拍一拍"})
	Create(&CreateParam{"RevokeMsg", "string", "bot", "", "撤回提醒", "私聊撤回消息时响应的内容，留空则忽略"})
	Create(&CreateParam{"WhiteLimit", "bool", "bot", "", "白名单", "开启后只有白名单内的群或好友可以使用机器人"})
	Create(&CreateParam{"ModelDefault", "string", "bot", "", "默认模型", "用户的默认大模型代码"})
	Create(&CreateParam{"ModelContext", "string", "bot", "", "模型预定义", "定义模型扮演的身份"})
	Create(&CreateParam{"ModelHistory", "number", "bot", "", "上下文总量", "历史消息最大数量"})

	Save() // 将默认配置存入数据库

}
