package aichat

import (
	"strings"

	"github.com/opentdp/wechat-rest/args"
)

func Text(id, msg string) string {

	var err error
	var res string

	if len(args.LLM.Models) == 0 {
		return "未配置大语言模型"
	}

	// 预设模型参数
	CountHistory(id) // init only
	llmc := args.GetMember(id).GetModel()
	text := strings.TrimSpace(strings.TrimPrefix(msg, "/ai"))

	// 调用接口生成文本
	switch llmc.Provider {
	case "google":
		res, err = GoogleText(id, text)
	case "openai":
		res, err = OpenaiText(id, text)
	case "xunfei":
		res, err = XunfeiText(id, text)
	default:
		res = "暂不支持此模型"
	}

	// 返回结果
	if err != nil {
		return err.Error()
	}
	return res

}

// Message History

type MsgHistory struct {
	Content string
	Role    string
}

var msgHistories = make(map[string][]*MsgHistory)

func ResetHistory(id string) {

	msgHistories[id] = []*MsgHistory{}

}

func CountHistory(id string) int {

	if _, ok := msgHistories[id]; !ok {
		ResetHistory(id)
	}

	return len(msgHistories[id])

}

func AppendHistory(id string, items ...*MsgHistory) {

	if len(msgHistories[id]) >= args.LLM.HistoryNum {
		msgHistories[id] = msgHistories[id][len(items):]
	}

	msgHistories[id] = append(msgHistories[id], items...)

}
