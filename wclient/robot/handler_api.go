package robot

import (
	"strings"

	"github.com/opentdp/go-helper/request"
	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/wcferry"
)

func apiHandler() {

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

	str := strings.Replace(strings.TrimSpace(msg.Content), " ", "/", 1)
	url := "https://api.rehi.org/format=yaml/" + str

	// 获取指令名称

	parts := strings.SplitN(str, "/", 2)
	cmd := parts[0]

	// 返回卡片消息

	if strings.Contains("news,port,iptv,weather", cmd) {
		digest := "请点击卡片查看结果"
		icon := "https://api.rehi.org/assets/icon.png"
		if msg.IsGroup {
			wc.CmdClient.SendRichText("小秘书", "mphelper", msg.Content, digest, url, icon, msg.Roomid)
		} else {
			wc.CmdClient.SendRichText("小秘书", "mphelper", msg.Content, digest, url, icon, msg.Sender)
		}
		return ""
	}

	// 获取结果后返回

	res, err := request.TextGet(url, request.H{
		"User-Agent": args.AppName + "/" + args.Version,
	})

	if err != nil {
		return err.Error()
	}

	if cmd == "" || cmd == "help" {
		lines := strings.Split(res, "\n")
		for k, line := range lines {
			line = strings.TrimLeft(line, "/")
			line = strings.Replace(line, "/", " ", 1)
			lines[k] = "/api " + strings.TrimSpace(line)
		}
		return strings.Join(lines, "\n")
	}

	if fileReply(msg, res) == 0 {
		return ""
	}

	return res

}
