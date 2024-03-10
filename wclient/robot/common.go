package robot

import (
	"strings"

	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/wechat-rest/dbase/setting"
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wclient"
)

var wc *wcferry.Client
var selfInfo *wcferry.UserInfo

func Start() {

	setting.Laod()

	if !setting.BotEnable {
		logman.Warn("robot disabled")
		return
	}

	if wc != nil {
		logman.Warn("robot already started")
		return
	}

	initHandlers()

	wc = wclient.Register()
	_, err := wc.EnrollReceiver(true, receiver)
	if err != nil {
		logman.Fatal("robot start failed", "error", err)
	}

}

func Redo() {

	selfInfo = nil
	initHandlers()

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
		if user != nil && user.Name != "" {
			text = "@" + user.Name + "\n" + text
		}
		return wc.CmdClient.SendTxt(text, msg.Roomid, msg.Sender)
	}

	return wc.CmdClient.SendTxt(text, msg.Sender, "")

}
