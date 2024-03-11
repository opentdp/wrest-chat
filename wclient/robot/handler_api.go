package robot

import (
	"net/url"
	"strings"

	"github.com/opentdp/go-helper/request"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/dbase/setting"
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wclient"
)

func apiHandler() []*Handler {

	cmds := []*Handler{}

	if len(setting.ApiEndpoint) < 10 {
		return cmds
	}

	cmds = append(cmds, &Handler{
		Level:    0,
		Order:    200,
		ChatAble: true,
		RoomAble: true,
		Command:  "/api",
		Describe: "调用远程接口",
		Callback: apiCallback,
	})

	return cmds

}

func apiCallback(msg *wcferry.WxMsg) string {

	self := wc.CmdClient.GetSelfInfo()

	cmd := []string{"help"}
	if msg.Content != "" {
		cmd = strings.SplitN(msg.Content, " ", 2)
		if len(cmd) > 1 {
			cmd[1] = url.QueryEscape(cmd[1])
		}
	}

	// 获取结果
	url := setting.ApiEndpoint + strings.Join(cmd, "/")
	res, err := request.TextGet(url, request.H{
		"Client-Uid": self.Wxid + "," + msg.Sender,
		"User-Agent": args.AppName + "/" + args.Version,
	})
	if err != nil {
		return err.Error()
	}

	// 返回卡片消息
	if strings.Count(res, "\n") > 20 || len(res) > 900 {
		receiver := msg.Sender
		if msg.IsGroup {
			receiver = msg.Roomid
		}
		title := msg.Content
		digest := "请点击卡片查看结果"
		icon := setting.ApiEndpointIcon
		wc.CmdClient.SendRichText(self.Name, self.Wxid, title, digest, url, icon, receiver)
		return ""
	}

	// 尝试发送文件
	if wclient.SendFlexMsg(res, msg.Sender, msg.Roomid) == 0 {
		return ""
	}

	return res

}
