package robot

import (
	"strings"
	"unicode/utf8"

	"github.com/opentdp/go-helper/request"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/dbase/setting"
	"github.com/opentdp/wechat-rest/wcferry"
)

func apiHandler() {

	if len(setting.ApiEndpoint) < 10 {
		return
	}

	handlers["/api"] = &Handler{
		Level:    0,
		Order:    20,
		ChatAble: true,
		RoomAble: true,
		Describe: "调用远程接口",
		Callback: apiCallback,
	}

}

func apiCallback(msg *wcferry.WxMsg) string {

	cmd := []string{"help"}
	if msg.Content != "" {
		cmd = strings.SplitN(msg.Content, " ", 2)
	}

	// 获取结果
	url := setting.ApiEndpoint + strings.Join(cmd, "/")
	res, err := request.TextGet(url, request.H{
		"User-Agent": args.AppName + "/" + args.Version,
		"Client-Id":  self().Wxid + "," + msg.Sender,
	})
	if err != nil {
		return err.Error()
	}

	// 返回卡片消息
	if utf8.RuneCountInString(res) > 120 {
		receiver := msg.Sender
		if msg.IsGroup {
			receiver = msg.Roomid
		}
		title := msg.Content
		digest := "请点击卡片查看结果"
		icon := setting.ApiEndpointIcon
		wc.CmdClient.SendRichText(self().Name, self().Wxid, title, digest, url, icon, receiver)
		return ""
	}

	// 尝试发送文件
	if wc.CmdClient.SendFlexMsg(res, msg.Sender, msg.Roomid) == 0 {
		return ""
	}

	return res

}
