package model

import (
	"strings"

	"github.com/opentdp/wechat-rest/args"
)

func AiChat(id, msg string) string {

	var err error
	var res string

	if len(args.LLM.Models) == 0 {
		return "未配置大语言模型"
	}

	// 预设模型参数
	llmc, _ := GetUserModel(id), CountHistory(id)
	text := strings.TrimSpace(strings.TrimPrefix(msg, "/ai"))

	// 调用接口生成文本
	switch llmc.Provider {
	case "google":
		res, err = GoogleChat(id, text)
	case "openai":
		res, err = OpenaiChat(id, text)
	default:
		res = "暂不支持此模型"
	}

	// 返回结果
	if err != nil {
		return err.Error()
	}
	return res

}

// User Models

var UserModels = make(map[string]*args.LLModel)

func SetUserModel(id string, m *args.LLModel) string {

	ResetHistory(id)
	UserModels[id] = m

	return "对话模型已切换为 " + m.Name + " [" + m.Model + "]"

}

func GetUserModel(id string) *args.LLModel {

	if _, exists := UserModels[id]; !exists {
		SetUserModel(id, args.LLM.Models[0])
	}

	return UserModels[id]

}

// Message History

type HistoryItem struct {
	Content string
	Role    string
}

var MsgHistory = make(map[string][]*HistoryItem)

func ResetHistory(id string) string {

	MsgHistory[id] = []*HistoryItem{}
	return "已重置上下文"

}

func CountHistory(id string) int {

	if _, exists := MsgHistory[id]; !exists {
		ResetHistory(id)
	}

	return len(MsgHistory[id])

}

func AppendHistory(id string, items ...*HistoryItem) {

	if len(MsgHistory[id]) >= args.LLM.HistoryNum {
		MsgHistory[id] = MsgHistory[id][len(items):]
	}

	MsgHistory[id] = append(MsgHistory[id], items...)

}
