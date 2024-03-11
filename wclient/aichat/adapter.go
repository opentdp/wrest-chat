package aichat

import (
	"encoding/base64"
	"net/http"
	"os"

	"github.com/opentdp/wechat-rest/dbase/chatroom"
	"github.com/opentdp/wechat-rest/dbase/llmodel"
	"github.com/opentdp/wechat-rest/dbase/profile"
	"github.com/opentdp/wechat-rest/dbase/setting"
	"github.com/opentdp/wechat-rest/dbase/tables"
)

func Text(id, rid, msg string) string {

	var err error
	var res string

	// 预设模型参数
	CountHistory(id)
	llmc := UserModel(id, rid)

	// 调用接口生成文本
	switch llmc.Provider {
	case "google":
		res, err = GoogleText(id, rid, msg)
	case "openai":
		res, err = OpenaiText(id, rid, msg)
	case "xunfei":
		res, err = XunfeiText(id, rid, msg)
	case "baidu":
		res, err = BaiDuText(id, rid, msg)
	case "tencent":
		res, err = TencentText(id, rid, msg)
	case "aliyun":
		res, err = AliyunText(id, rid, msg)
	case "":
		res = "当前模型已失效，请重新选择"
	default:
		res = "暂不支持此模型"
	}

	// 返回结果
	if err != nil {
		return err.Error()
	}
	return res

}

func Image(id, rid, msg, img string) string {

	var err error
	var res string

	// 预设模型参数
	CountHistory(id)
	llmc := UserModel(id, rid)

	// 调用接口生成文本
	switch llmc.Provider {
	case "google":
		res, err = GoogleImage(id, rid, msg, img)
	case "":
		res = "当前模型已失效，请重新选择"
	default:
		res = "当前模型不支持分析图片"
	}

	// 返回结果
	if err != nil {
		return err.Error()
	}
	return res

}

// 读取图片

func ReadImage(img string) (string, string) {

	fileContent, err := os.ReadFile(img)
	if err != nil {
		return "", ""
	}

	base64String := base64.StdEncoding.EncodeToString(fileContent)
	mimeType := http.DetectContentType(fileContent)

	return base64String, mimeType

}

// 用户模型

type UserLLModel struct {
	RoleContext  string
	ModelHistory int
	*tables.LLModel
}

func UserModel(id, rid string) *UserLLModel {

	llmc := &tables.LLModel{}
	modelContext := setting.ModelContext
	modelHistory := setting.ModelHistory

	// 用户模型配置
	up, _ := profile.Fetch(&profile.FetchParam{Wxid: id, Roomid: rid})
	if up.Rd > 0 {
		llmc, _ = llmodel.Fetch(&llmodel.FetchParam{Mid: up.AiModel})
	}

	// 群聊默认配置
	rc, _ := chatroom.Fetch(&chatroom.FetchParam{Roomid: rid})
	if rc.Rd > 0 {
		if llmc.Rd == 0 && len(rc.ModelDefault) > 1 {
			llmc, _ = llmodel.Fetch(&llmodel.FetchParam{Mid: rc.ModelDefault})
		}
		if len(rc.ModelContext) > 1 {
			modelContext = rc.ModelContext
		}
		if rc.ModelHistory > 1 {
			modelHistory = rc.ModelHistory
		}
	}

	// 全局默认模型
	if llmc.Rd == 0 {
		llmc, _ = llmodel.Fetch(&llmodel.FetchParam{Mid: setting.ModelDefault})
	}

	// 从数据库取第一个
	if llmc.Rd == 0 {
		llmc, _ = llmodel.Fetch(&llmodel.FetchParam{})
	}

	return &UserLLModel{modelContext, modelHistory, llmc}

}

// 历史消息

type MsgHistory struct {
	Content string
	Role    string // user,assistant
}

var msgHistories = map[string][]*MsgHistory{}

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

	if len(msgHistories[id]) >= setting.ModelHistory {
		msgHistories[id] = msgHistories[id][len(items):]
	}

	msgHistories[id] = append(msgHistories[id], items...)

}
