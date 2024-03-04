package robot

import (
	"strings"

	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/wechat-rest/dbase/setting"
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wclient"
)

var wc *wcferry.Client

var clientId string
var selfInfo *wcferry.UserInfo

func Start() {

	setting.Laod()

	if !setting.BotEnable {
		logman.Warn("robot disabled")
		return
	}

	if clientId != "" {
		logman.Warn("robot already started")
		return
	}

	wc = wclient.Register()

	clientId, _ = wc.EnrollReceiver(true, receiver)

}

func Stop() {

	wc.DisableReceiver(clientId)

	clearHandlers()
	selfInfo = nil
	clientId = ""

}

func Redo() {

	clearHandlers()
	selfInfo = nil

}

// 个人信息
func self() *wcferry.UserInfo {

	if selfInfo == nil {
		selfInfo = wc.CmdClient.GetSelfInfo()
	}
	return selfInfo

}

// 会话场景
func prid(msg *wcferry.WxMsg) string {

	if msg.IsGroup {
		return msg.Roomid
	}
	return "-"

}

// 回复消息
func reply(msg *wcferry.WxMsg, text string) int32 {

	if msg.IsSelf {
		return -2
	}

	if text = strings.TrimSpace(text); text == "" {
		return -1
	}

	if msg.IsGroup {
		user := wc.CmdClient.GetInfoByWxid(msg.Sender)
		return wc.CmdClient.SendTxt("@"+user.Name+"\n"+text, msg.Roomid, msg.Sender)
	} else {
		return wc.CmdClient.SendTxt(text, msg.Sender, "")
	}

}
