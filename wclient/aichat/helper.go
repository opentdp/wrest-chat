package aichat

import (
	"github.com/rehiy/one-llm/aichat"

	"github.com/opentdp/wrest-chat/dbase/chatroom"
	"github.com/opentdp/wrest-chat/dbase/llmodel"
	"github.com/opentdp/wrest-chat/dbase/profile"
	"github.com/opentdp/wrest-chat/dbase/setting"
	"github.com/opentdp/wrest-chat/dbase/tables"
)

func Text(msg, id, rid string) string {

	llmc := UserConfig(id, rid)
	return aichat.Text(msg, llmc)

}

func Vison(msg, img, id, rid string) string {

	llmc := UserConfig(id, rid)
	return aichat.Vison(msg, img, llmc)

}

// 用户配置信息

var UserConfigs = map[string]*aichat.UserConfig{}

func UserReset(id, rid string) {

	if id == "" {
		UserConfigs = map[string]*aichat.UserConfig{}
	} else {
		delete(UserConfigs, id+"/"+rid)
		UserConfig(id, rid)
	}

}

func UserConfig(id, rid string) *aichat.UserConfig {

	k := id + "/" + rid
	if v, ok := UserConfigs[k]; ok {
		return v
	}

	// 获取默认配置
	llmc := &tables.LLModel{}
	roleContext := setting.ModelContext
	msgHistoryMax := setting.ModelHistory

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
			roleContext = rc.ModelContext
		}
		if rc.ModelHistory > 1 {
			msgHistoryMax = rc.ModelHistory
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

	// 缓存用户配置
	UserConfigs[k] = &aichat.UserConfig{
		Family:        llmc.Family,
		Provider:      llmc.Provider,
		Model:         llmc.Model,
		Secret:        llmc.Secret,
		Endpoint:      llmc.Endpoint,
		RoleContext:   roleContext,
		MsgHistorys:   []*aichat.MsgHistory{},
		MsgHistoryMax: msgHistoryMax,
	}

	return UserConfigs[k]

}
