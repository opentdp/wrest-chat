package robot

import (
	"encoding/xml"
	"fmt"
	"math/rand"
	"regexp"
	"sort"
	"strings"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wclient/model"
	"github.com/opentdp/wechat-rest/wclient/types"
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

	helper1 := []string{} // 私聊指令
	helper2 := []string{} // 群聊指令

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
			model.ResetHistory(msg.Sender)
			return "已重置上下文"
		},
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
					if v != "" && !contains(args.Bot.BlackList, v) {
						args.Bot.BlackList = append(args.Bot.BlackList, v)
					}
				}
				return fmt.Sprintf("已禁止用户数：%d", len(args.Bot.BlackList))
			}
			return "参数错误"
		},
	}

	handlers["/mr"] = &Handler{
		Level:    0,
		ChatAble: true,
		RoomAble: true,
		Describe: "随机选择模型",
		Callback: func(msg *wcferry.WxMsg) string {
			k := rand.Intn(len(args.LLM.Models))
			v := args.LLM.Models[k]
			model.GetUserConfig(msg.Sender).LLModel = v
			return "对话模型切换为 " + v.Name + " [" + v.Model + "]"
		},
	}

	handlers["/wake"] = &Handler{
		Level:    0,
		ChatAble: true,
		RoomAble: false,
		Describe: "设置唤醒词",
		Callback: func(msg *wcferry.WxMsg) string {
			model.GetUserConfig(msg.Sender).WakeWord = msg.Content
			if msg.Content == "" {
				return "唤醒词设置为 " + msg.Content
			}
			return "已禁用唤醒词"
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
				model.GetUserConfig(msg.Sender).LLModel = v
				return "对话模型切换为 " + v.Name + " [" + v.Model + "]"
			},
		}
	}

	for _, v := range args.Bot.HostedRooms {
		v := v // copy it
		cmdkey := "/room:" + v.Mask
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
				text += strings.Join(helper2, "\n") + "\n"
			} else {
				text += strings.Join(helper1, "\n") + "\n"
			}
			text += "对话模型 " + model.GetUserConfig(msg.Sender).LLModel.Name + "，"
			text += fmt.Sprintf("上下文长度 %d/%d", model.CountHistory(msg.Sender), args.LLM.HistoryNum)
			return text
		},
	}

	// 生成帮助信息

	for k, v := range handlers {
		if v.ChatAble {
			helper1 = append(helper1, k+" "+v.Describe)
		}
		if v.RoomAble {
			helper2 = append(helper2, k+" "+v.Describe)
		}
	}

	sort.Strings(helper1)
	sort.Strings(helper2)

}

func applyHandler(msg *wcferry.WxMsg) string {

	// 匹配名单
	if len(args.Bot.WhiteList) > 0 && !contains(args.Bot.WhiteList, msg.Sender) {
		return ""
	}
	if len(args.Bot.BlackList) > 0 && contains(args.Bot.BlackList, msg.Sender) {
		return ""
	}

	// 定制唤醒
	if msg.Content[0:1] != "/" {
		if strings.Contains(msg.Content, "@"+selfInfo.Name) {
			msg.Content = "/ai " + msg.Content
		} else {
			wakeWord := model.GetUserConfig(msg.Sender).WakeWord
			if wakeWord == "" {
				if !msg.IsGroup {
					msg.Content = "/ai " + msg.Content
				}
			} else if strings.HasPrefix(msg.Content, wakeWord) {
				msg.Content = "/ai " + strings.Replace(msg.Content, wakeWord, "", 1)
			}
		}
	}

	// 解析指令
	re := regexp.MustCompile(`^(/[\w:-]{2,20})(\s+|$)`)
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
	if hd.Level > 0 && !contains(args.Bot.Managers, msg.Sender) {
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

func contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
