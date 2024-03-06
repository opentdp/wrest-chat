package robot

import (
	"fmt"

	"github.com/opentdp/wechat-rest/dbase/message"
	"github.com/opentdp/wechat-rest/wcferry"
)

func topHandler() {

	handlers["/top"] = &Handler{
		Level:    7,
		Order:    50,
		ChatAble: false,
		RoomAble: true,
		Describe: "获取群聊统计信息",
		Callback: func(msg *wcferry.WxMsg) string {
			res := ""
			// 聊天统计
			if items := message.TalkTop10(msg.Roomid); len(items) > 0 {
				res += "\n今日水王\n----------------\n"
				for _, v := range items {
					u := wc.CmdClient.GetAliasInChatRoom(v.Sender, msg.Roomid)
					res += fmt.Sprintf("%s:   %d 次\n", u, v.RecordCount)
				}
			}
			// 图片统计
			if items := message.ImageTop10(msg.Roomid); len(items) > 0 {
				res += "\n今日图王\n----------------\n"
				for _, v := range items {
					u := wc.CmdClient.GetAliasInChatRoom(v.Sender, msg.Roomid)
					res += fmt.Sprintf("%s:   %d 次\n", u, v.RecordCount)
				}
			}
			return res
		},
	}

}
