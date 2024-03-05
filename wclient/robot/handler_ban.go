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
		Describe: "拉黑指定的用户",
		Callback: func(msg *wcferry.WxMsg) string {
			ret := &types.AtMsgSource{}
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
					// 权限检查
					up, _ := profile.Fetch(&profile.FetchParam{Wxid: v, Roomid: prid(msg)})
					if up.Level >= 7 {
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
	}

	handlers["/unban"] = &Handler{
		Level:    7,
		Order:    41,
		ChatAble: false,
		RoomAble: true,
		Describe: "解封拉黑的用户",
		Callback: func(msg *wcferry.WxMsg) string {
			ret := &types.AtMsgSource{}
			err := xml.Unmarshal([]byte(msg.Xml), ret)
			if err == nil && ret.AtUserList != "" {
				users := strings.Split(ret.AtUserList, ",")
				for _, v := range users {
					if v == "" {
						continue
					}
					// 权限检查
					up, _ := profile.Fetch(&profile.FetchParam{Wxid: v, Roomid: prid(msg)})
					if up.Level >= 7 {
						return "禁止操作管理员"
					}
					// 解封用户
					profile.Replace(&profile.ReplaceParam{Wxid: v, Roomid: prid(msg), BanExpire: -1})
				}
				return "已解封用户"
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

	up, _ := profile.Fetch(&profile.FetchParam{Wxid: msg.Sender, Roomid: prid(msg)})
	if up.Level == 1 || up.BanExpire > time.Now().Unix() {
		msg.Content = ""
	}

	return ""

}
