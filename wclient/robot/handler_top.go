package robot

import (
	"fmt"
	"strings"

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
			text := []string{}
			// 聊天统计
			text = append(text, "", "今日灌水排行", "----------------")
			for _, v := range message.TalkTop10(msg.Roomid) {
				u := wc.CmdClient.GetAliasInChatRoom(v.Sender, msg.Roomid)
				text = append(text, fmt.Sprintf("%s:   %d 次", u, v.RecordCount))
			}
			// 图片统计
			text = append(text, "", "今日斗图排行", "----------------")
			for _, v := range message.ImageTop10(msg.Roomid) {
				u := wc.CmdClient.GetAliasInChatRoom(v.Sender, msg.Roomid)
				text = append(text, fmt.Sprintf("%s:   %d 张", u, v.RecordCount))
			}
			return strings.Join(text, "\n")
		},
	}

}
