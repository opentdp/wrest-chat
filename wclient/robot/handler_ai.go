package robot

import (
	"fmt"

	"github.com/opentdp/wechat-rest/dbase/llmodel"
	"github.com/opentdp/wechat-rest/dbase/profile"
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wclient/aichat"
)

func aiHandler() []*Handler {

	cmds := []*Handler{}

	models, err := llmodel.FetchAll(&llmodel.FetchAllParam{})
	if err != nil || len(models) == 0 {
		return cmds
	}

	cmds = append(cmds, &Handler{
		Level:    0,
		Order:    10,
		ChatAble: true,
		RoomAble: true,
		Command:  "/ai",
		Describe: "提问或交谈",
		Callback: func(msg *wcferry.WxMsg) string {
			if msg.Content == "" {
				return "请在指令后输入问题"
			}
			if msg.Extra == "image-txt" {
				if msg.Thumb == "" {
					return "提取消息图片失败"
				}
				return aichat.Image(msg.Sender, msg.Roomid, msg.Content, msg.Thumb)
			}
			return aichat.Text(msg.Sender, msg.Roomid, msg.Content)
		},
	})

	cmds = append(cmds, &Handler{
		Level:    0,
		Order:    11,
		ChatAble: true,
		RoomAble: true,
		Command:  "/ai:new",
		Describe: "重置上下文内容",
		Callback: func(msg *wcferry.WxMsg) string {
			aichat.ResetHistory(msg.Sender)
			return "已重置上下文"
		},
	})

	if len(models) > 3 {
		cmds = append(cmds, &Handler{
			Level:    0,
			Order:    13,
			ChatAble: true,
			RoomAble: true,
			Command:  "/ai:rand",
			Describe: "随机选择模型",
			Callback: func(msg *wcferry.WxMsg) string {
				up, _ := profile.Fetch(&profile.FetchParam{Wxid: msg.Sender, Roomid: prid(msg)})
				for _, v := range models {
					if v.Level <= up.Level {
						profile.Replace(&profile.ReplaceParam{Wxid: msg.Sender, Roomid: prid(msg), AiModel: v.Mid})
						return "对话模型切换为 " + v.Family + " [" + v.Model + "]"
					}
				}
				return fmt.Sprintf("没有可用的模型（Level ≤ %d）", up.Level)
			},
		})
	}

	for _, v := range models {
		v := v // copy
		cmdkey := "/cm:" + v.Mid
		cmds = append(cmds, &Handler{
			Level:    v.Level,
			Order:    14,
			ChatAble: true,
			RoomAble: true,
			Command:  cmdkey,
			Describe: "换模型：" + v.Family + " [" + v.Model + "]",
			Callback: func(msg *wcferry.WxMsg) string {
				profile.Replace(&profile.ReplaceParam{Wxid: msg.Sender, Roomid: prid(msg), AiModel: v.Mid})
				return "对话模型切换为 " + v.Family + " [" + v.Model + "]"
			},
		})
	}

	return cmds

}
