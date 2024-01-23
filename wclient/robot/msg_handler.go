package robot

import (
	"net/url"
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
	apiHandler()
	badHandler()
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

	if txt := banMessagePrefix(msg); txt != "" {
		return txt
	}

	if txt := badMessagePrefix(msg); txt != "" {
		return txt
	}

	if txt := aiMessagePrefix(msg); txt != "" {
		return txt
	}

	// 空白消息
	msg.Content = strings.TrimSpace(msg.Content)
	if len(msg.Content) == 0 {
		return ""
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

// helper functions

func textReply(msg *wcferry.WxMsg, text string) int32 {

	if text = strings.TrimSpace(text); text == "" {
		return -1
	}

	if msg.IsGroup {
		user := wc.CmdClient.GetInfoByWxid(msg.Sender)
		return wc.CmdClient.SendTxt("@"+user.Name+"\n"+text, msg.Roomid, msg.Sender)
	} else {
		return wc.CmdClient.SendTxt(text, msg.Sender, "")
	}

}

func fileReply(msg *wcferry.WxMsg, text string) int32 {

	if text = strings.TrimSpace(text); text == "" {
		return -1
	}

	if u, err := url.Parse(text); err == nil {
		if u.Scheme == "http" || u.Scheme == "https" {
			if isImageFile(u.Path) {
				if msg.IsGroup {
					return wc.CmdClient.SendImg(text, msg.Roomid)
				} else {
					return wc.CmdClient.SendImg(text, msg.Sender)
				}
			} else {
				if msg.IsGroup {
					return wc.CmdClient.SendFile(text, msg.Roomid)
				} else {
					return wc.CmdClient.SendFile(text, msg.Sender)
				}
			}
		}
		return -1
	}

	return -2

}

func isImageFile(text string) bool {

	imageExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp", ".tiff", ".svg"}

	ext := strings.ToLower(text)
	for _, imageExt := range imageExtensions {
		if ext == imageExt {
			return true
		}
	}

	return false

}
