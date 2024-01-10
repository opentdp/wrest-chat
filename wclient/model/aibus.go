package model

import (
	"strings"

	"github.com/opentdp/wechat-rest/args"
)

type HistoryItem struct {
	Role    string
	Content string
}

var History = make(map[string][]HistoryItem)

var Models = make(map[string]int)

func Clear(id, msg string) string {

	History[id] = []HistoryItem{}
	return "已清空上下文"

}

func Model(id string) *args.LLMModel {

	if len(args.LLM.Models) == 0 {
		return nil
	}
	if Models[id] >= len(args.LLM.Models) {
		return nil
	}

	return args.LLM.Models[Models[id]]

}

func AiChat(id, msg string) string {

	var err error
	var res string

	if len(args.LLM.Models) == 0 {
		return "未配置大语言模型"
	}

	// 预设模型参数
	if _, exists := Models[id]; !exists {
		Models[id] = 0
	}
	if _, exists := History[id]; !exists {
		History[id] = []HistoryItem{}
	}

	// 防止模型越界
	if Models[id] >= len(args.LLM.Models) {
		Models[id] = 0
	}
	if len(History[id]) > args.LLM.HistoryNum {
		History[id] = History[id][2:]
	}

	// 调用接口生成文本
	text := strings.TrimSpace(strings.TrimPrefix(msg, "/ai"))
	switch Model(id).Provider {
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
