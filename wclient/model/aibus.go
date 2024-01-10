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
	if _, exists := UserModels[id]; !exists {
		UserModels[id] = 0
	}
	if _, exists := MsgHistory[id]; !exists {
		MsgHistory[id] = []*HistoryItem{}
	}

	// 防止模型越界
	if UserModels[id] >= len(args.LLM.Models) {
		UserModels[id] = 0
	}
	if len(MsgHistory[id]) > args.LLM.HistoryNum {
		MsgHistory[id] = MsgHistory[id][2:]
	}

	// 调用接口生成文本
	text := strings.TrimSpace(strings.TrimPrefix(msg, "/ai"))
	switch GetUserModel(id).Provider {
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

// History

type HistoryItem struct {
	Role    string
	Content string
}

var MsgHistory = make(map[string][]*HistoryItem)

func AddHistory(id string, items ...*HistoryItem) {

	MsgHistory[id] = append(MsgHistory[id], items...)

}

func CountHistory(id string) int {

	return len(MsgHistory[id])

}

func ClearHistory(id string) string {

	MsgHistory[id] = []*HistoryItem{}
	return "已清空上下文"

}

// UserModels

var UserModels = make(map[string]int)

func SetUserModel(id string, k int) string {

	if k >= len(args.LLM.Models) {
		return "模型未定义"
	}

	UserModels[id] = k
	MsgHistory[id] = []*HistoryItem{}

	v := args.LLM.Models[k]
	return "对话模型已切换为 " + v.Name + " [" + v.Model + "]"

}

func GetUserModel(id string) *args.LLModel {

	if len(args.LLM.Models) == 0 {
		return nil
	}
	if UserModels[id] >= len(args.LLM.Models) {
		return nil
	}

	return args.LLM.Models[UserModels[id]]

}
