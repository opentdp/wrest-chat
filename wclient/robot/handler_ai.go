package robot

import (
	"strings"

	"github.com/opentdp/wechat-rest/dbase/llmodel"
	"github.com/opentdp/wechat-rest/dbase/profile"
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wclient/aichat"
)

func aiHandler() {

	models, err := llmodel.FetchAll(&llmodel.FetchAllParam{})
	if err != nil || len(models) == 0 {
		return
	}

	handlers["/ai"] = &Handler{
		Level:    0,
		Order:    10,
		ChatAble: true,
		RoomAble: true,
		Describe: "提问或交谈",
		Callback: func(msg *wcferry.WxMsg) string {
			text := strings.TrimSpace(msg.Content)
			if text != "" {
				return aichat.Text(msg.Sender, msg.Roomid, text)
			}
			return "请在指令后输入问题"
		},
		PreCheck: aiPreCheck,
	}

	handlers["/new"] = &Handler{
		Level:    0,
		Order:    11,
		ChatAble: true,
		RoomAble: true,
		Describe: "重置上下文内容",
		Callback: func(msg *wcferry.WxMsg) string {
			aichat.ResetHistory(msg.Sender)
			return "已重置上下文"
		},
	}

	for _, v := range models {
		v := v // copy
		cmdkey := "/m:" + v.Mid
		handlers[cmdkey] = &Handler{
			Level:    0,
			Order:    12,
			ChatAble: true,
			RoomAble: true,
			Describe: "切换为 " + v.Family + " [" + v.Model + "]",
			Callback: func(msg *wcferry.WxMsg) string {
				profile.Migrate(&profile.MigrateParam{Wxid: msg.Sender, Roomid: msg.Roomid, AiModel: v.Mid})
				return "对话模型切换为 " + v.Family + " [" + v.Model + "]"
			},
		}
	}

	handlers["/mr"] = &Handler{
		Level:    0,
		Order:    13,
		ChatAble: true,
		RoomAble: true,
		Describe: "随机选择模型",
		Callback: func(msg *wcferry.WxMsg) string {
			for _, v := range models {
				profile.Migrate(&profile.MigrateParam{Wxid: msg.Sender, Roomid: msg.Roomid, AiModel: v.Mid})
				return "对话模型切换为 " + v.Family + " [" + v.Model + "]"
			}
			return "没有可用的模型"
		},
	}

	handlers["/wake"] = &Handler{
		Level:    0,
		Order:    14,
		ChatAble: true,
		RoomAble: true,
		Describe: "自定义唤醒词",
		Callback: func(msg *wcferry.WxMsg) string {
			argot := strings.TrimSpace(msg.Content)
			// 校验唤醒词
			if strings.Contains(argot, "@") || strings.Contains(argot, "/") {
				return "唤醒词不允许包含 @ 或 /"
			} else if argot == "" {
				argot = "-"
			}
			// 更新唤醒词
			profile.Migrate(&profile.MigrateParam{Wxid: msg.Sender, Roomid: msg.Roomid, AiArgot: argot})
			if argot == "-" {
				if msg.IsGroup {
					return "已禁用自定义唤醒词"
				}
				return "已启用无唤醒词对话模式"
			}
			return "唤醒词设置为 " + argot
		},
	}

}

func aiPreCheck(msg *wcferry.WxMsg) string {

	if len(msg.Content) == 0 {
		return ""
	}

	if msg.Content[0:1] != "/" {
		// 处理 @机器人 的消息
		if strings.Contains(msg.Xml, self().Wxid) {
			msg.Content = "/ai " + msg.Content
			return ""
		}
		// 处理用户自定义的唤醒词
		up, _ := profile.Fetch(&profile.FetchParam{Wxid: msg.Sender, Roomid: msg.Roomid})
		if up.AiArgot == "-" {
			if !msg.IsGroup {
				msg.Content = "/ai " + msg.Content
			}
		} else if up.AiArgot != "" {
			if strings.HasPrefix(msg.Content, up.AiArgot) {
				msg.Content = strings.Replace(msg.Content, up.AiArgot, "/ai ", 1)
			}
		}
	}

	return ""

}
