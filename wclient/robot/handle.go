package robot

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wclient/model"
)

type Handler struct {
	Level    int    // 0:不限制 9:管理员
	ChatAble bool   // 是否允许在私聊使用
	RoomAble bool   // 是否允许在群聊使用
	Describe string // 指令描述
	Callback func(msg *wcferry.WxMsg) string
}

var handlers = make(map[string]*Handler)

func initHandlers() {

	helper1 := "" // 私聊指令
	helper2 := "" // 群聊指令

	handlers["/ai"] = &Handler{
		Level:    0,
		ChatAble: true,
		RoomAble: true,
		Describe: "提问或交谈",
		Callback: func(msg *wcferry.WxMsg) string {
			return model.AiChat(msg.Sender, msg.Content)
		},
	}

	handlers["/new"] = &Handler{
		Level:    0,
		ChatAble: true,
		RoomAble: true,
		Describe: "重置上下文内容",
		Callback: func(msg *wcferry.WxMsg) string {
			return model.ClearHistory(msg.Sender)
		},
	}

	handlers["/mr"] = &Handler{
		Level:    0,
		ChatAble: true,
		RoomAble: true,
		Describe: "随机选择一个模型",
		Callback: func(msg *wcferry.WxMsg) string {
			k := rand.Intn(len(args.LLM.Models))
			return model.SetUserModel(msg.Sender, args.LLM.Models[k])
		},
	}

	for _, v := range args.LLM.Models {
		v := v // copy it
		cmdkey := "/m:" + v.Name
		handlers[cmdkey] = &Handler{
			Level:    0,
			ChatAble: true,
			RoomAble: true,
			Describe: "切换为 " + v.Model + " 模型",
			Callback: func(msg *wcferry.WxMsg) string {
				return model.SetUserModel(msg.Sender, v)
			},
		}
	}

	for _, v := range args.Bot.InvitableRooms {
		v := v // copy it
		cmdkey := "/room:" + v.Mark
		handlers[cmdkey] = &Handler{
			Level:    0,
			ChatAble: true,
			RoomAble: false,
			Describe: "加入群聊 " + v.Name,
			Callback: func(msg *wcferry.WxMsg) string {
				resp := wc.CmdClient.InviteChatroomMembers(v.RoomId, msg.Sender)
				if resp == 1 {
					return "已发送群邀请，稍后请点击进入"
				} else {
					return "发送群邀请失败"
				}
			},
		}
	}

	handlers["/help"] = &Handler{
		Level:    0,
		ChatAble: true,
		RoomAble: true,
		Describe: "查看帮助信息",
		Callback: func(msg *wcferry.WxMsg) string {
			text := ""
			if msg.IsGroup {
				text += helper2
			} else {
				text += helper1
			}
			text += "对话模型 " + model.GetUserModel(msg.Sender).Name + "，"
			text += fmt.Sprintf("上下文长度 %d/%d", model.CountHistory(msg.Sender), args.LLM.HistoryNum)
			return text
		},
	}

	for k, v := range handlers {
		if v.ChatAble {
			helper1 += k + " " + v.Describe + "\n"
		}
		if v.RoomAble {
			helper2 += k + " " + v.Describe + "\n"
		}
	}

}

func chatHandler(msg *wcferry.WxMsg) string {

	// 解析指令
	re := regexp.MustCompile(`^(/[\w:-]{2,20})(\s|$)`)
	matches := re.FindStringSubmatch(msg.Content)
	if matches == nil || len(matches) < 2 {
		return ""
	}

	// 清理指令
	command := matches[1]
	msg.Content = strings.TrimSpace(msg.Content[len(matches[0]):])
	if command == "/ai" && msg.Content == "" {
		command = "/help"
	}

	// 查找指令
	hd, exists := handlers[command]
	if !exists {
		return "指令或参数错误"
	}

	// 检查权限
	if hd.Level > 0 {
		return "无权限使用该指令"
	}

	// 检查场景
	if msg.IsGroup {
		if !hd.RoomAble {
			return "该指令只能在私聊中使用"
		}
	} else {
		if !hd.ChatAble {
			return "该指令只能在群聊中使用"
		}
	}

	// 执行指令
	return hd.Callback(msg)

}
