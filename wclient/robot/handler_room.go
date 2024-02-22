package robot

import (
	"github.com/opentdp/wechat-rest/dbase/chatroom"
	"github.com/opentdp/wechat-rest/wcferry"
)

func roomHandler() {

	rooms, err := chatroom.FetchAll(&chatroom.FetchAllParam{})
	if err != nil {
		return
	}

	for _, v := range rooms {
		room := v // copy
		cmdkey := "/g:" + room.JoinArgot
		handlers[cmdkey] = &Handler{
			Level:    0,
			ChatAble: true,
			RoomAble: false,
			Describe: "加入群聊 " + room.Name,
			Callback: func(msg *wcferry.WxMsg) string {
				resp := wc.CmdClient.InviteChatroomMembers(room.Roomid, msg.Sender)
				if resp == 1 {
					return "已发送群邀请，稍后请点击进入"
				} else {
					return "发送群邀请失败"
				}
			},
		}
	}

}
