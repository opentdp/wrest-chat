package robot

import (
	"sort"
	"strings"

	"github.com/opentdp/wechat-rest/dbase/chatroom"
	"github.com/opentdp/wechat-rest/dbase/profile"
	"github.com/opentdp/wechat-rest/dbase/setting"
	"github.com/opentdp/wechat-rest/wcferry"
)

type Handler struct {
	Level    int32                       // 0:不限制 7:群管理 9:创始人
	Order    int32                       // 排序，越小越靠前
	ChatAble bool                        // 是否允许在私聊使用
	RoomAble bool                        // 是否允许在群聊使用
	Command  string                      // 指令
	Describe string                      // 指令的描述信息
	PreCheck func(*wcferry.WxMsg) string // 前置检查，可拦截所有聊天内容
	Callback func(*wcferry.WxMsg) string // 指令回调，返回回复内容
}

var Handlers = []*Handler{}
var HandlerMap = map[string]*Handler{}

func initHandlers() {

	list := []*Handler{}
	lmap := map[string]*Handler{}

	list = append(list, aiHandler()...)
	list = append(list, apiHandler()...)
	list = append(list, badHandler()...)
	list = append(list, banHandler()...)
	list = append(list, topHandler()...)
	list = append(list, roomHandler()...)
	list = append(list, wgetHandler()...)
	list = append(list, helpHandler()...)

	// 格式化
	for _, v := range list {
		lmap[v.Command] = v
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Order < list[j].Order
	})

	// 更新列表
	Handlers = list
	HandlerMap = lmap

}

func applyHandlers(msg *wcferry.WxMsg) string {

	// 白名单
	if txt := whiteLimit(msg); txt != "" {
		return txt
	}

	// 前置钩子
	for _, v := range Handlers {
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
	handler := HandlerMap[matches[0]]
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
