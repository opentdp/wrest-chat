package robot

import (
	"regexp"
	"strings"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/wcferry"
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
	wgetHandler()

	helpHandler()

}

func applyHandlers(msg *wcferry.WxMsg) string {

	if !validMessage(msg) {
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
			wakeWord := args.GetMember(msg.Sender).AiArgot
			if wakeWord == "" {
				if !msg.IsGroup {
					msg.Content = "/ai " + msg.Content
				}
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

	// 验证场景
	if msg.IsGroup {
		if !handler.RoomAble {
			return "此指令仅在私聊中可用"
		}
	} else {
		if !handler.ChatAble {
			return "此指令仅在群聊中可用"
		}
	}

	// 指令权限
	if handler.Level > 0 {
		if args.GetMember(msg.Sender).Level < handler.Level {
			return "无权限使用此指令"
		}
		if msg.IsGroup {
			room := args.GetChatRoom(msg.Roomid)
			if room.GetMember(msg.Sender).Level < handler.Level {
				return "无权限使用此指令"
			}
		}
	}

	// 执行指令
	return handler.Callback(msg)

}

func validMessage(msg *wcferry.WxMsg) bool {

	// 空白指令
	msg.Content = strings.TrimSpace(msg.Content)
	if len(msg.Content) == 0 {
		return false
	}

	// 全局权限
	user := args.GetMember(msg.Sender)
	if args.Bot.WhiteMember {
		if user.Level <= 1 {
			return false
		}
	} else {
		if user.Level == 1 {
			return false
		}
	}

	// 群聊权限
	if msg.IsGroup {
		room := args.GetChatRoom(msg.Roomid)
		user := room.GetMember(msg.Sender)
		if args.Bot.WhiteChatRoom {
			if room.Level <= 1 {
				return false
			}
		} else {
			if room.Level == 1 {
				return false
			}
			if user.Level == 1 {
				return false
			}
		}
	}

	return true

}
