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
	llmc, _ := GetUserConfig(id).LLModel, CountHistory(id)
	text := strings.TrimSpace(strings.TrimPrefix(msg, "/ai"))

	// 调用接口生成文本
	switch llmc.Provider {
	case "google":
		res, err = GoogleChat(id, text)
	case "openai":
		res, err = OpenaiChat(id, text)
	case "xunfei":
		res, err = XunfeiChat(id, text)
	default:
		res = "暂不支持此模型"
	}

	// 返回结果
	if err != nil {
		return err.Error()
	}
	return res

}

// User Config

type UserConfig struct {
	WakeWord string
	LLModel  *args.LLModel
}

var userConfigMap = make(map[string]*UserConfig)

func GetUserConfig(id string) *UserConfig {

	if _, ok := userConfigMap[id]; !ok {
		userConfigMap[id] = &UserConfig{"", args.LLM.Models[args.LLM.Default]}
	}

	return userConfigMap[id]

}

// Message History

type MsgHistory struct {
	Content string
	Role    string
}

var msgHistoryMap = make(map[string][]*MsgHistory)

func ResetHistory(id string) {

	msgHistoryMap[id] = []*MsgHistory{}

}

func CountHistory(id string) int {

	if _, ok := msgHistoryMap[id]; !ok {
		ResetHistory(id)
	}

	return len(msgHistoryMap[id])

}

func AppendHistory(id string, items ...*MsgHistory) {

	if len(msgHistoryMap[id]) >= args.LLM.HistoryNum {
		msgHistoryMap[id] = msgHistoryMap[id][len(items):]
	}

	msgHistoryMap[id] = append(msgHistoryMap[id], items...)

}
