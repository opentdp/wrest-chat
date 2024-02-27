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

	cmd := []string{"help"}
	if msg.Content != "" {
		cmd = strings.SplitN(msg.Content, " ", 2)
	}

	url := "https://api.rehi.org/format=yaml/" + strings.Join(cmd, "/")

	// 返回卡片消息
	cards := "news,port,iptv,weather"
	if strings.Contains(cards, cmd[0]) {
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

	// 处理帮助信息
	if cmd[0] == "help" {
		lines := strings.Split(res, "\n")
		for k, line := range lines {
			line = strings.TrimLeft(line, "/")
			line = strings.Replace(line, "/", " ", 1)
			lines[k] = strings.TrimSpace(line)
		}
		return "/api " + strings.Join(cmd, " ") + "\n" + strings.Join(lines, "\n")
	}

	// 尝试发送文件
	if fileReply(msg, res) == 0 {
		return ""
	}

	return res

}
