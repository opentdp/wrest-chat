package robot

import (
	"github.com/opentdp/wechat-rest/wcferry"
)

func revokeHandler() []*Handler {

	cmds := []*Handler{}

	cmds = append(cmds, &Handler{
		Level:    7,
		Order:    800,
		ChatAble: true,
		RoomAble: true,
		Command:  "/revoke",
		Describe: "撤回引用的消息（机器人发送的）",
		Callback: func(msg *wcferry.WxMsg) string {
			if msg.Extra == "refer-msg" {
				wc.CmdClient.RevokeMsg(msg.Id)
			}
			return ""
		},
	})

	return cmds

}
