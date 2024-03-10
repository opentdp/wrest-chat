package robot

import (
	"github.com/opentdp/wechat-rest/dbase/chatroom"
	"github.com/opentdp/wechat-rest/wcferry"
)

func roomHandler() []*Handler {

	cmds := []*Handler{}

	rooms, err := chatroom.FetchAll(&chatroom.FetchAllParam{})
	if err != nil {
		return cmds
	}

	for k, v := range rooms {
		if len(v.JoinArgot) < 2 {
			continue
		}
		v := v // copy
		cmdkey := "/jr:" + v.JoinArgot
		cmds = append(cmds, &Handler{
			Level:    0,
			Order:    400 + int32(k),
			ChatAble: true,
			RoomAble: false,
			Command:  cmdkey,
			Describe: "加群聊：" + v.Name,
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

	return cmds

}
