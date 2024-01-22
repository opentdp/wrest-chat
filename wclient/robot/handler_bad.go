package robot

import (
	"strconv"
	"strings"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/wcferry"
)

var badMember = map[string]int{}

func badHandler() {

	handlers["/bad"] = &Handler{
		Level:    7,
		ChatAble: true,
		RoomAble: true,
		Describe: "添加违规关键词",
		Callback: func(msg *wcferry.WxMsg) string {
			v := msg.Content
			if args.Bot.BadWord[v] == 0 {
				args.Bot.BadWord[v] = 1
				return "添加成功"
			}
			return "关键词已存在"
		},
	}

	handlers["/unbad"] = &Handler{
		Level:    7,
		ChatAble: true,
		RoomAble: true,
		Describe: "添加违规关键词",
		Callback: func(msg *wcferry.WxMsg) string {
			v := msg.Content
			if args.Bot.BadWord[v] > 0 {
				delete(args.Bot.BadWord, v)
				return "删除成功"
			}
			return "关键词不存在"
		},
	}

}

func badMessagePrefix(msg *wcferry.WxMsg) string {

	if !msg.IsGroup || args.GetMember(msg.Sender).Level >= 7 {
		return ""
	}

	for k, v := range args.Bot.BadWord {
		if v > 0 && strings.Contains(msg.Content, k) {
			badMember[msg.Sender] += v
			if badMember[msg.Sender] > 10 {
				wc.CmdClient.DelChatRoomMembers(msg.Roomid, msg.Sender)
			}
			return "违规风险 +" + strconv.Itoa(v) + "，当前累计：" + strconv.Itoa(badMember[msg.Sender]) + "，大于 10 将赠与免费机票。"
		}
	}

	return ""

}
