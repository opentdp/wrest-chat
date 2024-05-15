package robot

import (
	"github.com/opentdp/wrest-chat/dbase/chatroom"
	"github.com/opentdp/wrest-chat/dbase/profile"
	"github.com/opentdp/wrest-chat/dbase/setting"
	"github.com/opentdp/wrest-chat/wcferry"
)

func receiver(msg *wcferry.WxMsg) {

	if whiteLimit(msg) {
		return // 白名单限制
	}

	switch msg.Type {
	case 1: //文字
		receiver1(copyMsg(msg))
	case 3: //图片
		receiver3(msg)
	case 37: //好友确认
		receiver37(msg)
	case 49: //混合消息
		receiver49(copyMsg(msg))
	case 10000: //红包、系统消息
		receiver10000(msg)
	case 10002: //撤回消息
		receiver10002(msg)
	}

}

// 复制消息
// return 深拷贝后的消息
func copyMsg(msg *wcferry.WxMsg) *wcferry.WxMsg {

	return &wcferry.WxMsg{
		IsSelf:  msg.IsSelf,
		IsGroup: msg.IsGroup,
		Type:    msg.Type,
		Ts:      msg.Ts,
		Roomid:  msg.Roomid,
		Content: msg.Content,
		Sender:  msg.Sender,
		Sign:    msg.Sign,
		Thumb:   msg.Thumb,
		Extra:   msg.Extra,
		Xml:     msg.Xml,
	}

}

// 白名单限制
// return 验证结果 [true 受限, false 忽略]
func whiteLimit(msg *wcferry.WxMsg) bool {

	// 无需验证
	if !setting.WhiteLimit1 && !setting.WhiteLimit2 {
		return false
	}

	// 管理员豁免
	up, _ := profile.Fetch(&profile.FetchParam{Wxid: msg.Sender, Roomid: prid(msg)})
	if up.Level > 6 {
		return false
	}

	// 白名单验证
	if msg.IsGroup {
		if setting.WhiteLimit1 {
			room, _ := chatroom.Fetch(&chatroom.FetchParam{Roomid: msg.Roomid})
			return room.Level < 2
		}
	} else {
		if setting.WhiteLimit2 {
			return up.Level < 2
		}
	}

	// 默认不受限
	return false

}

// 组策略限制
// return 验证结果 [true 受限, false 忽略]
func groupLimit(msg *wcferry.WxMsg, level int32, roomid string) bool {

	// 验证权限
	if level > 0 {
		up, _ := profile.Fetch(&profile.FetchParam{Wxid: msg.Sender, Roomid: prid(msg)})
		if up.Level < level {
			return true
		}
	}

	// 验证场景
	if msg.IsGroup {
		if roomid != "*" && roomid != "+" && roomid != msg.Roomid {
			return true
		}
	} else {
		if roomid != "*" && roomid != "-" {
			return true
		}
	}

	return false

}
