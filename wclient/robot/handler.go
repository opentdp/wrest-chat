package robot

import (
	"sort"
	"strings"

	"github.com/opentdp/wechat-rest/dbase/keyword"
	"github.com/opentdp/wechat-rest/dbase/profile"
	"github.com/opentdp/wechat-rest/dbase/setting"
	"github.com/opentdp/wechat-rest/wcferry"
)

var Handlers = []*Handler{}
var HandlerMap = map[string]*Handler{}

type Handler struct {
	Level    int32                       `json:"level"`     // 0:不限制 7:群管理 9:创始人
	Order    int32                       `json:"order"`     // 排序，越小越靠前
	ChatAble bool                        `json:"chat_able"` // 是否允许在私聊使用
	RoomAble bool                        `json:"room_able"` // 是否允许在群聊使用
	Command  string                      `json:"command"`   // 指令
	Describe string                      `json:"describe"`  // 指令的描述信息
	PreCheck func(*wcferry.WxMsg) string // 前置检查，可拦截文本聊天内容
	Callback func(*wcferry.WxMsg) string // 指令回调，返回回复内容
}

func initHandlers() {

	hlst := []*Handler{}
	hmap := map[string]*Handler{}

	// 获取指令列表
	hlst = append(hlst, aiHandler()...)
	hlst = append(hlst, apiHandler()...)
	hlst = append(hlst, badHandler()...)
	hlst = append(hlst, banHandler()...)
	hlst = append(hlst, topHandler()...)
	hlst = append(hlst, roomHandler()...)
	hlst = append(hlst, wgetHandler()...)
	hlst = append(hlst, helpHandler()...)

	// 指令列表排序
	sort.Slice(hlst, func(i, j int) bool {
		return hlst[i].Order < hlst[j].Order
	})

	// 获取指令映射
	for _, v := range hlst {
		hmap[v.Command] = v
	}

	// 获取别名数据
	kws, err := keyword.FetchAll(&keyword.FetchAllParam{Group: "handler"})
	if err == nil && len(kws) > 0 {
		for _, v := range kws {
			if hmap[v.Target] != nil {
				hmap[v.Phrase+"@"+v.Roomid] = hmap[v.Target]
			}
		}
	}

	// 更新全局数据
	Handlers, HandlerMap = hlst, hmap

}

func applyHandlers(msg *wcferry.WxMsg) string {

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

	// 修正空白
	if msg.Content[0] == '@' {
		msg.Content = strings.Replace(msg.Content, " ", " ", 1)
	}

	// 解析指令
	params := strings.SplitN(msg.Content, " ", 2)
	handler := HandlerMap[params[0]] // 默认
	if handler == nil {
		handler = HandlerMap[params[0]+"@"+prid(msg)] // 群聊
		if handler == nil {
			handler = HandlerMap[params[0]+"@-"] // 全局
			if handler == nil {
				if msg.Content[0] == '/' {
					return setting.InvalidHandler
				}
				return ""
			}
		}
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
	if len(params) == 2 {
		msg.Content = strings.TrimSpace(params[1])
	} else {
		msg.Content = ""
	}

	// 执行指令
	return handler.Callback(msg)

}
