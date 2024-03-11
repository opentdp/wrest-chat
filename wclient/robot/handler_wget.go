package robot

import (
	"strings"

	"github.com/opentdp/wechat-rest/wcferry"
)

func wgetHandler() []*Handler {

	cmds := []*Handler{}

	cmds = append(cmds, &Handler{
		Level:    7,
		Order:    90,
		ChatAble: true,
		RoomAble: true,
		Describe: "获取图片或文件",
		Callback: func(msg *wcferry.WxMsg) string {
			u := msg.Content
			if !strings.HasPrefix(u, "http") {
				return "请输入正确的网址"
			}
			if wc.CmdClient.SendFlexMsg(u, msg.Sender, msg.Roomid) != 0 {
				return "文件获取失败"
			}
			return ""
		},
	})

	return cmds

}
