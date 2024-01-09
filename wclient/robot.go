package wclient

import (
	"encoding/xml"
	"strings"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wclient/proto"
)

func Robot() {

	Connect()

	wc.EnrollReceiver(true, Reciver)

}

func Reciver(msg *wcferry.WxMsg) {

	switch msg.Type {
	case 1:
		// 忽略公号消息
		// if strings.HasPrefix(msg.Sender, "gh_") {
		// 	return
		// }
		return
	case 37:
		// 自动接受好友请求
		dat := &proto.FriendRequestMsg{}
		err := xml.Unmarshal([]byte(msg.Content), dat)
		if err == nil && dat.FromUserName != "" {
			wc.CmdClient.AcceptNewFriend(dat.EncryptUserName, dat.Ticket, dat.Scene)
		}
		return
	case 10000:
		// 拍一拍
		if strings.Contains(msg.Content, "拍了拍我") {
			wc.CmdClient.SendPatMsg(msg.Roomid, msg.Sender)
			return
		}
		// 添加好友后自动回复
		if strings.Contains(msg.Content, "现在可以开始聊天了") {
			wc.CmdClient.SendTxt(args.Bot.FriendWelcome, msg.Sender, "")
			return
		}
		return
	}

}
