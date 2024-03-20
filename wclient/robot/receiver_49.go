package robot

import (
	"encoding/xml"

	"github.com/opentdp/wrest-chat/wcferry"
	"github.com/opentdp/wrest-chat/wcferry/types"
)

// 处理混合类消息
func receiver49(msg *wcferry.WxMsg) {

	ret := types.MsgContent49{}
	err := xml.Unmarshal([]byte(msg.Content), &ret)
	if err != nil {
		return
	}

	// 引用消息
	if ret.AppMsg.Type == 57 {
		msg.Extra = msg.Content
		msg.Content = ret.AppMsg.Title
		msg.Id = ret.AppMsg.ReferMsg.Svrid
		msg.Type = ret.AppMsg.ReferMsg.Type
		msg.Sign = "refer-msg"
		receiver1(msg)
		return
	}

}
