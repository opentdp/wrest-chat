package robot

import (
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

func setupHandlers() {

	aiHandler()
	banHandler()
	newHandler()
	modelHandler()
	roomHandler()
	wakeHandler()

	helpHandler()

}

func applyHandlers(msg *wcferry.WxMsg) string {

	if len(handlers) == 0 {
		setupHandlers()
	}

	// 前置过滤
	if len(args.Bot.WhiteList) > 0 && !sliceContains(args.Bot.WhiteList, msg.Sender) {
		return ""
	} else if sliceContains(args.Bot.BlackList, msg.Sender) {
		return ""
	} else if len(msg.Content) == 0 {
		return ""
	}

	// 定制唤醒
	if msg.Content[0:1] != "/" {
		if strings.Contains(msg.Content, "@"+selfInfo.Name) {
			msg.Content = "/ai " + msg.Content
		} else {
			wakeWord := model.GetUserConfig(msg.Sender).WakeWord
			if wakeWord == "" {
				msg.Content = "/ai " + msg.Content
			} else if strings.HasPrefix(msg.Content, wakeWord) {
				msg.Content = strings.Replace(msg.Content, wakeWord, "/ai ", 1)
			}
		}
	}

	// 解析指令
	re := regexp.MustCompile(`^(/[\w:-]{2,20})\s*(.*)$`)
	matches := re.FindStringSubmatch(msg.Content)
	if len(matches) == 3 {
		msg.Content = matches[2]
	} else {
		return ""
	}

	// 查找指令
	handler, exists := handlers[matches[1]]
	if !exists {
		return "指令或参数错误, 回复 /help 获取帮助"
	}

	// 检查权限
	if handler.Level > 0 && !sliceContains(args.Bot.Managers, msg.Sender) {
		return "无权限使用此指令"
	}

	// 检查场景
	if msg.IsGroup {
		if !handler.RoomAble {
			return "此指令仅在私聊中可用"
		}
	} else {
		if !handler.ChatAble {
			return "此指令仅在群聊中可用"
		}
	}

	// 执行指令
	return handler.Callback(msg)

}
