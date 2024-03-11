package robot

import (
	"encoding/xml"
	"strings"

	"github.com/opentdp/wechat-rest/dbase/message"
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wcferry/types"
)

// 处理新朋友通知
// 处理混合类消息
func hook49(msg *wcferry.WxMsg) {

	ret := types.MsgContent49{}
	err := xml.Unmarshal([]byte(msg.Content), &ret)
	if err != nil {
		return
	}

	title := ret.AppMsg.Title
	refId := ret.AppMsg.ReferMsg.Svrid

	if ret.AppMsg.Type == 57 {
		// 撤回引用的消息
		// TDOO: 未实现鉴权
		if strings.HasPrefix(title, "撤回") {
			wc.CmdClient.RevokeMsg(refId)
			return
		}
		// 引用图片
		if ret.AppMsg.ReferMsg.Type == 3 {
			origin, err := message.Fetch(&message.FetchParam{Id: refId})
			if err == nil && origin.Remark != "" {
				msg.Thumb = origin.Remark
			}
			msg.Content = title
			msg.Extra = "image-txt"
			receiver1(msg)
			return
		}
		// 引用混合类消息
		if ret.AppMsg.ReferMsg.Type == 49 {
			origin, err := message.Fetch(&message.FetchParam{Id: refId})
			if err == nil && origin.Content != "" {
				msg.Content = title + "\nXML数据如下:\n" + origin.Content
				msg.Extra = "record-txt"
				receiver1(msg)
				return
			}
			return
		}
	}

}
