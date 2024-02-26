package robot

import (
	"fmt"
	"strings"

	"github.com/opentdp/wechat-rest/dbase/keyword"
	"github.com/opentdp/wechat-rest/dbase/profile"
	"github.com/opentdp/wechat-rest/dbase/tables"
	"github.com/opentdp/wechat-rest/wcferry"
)

var badMember = map[string]int{}
var keywordList = []*tables.Keyword{}

func badHandler() {

	handlers["/bad"] = &Handler{
		Level:    7,
		Order:    30,
		ChatAble: true,
		RoomAble: true,
		Describe: "添加违规关键词",
		Callback: func(msg *wcferry.WxMsg) string {
			v := msg.Content
			_, err := keyword.Create(&keyword.CreateParam{
				Roomid: msg.Roomid,
				Phrase: v,
				Level:  1,
			})
			if err == nil {
				updateBadWord()
				return "添加成功"
			}
			return "关键词已存在"
		},
		PreCheck: badMessagePrefix,
	}

	handlers["/unbad"] = &Handler{
		Level:    7,
		Order:    31,
		ChatAble: true,
		RoomAble: true,
		Describe: "删除违规关键词",
		Callback: func(msg *wcferry.WxMsg) string {
			v := msg.Content
			err := keyword.Delete(&keyword.DeleteParam{
				Roomid: msg.Roomid,
				Phrase: v,
			})
			if err == nil {
				updateBadWord()
				return "删除成功"
			}
			return "关键词不存在"
		},
	}

}

func updateBadWord() {

	list, _ := keyword.FetchAll(&keyword.FetchAllParam{})
	keywordList = list

}

func badMessagePrefix(msg *wcferry.WxMsg) string {

	up, _ := profile.Fetch(&profile.FetchParam{Wxid: msg.Sender, Roomid: msg.Roomid})
	if !msg.IsGroup || up.Level >= 7 {
		return ""
	}

	for _, v := range keywordList {
		if v.Level > 0 && strings.Contains(msg.Content, v.Phrase) {
			badMember[msg.Sender] += int(v.Level)
			if badMember[msg.Sender] > 10 {
				wc.CmdClient.DelChatRoomMembers(msg.Roomid, msg.Sender)
				delete(badMember, msg.Sender)
				return "送你离开，天涯之外"
			}
			str := "违规风险 %d，当前累计：%d，大于 10 将赠与免费机票。"
			return fmt.Sprintf(str, v.Level, badMember[msg.Sender])
		}
	}

	return ""

}
