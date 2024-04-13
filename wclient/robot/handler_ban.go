package robot

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/opentdp/wrest-chat/dbase/chatroom"
	"github.com/opentdp/wrest-chat/dbase/profile"
	"github.com/opentdp/wrest-chat/wcferry"
	"github.com/opentdp/wrest-chat/wcferry/types"
)

func banHandler() []*Handler {

	cmds := []*Handler{}

	cmds = append(cmds, &Handler{
		Level:    7,
		Order:    320,
		Roomid:   "+",
		Command:  "/ban",
		Describe: "拉黑指定的用户",
		Callback: func(msg *wcferry.WxMsg) string {
			ret := &types.MsgXmlAtUser{}
			err := xml.Unmarshal([]byte(msg.Xml), ret)
			if err == nil && ret.AtUserList != "" {
				// 获取拉黑时限
				parts := strings.Split(msg.Content, " ")
				second, err := strconv.Atoi(parts[0])
				if err != nil {
					second = 86400
				}
				// 批量操作拉黑
				users := strings.Split(ret.AtUserList, ",")
				for _, v := range users {
					if v == "" {
						continue
					}
					// 管理豁免
					up, _ := profile.Fetch(&profile.FetchParam{Wxid: v, Roomid: prid(msg)})
					if up.Level > 6 {
						return "禁止操作管理员"
					}
					// 拉黑用户
					expire := time.Now().Unix() + int64(second)
					profile.Replace(&profile.ReplaceParam{Wxid: v, Roomid: prid(msg), BanExpire: expire})
				}
				return fmt.Sprintf("已拉黑，有效期 %d 秒", second)
			}
			return "参数错误"
		},
		PreCheck: banPreCheck,
	})

	cmds = append(cmds, &Handler{
		Level:    7,
		Order:    321,
		Roomid:   "+",
		Command:  "/ban:rm",
		Describe: "解封拉黑的用户",
		Callback: func(msg *wcferry.WxMsg) string {
			ret := &types.MsgXmlAtUser{}
			err := xml.Unmarshal([]byte(msg.Xml), ret)
			if err == nil && ret.AtUserList != "" {
				users := strings.Split(ret.AtUserList, ",")
				for _, v := range users {
					if v == "" {
						continue
					}
					// 管理豁免
					up, _ := profile.Fetch(&profile.FetchParam{Wxid: v, Roomid: prid(msg)})
					if up.Level > 6 {
						return "禁止操作管理员"
					}
					// 解封用户
					profile.Replace(&profile.ReplaceParam{Wxid: v, Roomid: prid(msg), BanExpire: -1})
				}
				return "已解封用户"
			}
			return "参数错误"
		},
	})

	return cmds

}

func banPreCheck(msg *wcferry.WxMsg) string {

	// 管理豁免
	up, _ := profile.Fetch(&profile.FetchParam{Wxid: msg.Sender, Roomid: prid(msg)})
	if up.Level > 6 {
		return ""
	}

	// 群聊已拉黑
	if msg.IsGroup {
		room, _ := chatroom.Fetch(&chatroom.FetchParam{Roomid: msg.Roomid})
		if room.Level == 1 {
			return "-"
		}
	}

	// 用户已拉黑
	if up.Level == 1 || up.BanExpire > time.Now().Unix() {
		return "-"
	}

	return ""

}
