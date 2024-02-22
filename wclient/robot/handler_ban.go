package robot

import (
	"encoding/xml"
	"strings"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/dbase/chatroom"
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
					// 全局禁止
					p, _ := profile.Fetch(&profile.FetchParam{Wxid: msg.Sender})
					if p.Level == 9 {
						return "无法操作管理员"
					}
					// 群内禁止
					p2, _ := profile.Fetch(&profile.FetchParam{Wxid: msg.Sender, Roomid: msg.Roomid})
					if p2.Level == 9 {
						return "无法操作管理员"
					}
					// 禁止使用
					profile.Update(&profile.UpdateParam{Wxid: v, Level: 1})
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
					p, _ := profile.Fetch(&profile.FetchParam{Wxid: v})
					if p.Level == 1 {
						p.Level = 0
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
	user, _ := profile.Fetch(&profile.FetchParam{Wxid: msg.Sender})
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
		room, _ := chatroom.Fetch(&chatroom.FetchParam{Roomid: msg.Roomid})
		user, _ := profile.Fetch(&profile.FetchParam{Wxid: msg.Sender, Roomid: msg.Roomid})
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
