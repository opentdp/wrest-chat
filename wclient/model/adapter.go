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

func GetUser(id string) *args.Member {

	if _, ok := args.Usr.Member[id]; !ok {
		args.Usr.Member[id] = &args.Member{
			AiModel: args.LLM.Default,
		}
	}

	return args.Usr.Member[id]

}

func GetUserModel(id string) *args.LLModel {

	user := GetUser(id)
	llmc := args.LLM.Models[user.AiModel]

	if llmc == nil {
		for k, v := range args.LLM.Models {
			user.AiModel = k
			return v
		}
	}

	return llmc

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
