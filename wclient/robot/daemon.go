package robot

import (
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

	initHandlers()

	wc = wclient.Register()
	clientId, _ = wc.EnrollReceiver(true, receiver)

}

func Redo() {

	selfInfo = nil
	initHandlers()

}
