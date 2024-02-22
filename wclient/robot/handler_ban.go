package robot

import (
	"encoding/xml"
	"strings"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/dbase/profile"
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wcferry/types"
)

func banHandler() {

	handlers["/ban"] = &Handler{
		Level:    7,
		ChatAble: false,
		RoomAble: true,
		Describe: "禁止用户使用助手",
		Callback: func(msg *wcferry.WxMsg) string {
			ret := &types.AtMsgSource{}
			err := xml.Unmarshal([]byte(msg.Xml), ret)
			if err == nil && ret.AtUserList != "" {
				list := strings.Split(ret.AtUserList, ",")
				for _, v := range list {
					if v == "" {
						continue
					}
					// 群内禁止
					if profile.Get(v, msg.Roomid).Level == 9 {
						return "无法操作管理员"
					}
					// 全局禁止
					user := profile.Get(v, "")
					if user != nil && user.Level == 9 {
						return "无法操作管理员"
					}
					user.Level = 1
				}
				return "操作成功"
			}
			return "参数错误"
		},
	}

	handlers["/unban"] = &Handler{
		Level:    1,
		ChatAble: true,
		RoomAble: true,
		Describe: "允许用户使用助手",
		Callback: func(msg *wcferry.WxMsg) string {
			ret := &types.AtMsgSource{}
			err := xml.Unmarshal([]byte(msg.Xml), ret)
			if err == nil && ret.AtUserList != "" {
				list := strings.Split(ret.AtUserList, ",")
				for _, v := range list {
					if v == "" {
						continue
					}
					user := profile.Get(v, "")
					if user.Level == 1 {
						user.Level = 0
					}
				}
				return "操作成功"
			}
			return "参数错误"
		},
	}

}

func banMessagePrefix(msg *wcferry.WxMsg) string {

	// 全局权限
	user := profile.Get(msg.Sender, "")
	if args.Bot.WhiteMember {
		if user.Level <= 1 {
			msg.Content = ""
			return ""
		}
	} else {
		if user.Level == 1 {
			msg.Content = ""
			return ""
		}
	}

	// 群聊权限
	if msg.IsGroup {
		room := profile.Get("", msg.Roomid)
		user := profile.Get(msg.Sender, msg.Roomid)
		if args.Bot.WhiteChatRoom {
			if room.Level <= 1 {
				msg.Content = ""
				return ""
			}
		} else {
			if room.Level == 1 {
				msg.Content = ""
				return ""
			}
			if user.Level == 1 {
				msg.Content = ""
				return ""
			}
		}
	}

	return ""

}
