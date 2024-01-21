package robot

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wcferry/types"
)

func banHandler() {

	if len(args.Bot.WhiteList) > 0 {
		return
	}

	handlers["/ban"] = &Handler{
		Level:    1,
		ChatAble: true,
		RoomAble: true,
		Describe: "禁止用户使用助手",
		Callback: func(msg *wcferry.WxMsg) string {
			ret := &types.AtMsgSource{}
			err := xml.Unmarshal([]byte(msg.Xml), ret)
			if err == nil && ret.AtUserList != "" {
				users := strings.Split(ret.AtUserList, ",")
				for _, v := range users {
					if v != "" && !sliceContains(args.Bot.BlackList, v) {
						args.Bot.BlackList = append(args.Bot.BlackList, v)
					}
				}
				if err := args.Co.SaveYaml(); err != nil {
					return fmt.Sprintf("写入配置错误：%s", err)
				}
				return fmt.Sprintf("已禁止用户数：%d", len(args.Bot.BlackList))
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
				users := strings.Split(ret.AtUserList, ",")
				for _, v := range users {
					if v != "" {
						args.Bot.BlackList = sliceRemove(args.Bot.BlackList, v)
					}
				}
				if err := args.Co.SaveYaml(); err != nil {
					return fmt.Sprintf("写入配置错误：%s", err)
				}
				return fmt.Sprintf("已禁止用户数：%d", len(args.Bot.BlackList))
			}
			return "参数错误"
		},
	}

}
