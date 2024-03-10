package robot

import (
	"encoding/xml"

	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wcferry/types"
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
		msg.Id = ret.AppMsg.ReferMsg.Svrid
		msg.Content = ret.AppMsg.Title
		msg.Sign = "refer-msg"
		receiver1(msg)
		return
	}

}
