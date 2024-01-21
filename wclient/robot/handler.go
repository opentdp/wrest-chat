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
	saveHandler()

	helpHandler()

}

func applyHandlers(msg *wcferry.WxMsg) string {

	if ingoreMessage(msg) {
		return ""
	}

	if len(handlers) == 0 {
		setupHandlers()
	}

	// 定制唤醒
	if msg.Content[0:1] != "/" {
		if strings.Contains(msg.Content, "@"+selfInfo.Name) {
			msg.Content = "/ai " + msg.Content
		} else {
			wakeWord := model.GetUser(msg.Sender).AiArgot
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
	handler, ok := handlers[matches[1]]
	if !ok {
		return "指令或参数错误, 回复 /help 获取帮助"
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

	// 检查权限
	if handler.Level > 0 {
		user, ok := args.Usr.Member[msg.Sender]
		if ok && user.Level >= handler.Level {
			return "无权限使用此指令"
		}
		room, ok := args.Usr.Room[msg.Roomid]
		if ok && room.Level >= handler.Level {
			return "无权限使用此指令"
		}
	}

	// 执行指令
	return handler.Callback(msg)

}

func ingoreMessage(msg *wcferry.WxMsg) bool {

	// 空指令
	if len(msg.Content) == 0 {
		return true
	}

	// 用户权限
	if user, ok := args.Usr.Member[msg.Sender]; ok {
		if user.Level == 9 { // 管理员
			return false
		}
		if user.Level == 4 { // 已禁止
			return true
		}
	}

	// 群聊权限
	if room, ok := args.Usr.Room[msg.Roomid]; ok {
		if room.Level == 4 { // 已禁止
			return true
		}
	}

	return false

}
