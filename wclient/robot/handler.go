package robot

import (
	"regexp"
	"strings"

	"github.com/opentdp/wechat-rest/dbase/profile"
	"github.com/opentdp/wechat-rest/dbase/tables"
	"github.com/opentdp/wechat-rest/wcferry"
)

type Handler struct {
	Level    int32  // 0:不限制 9:管理员
	Order    int32  // 排序，越小越靠前
	ChatAble bool   // 是否允许在私聊使用
	RoomAble bool   // 是否允许在群聊使用
	Describe string // 指令描述
	Callback func(msg *wcferry.WxMsg) string
	PreCheck func(msg *wcferry.WxMsg) string
}

var handlers = map[string]*Handler{}

func setupHandlers() {

	aiHandler()
	apiHandler()
	badHandler()
	banHandler()
	roomHandler()
	wgetHandler()

	helpHandler()

}

func clearHandlers() {

	badMember = map[string]int{}
	keywordList = []*tables.Keyword{}

	handlers = map[string]*Handler{}

}

func applyHandlers(msg *wcferry.WxMsg) string {

	// 前置检查
	for _, v := range handlers {
		if v.PreCheck != nil {
			if txt := v.PreCheck(msg); txt != "" {
				return txt
			}
		}
	}

	// 空白消息
	content := strings.TrimSpace(msg.Content)
	if len(content) == 0 {
		return ""
	}

	// 解析指令
	re := regexp.MustCompile(`^(/[\w:-]{2,20})\s*(.*)$`)
	matches := re.FindStringSubmatch(content)
	if len(matches) != 3 {
		return ""
	}

	// 查找指令
	handler, ok := handlers[matches[1]]
	if !ok {
		return "指令未注册或参数错误"
	}

	// 验证级别
	if handler.Level > 0 {
		up, _ := profile.Fetch(&profile.FetchParam{Wxid: msg.Sender, Roomid: prid(msg)})
		if up.Level < handler.Level {
			return "此指令已被限制使用"
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

	// 执行指令
	msg.Content = strings.TrimSpace(matches[2])
	return handler.Callback(msg)

}
