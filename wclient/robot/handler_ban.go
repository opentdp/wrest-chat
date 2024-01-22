package robot

import (
	"encoding/xml"
	"strings"

	"github.com/opentdp/wechat-rest/args"
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
				room := args.GetChatRoom(msg.Roomid)
				list := strings.Split(ret.AtUserList, ",")
				for _, v := range list {
					if v == "" {
						continue
					}
					// 群内禁止
					if room.GetMember(v).Level == 9 {
						return "无法操作管理员"
					}
					// 全局禁止
					user := args.GetMember(v)
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
					user := args.GetMember(v)
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
