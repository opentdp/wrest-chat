package robot

import (
	"github.com/opentdp/wrest-chat/dbase/chatroom"
	"github.com/opentdp/wrest-chat/wcferry"
)

func roomHandler() []*Handler {

	cmds := []*Handler{}

	rooms, err := chatroom.FetchAll(&chatroom.FetchAllParam{})

	if err == nil && len(rooms) > 0 {
		for k, v := range rooms {
			if len(v.JoinArgot) < 2 {
				continue
			}
			v := v // copy
			cmdkey := "/jr:" + v.JoinArgot
			cmds = append(cmds, &Handler{
				Level:    -1,
				Order:    510 + int32(k),
				Roomid:   "-",
				Command:  cmdkey,
				Describe: v.Name,
				Callback: func(msg *wcferry.WxMsg) string {
					resp := wc.CmdClient.InviteChatroomMembers(v.Roomid, msg.Sender)
					if resp == 1 {
						return "已发送群邀请，稍后请点击进入"
					} else {
						return "发送群邀请失败"
					}
				},
			})
		}
	}

	return cmds

}
