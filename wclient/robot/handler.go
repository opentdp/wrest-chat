package robot

import (
	"sort"
	"strings"

	"github.com/opentdp/wechat-rest/dbase/chatroom"
	"github.com/opentdp/wechat-rest/dbase/profile"
	"github.com/opentdp/wechat-rest/dbase/setting"
	"github.com/opentdp/wechat-rest/dbase/tables"
	"github.com/opentdp/wechat-rest/wcferry"
)

type Handler struct {
	Level    int32                       // 0:不限制 7:群管理 9:创始人
	Order    int32                       // 排序，越小越靠前
	ChatAble bool                        // 是否允许在私聊使用
	RoomAble bool                        // 是否允许在群聊使用
	Describe string                      // 指令的描述信息
	PreCheck func(*wcferry.WxMsg) string // 前置检查，可拦截所有聊天内容
	Callback func(*wcferry.WxMsg) string // 指令回调，返回回复内容
}

var handlers = map[string]*Handler{}
var handlerKeys = []string{}

func setupHandlers() {

	aiHandler()
	apiHandler()
	badHandler()
	banHandler()
	topHandler()
	roomHandler()
	wgetHandler()

	helpHandler()

	// 对 handlers 进行排序

	for k := range handlers {
		handlerKeys = append(handlerKeys, k)
	}
	sort.Slice(handlerKeys, func(i, j int) bool {
		return handlers[handlerKeys[i]].Order < handlers[handlerKeys[j]].Order
	})

}

func clearHandlers() {

	handlers = map[string]*Handler{}
	handlerKeys = []string{}

	badMember = map[string]int{}
	keywordList = []*tables.Keyword{}

}

func applyHandlers(msg *wcferry.WxMsg) string {

	// 注册
	if len(handlers) == 0 {
		setupHandlers()
	}

	// 白名单
	if txt := whiteLimit(msg); txt != "" {
		return txt
	}

	// 前置钩子
	for _, k := range handlerKeys {
		v := handlers[k]
		if v.PreCheck != nil {
			if txt := v.PreCheck(msg); txt != "" {
				return txt
			}
		}
	}

	// 忽略空白
	msg.Content = strings.TrimSpace(msg.Content)
	if msg.Content == "" {
		return ""
	}

	// 解析指令
	matches := strings.SplitN(msg.Content, " ", 2)
	handler := handlers[matches[0]]
	if handler == nil {
		if msg.Content[0] == '/' {
			return setting.InvalidHandler
		}
		return ""
	}

	// 验证权限
	if handler.Level > 0 {
		up, _ := profile.Fetch(&profile.FetchParam{Wxid: msg.Sender, Roomid: prid(msg)})
		if up.Level < handler.Level {
			return setting.InvalidHandler
		}
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

	// 重写消息
	if len(matches) == 2 {
		msg.Content = strings.TrimSpace(matches[1])
	} else {
		msg.Content = ""
	}

	// 执行指令
	return handler.Callback(msg)

}

// 白名单模式
func whiteLimit(msg *wcferry.WxMsg) string {

	if !setting.WhiteLimit {
		return ""
	}

	// 管理豁免
	up, _ := profile.Fetch(&profile.FetchParam{Wxid: msg.Sender, Roomid: prid(msg)})
	if up.Level >= 7 {
		return ""
	}

	// 验证权限
	if msg.IsGroup {
		room, _ := chatroom.Fetch(&chatroom.FetchParam{Roomid: msg.Roomid})
		if room.Level < 2 {
			return "-"
		}
	} else if up.Level < 2 {
		return "-"
	}

	return ""

}
