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
var badwordList = []*tables.Keyword{}

var roomMemberAlias = map[string]string{}

func badHandler() []*Handler {

	updateBadWord()

	cmds := []*Handler{}

	cmds = append(cmds, &Handler{
		Level:    7,
		Order:    310,
		Roomid:   "*",
		Command:  "/bad",
		Describe: "添加违禁词",
		Callback: func(msg *wcferry.WxMsg) string {
			_, err := keyword.Create(&keyword.CreateParam{
				Group: "badword", Roomid: prid(msg), Phrase: msg.Content, Level: 1,
			})
			if err == nil {
				updateBadWord()
				return "违禁词添加成功"
			}
			return "违禁词已存在"
		},
		PreCheck: badPreCheck,
	})

	cmds = append(cmds, &Handler{
		Level:    7,
		Order:    311,
		Roomid:   "*",
		Command:  "/unbad",
		Describe: "删除违禁词",
		Callback: func(msg *wcferry.WxMsg) string {
			err := keyword.Delete(&keyword.DeleteParam{
				Group: "badword", Roomid: prid(msg), Phrase: msg.Content,
			})
			if err == nil {
				updateBadWord()
				return "违禁词删除成功"
			}
			return "违禁词删除失败"
		},
	})

	return cmds

}

func badPreCheck(msg *wcferry.WxMsg) string {

	// 私聊豁免
	if !msg.IsGroup {
		return ""
	}

	// 管理豁免
	up, _ := profile.Fetch(&profile.FetchParam{Wxid: msg.Sender, Roomid: prid(msg)})
	if up.Level >= 7 {
		return ""
	}

	// 遍历关键词
	text := roomMemberName(msg.Sender, msg.Roomid) + msg.Content
	for _, v := range badwordList {
		if msg.IsGroup {
			if v.Roomid != "*" && v.Roomid != "+" && v.Roomid != msg.Roomid {
				continue // 忽略
			}
		} else {
			if v.Roomid != "*" && v.Roomid != "-" {
				continue // 忽略
			}
		}
		if v.Level > 0 && strings.Contains(text, v.Phrase) {
			badMember[msg.Sender] += int(v.Level)
			if badMember[msg.Sender] > 10 {
				defer delete(badMember, msg.Sender)
				defer wc.CmdClient.DelChatRoomMembers(msg.Roomid, msg.Sender)
				return "我送你离开，天涯之外你是否还在"
			}
			str := "违规风险 +%d，当前累计：%d，大于 10 将被请出群聊"
			return fmt.Sprintf(str, v.Level, badMember[msg.Sender])
		}
	}

	return ""

}

func roomMemberName(wxid, roomid string) string {

	k := fmt.Sprintf("%s@%s", wxid, roomid)

	if roomMemberAlias[k] == "" {
		roomMemberAlias[k] = wc.CmdClient.GetAliasInChatRoom(wxid, roomid)
	}
	return roomMemberAlias[k]

}

func updateBadWord() {

	badwordList, _ = keyword.FetchAll(&keyword.FetchAllParam{
		Group: "badword",
	})

}
