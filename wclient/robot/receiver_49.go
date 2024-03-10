package robot

import (
	"encoding/xml"

	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wcferry/types"
)

// 处理新朋友通知
// 处理混合类消息
func receiver49(msg *wcferry.WxMsg) {

	ret := types.MsgContent49{}
	err := xml.Unmarshal([]byte(msg.Content), &ret)
	if err != nil {
		return
	}

	if ret.AppMsg.Type == 57 {
		msg.Id = ret.AppMsg.ReferMsg.Svrid
		msg.Content = ret.AppMsg.Title
		msg.Extra = "refer-msg"
		receiver1(msg)
		return
	}

}
