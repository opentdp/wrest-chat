package robot

import (
	"strings"

	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wclient"
)

func wgetHandler() []*Handler {

	cmds := []*Handler{}

	cmds = append(cmds, &Handler{
		Level:    7,
		Order:    380,
		Roomid:   "*",
		Command:  "/wget",
		Describe: "获取图片或文件",
		Callback: func(msg *wcferry.WxMsg) string {
			u := msg.Content
			if !strings.HasPrefix(u, "http") {
				return "请输入正确的网址"
			}
			if wclient.SendFlexMsg(u, msg.Sender, msg.Roomid) != 0 {
				return "文件获取失败"
			}
			return ""
		},
	})

	return cmds

}
