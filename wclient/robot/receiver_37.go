package robot

import (
	"encoding/xml"

	"github.com/opentdp/wechat-rest/dbase/setting"
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wcferry/types"
)

func receiver37(msg *wcferry.WxMsg) {

	// 自动接受新朋友
	if setting.FriendAccept {
		ret := types.MsgContent37{}
		err := xml.Unmarshal([]byte(msg.Content), &ret)
		if err == nil && ret.FromUserName != "" {
			wc.CmdClient.AcceptNewFriend(ret.EncryptUserName, ret.Ticket, ret.Scene)
		}
	}

}
