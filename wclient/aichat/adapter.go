package aichat

import (
	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/dbase/profile"
)

func Text(id, msg string) string {

	var err error
	var res string

	if len(args.LLM.Models) == 0 {
		return "未配置大语言模型"
	}

	// 预设模型参数
	CountHistory(id)
	llmc := UserModel(id)

	// 调用接口生成文本
	switch llmc.Provider {
	case "google":
		res, err = GoogleText(id, msg)
	case "openai":
		res, err = OpenaiText(id, msg)
	case "xunfei":
		res, err = XunfeiText(id, msg)
	default:
		res = "暂不支持此模型"
	}

	// 返回结果
	if err != nil {
		return err.Error()
	}
	return res

}

func UserModel(wxid string) *args.LLModel {

	var llmc *args.LLModel

	up, _ := profile.Fetch(&profile.FetchParam{Wxid: wxid})

	if up != nil {
		llmc = args.LLM.Models[up.AiModel]
	}

	if llmc == nil {
		llmc = args.LLM.Models[args.LLM.Default]
	}

	if llmc == nil {
		for _, v := range args.LLM.Models {
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
