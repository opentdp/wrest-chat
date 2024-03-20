package robot

import (
	"github.com/opentdp/wrest-chat/wcferry"
)

func revokeHandler() []*Handler {

	cmds := []*Handler{}

	cmds = append(cmds, &Handler{
		Level:    7,
		Order:    390,
		Roomid:   "*",
		Command:  "/revoke",
		Describe: "撤回引用的机器人发言",
		Callback: func(msg *wcferry.WxMsg) string {
			if msg.Sign == "refer-msg" {
				wc.CmdClient.RevokeMsg(msg.Id)
			}
			return ""
		},
	})

	return cmds

}
