package robot

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/opentdp/wechat-rest/dbase/chatroom"
	"github.com/opentdp/wechat-rest/dbase/profile"
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wcferry/types"
)

func banHandler() {

	handlers["/ban"] = &Handler{
		Level:    7,
		Order:    40,
		ChatAble: false,
		RoomAble: true,
		Describe: "禁止用户使用助手",
		Callback: func(msg *wcferry.WxMsg) string {
			ret := &types.AtMsgSource{}
			err := xml.Unmarshal([]byte(msg.Xml), ret)
			if err == nil && ret.AtUserList != "" {
				// 获取禁言时限
				parts := strings.Split(msg.Content, " ")
				second, err := strconv.Atoi(parts[0])
				if err != nil {
					second = 86400
				}
				// 批量操作禁言
				list := strings.Split(ret.AtUserList, ",")
				for _, v := range list {
					if v == "" {
						continue
					}
					// 权限检查
					up, _ := profile.Fetch(&profile.FetchParam{Wxid: v, Roomid: msg.Roomid})
					if up.Level >= 7 {
						return "禁止操作管理员"
					}
					// 禁止使用
					expire := time.Now().Unix() + int64(second)
					profile.Migrate(&profile.UpdateParam{Wxid: v, Roomid: msg.Roomid, BanExpire: expire})
				}
				return fmt.Sprintf("操作成功，有效期 %d 秒", second)
			}
			return "参数错误"
		},
		PreCheck: banPreCheck,
	}

	handlers["/unban"] = &Handler{
		Level:    7,
		Order:    41,
		ChatAble: false,
		RoomAble: true,
		Describe: "取消使用助手限制",
		Callback: func(msg *wcferry.WxMsg) string {
			ret := &types.AtMsgSource{}
			err := xml.Unmarshal([]byte(msg.Xml), ret)
			if err == nil && ret.AtUserList != "" {
				list := strings.Split(ret.AtUserList, ",")
				for _, v := range list {
					if v == "" {
						continue
					}
					// 权限检查
					up, _ := profile.Fetch(&profile.FetchParam{Wxid: v, Roomid: msg.Roomid})
					if up.Level >= 7 {
						return "禁止操作管理员"
					}
					// 取消禁止
					profile.Migrate(&profile.UpdateParam{Wxid: v, Roomid: msg.Roomid, BanExpire: -1})
				}
				return "已取消限制"
			}
			return "参数错误"
		},
	}

}

func banPreCheck(msg *wcferry.WxMsg) string {

	if msg.IsGroup {
		room, _ := chatroom.Fetch(&chatroom.FetchParam{Roomid: msg.Roomid})
		if room.Level == 1 {
			msg.Content = ""
			return ""
		}
	}

	up, _ := profile.Fetch(&profile.FetchParam{Wxid: msg.Sender, Roomid: msg.Roomid})
	if up.Level == 1 || up.BanExpire > time.Now().Unix() {
		msg.Content = ""
	}

	return ""

}
