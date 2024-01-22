package robot

import (
	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/wcferry"
)

func roomHandler() {

	if len(args.Usr.ChatRoom) == 0 {
		return
	}

	for _, v := range args.Bot.InvitableRoom {
		v := v // copy
		room := args.GetChatRoom(v)
		cmdkey := "/g:" + room.Argot
		handlers[cmdkey] = &Handler{
			Level:    0,
			ChatAble: true,
			RoomAble: false,
			Describe: "加入群聊 " + room.Name,
			Callback: func(msg *wcferry.WxMsg) string {
				resp := wc.CmdClient.InviteChatroomMembers(v, msg.Sender)
				if resp == 1 {
					return "已发送群邀请，稍后请点击进入"
				} else {
					return "发送群邀请失败"
				}
			},
		}
	}

}
