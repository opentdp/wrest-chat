package robot

import (
	"encoding/xml"
	"regexp"
	"strings"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wcferry/types"
	"github.com/opentdp/wechat-rest/wclient"
)

var wc *wcferry.Client
var selfInfo *wcferry.UserInfo

func Register() {

	wc = wclient.Register()

	selfInfo = wc.CmdClient.GetSelfInfo()

	wc.EnrollReceiver(true, reciver)

}

func reciver(msg *wcferry.WxMsg) {

	switch msg.Type {
	case 1:
		// 处理聊天指令
		if msg.IsGroup || wcferry.ContactType(msg.Sender) == "好友" {
			if output := applyHandlers(msg); output != "" {
				if msg.IsGroup {
					user := wc.CmdClient.GetInfoByWxid(msg.Sender)
					wc.CmdClient.SendTxt("@"+user.Name+"\n"+output, msg.Roomid, msg.Sender)
				} else {
					wc.CmdClient.SendTxt(output, msg.Sender, "")
				}
			}
		}
	case 37:
		// 自动接受新朋友
		ret := &types.FriendRequestMsg{}
		err := xml.Unmarshal([]byte(msg.Content), ret)
		if err == nil && ret.FromUserName != "" {
			wc.CmdClient.AcceptNewFriend(ret.EncryptUserName, ret.Ticket, ret.Scene)
		}
	case 10000:
		// 自动回应拍一拍
		if strings.Contains(msg.Content, "拍了拍我") {
			wc.CmdClient.SendPatMsg(msg.Roomid, msg.Sender)
			return
		}
		// 接受好友后响应
		if strings.Contains(msg.Content, "现在可以开始聊天了") {
			if args.Bot.Welcome != "" {
				wc.CmdClient.SendTxt(args.Bot.Welcome, msg.Sender, "")
			}
			return
		}
		// 有人进群时响应
		re := regexp.MustCompile(`邀请"(.+)"加入了群聊`)
		if matches := re.FindStringSubmatch(msg.Content); len(matches) > 1 {
			if room := args.GetChatRoom(msg.Roomid); room.Welcome != "" {
				wc.CmdClient.SendTxt("@"+matches[1]+"\n"+room.Welcome, msg.Roomid, "")
			}
			return
		}
	case 10002:
		// 撤回消息时响应
		ret := &types.SysMsg{}
		err := xml.Unmarshal([]byte(msg.Content), ret)
		if err == nil && ret.RevokeMsg.MsgID != "" && args.Bot.Revoke != "" {
			if msg.IsGroup {
				user := wc.CmdClient.GetInfoByWxid(msg.Sender)
				wc.CmdClient.SendTxt("@"+user.Name+" "+args.Bot.Revoke, msg.Roomid, msg.Sender)
			} else {
				wc.CmdClient.SendTxt(args.Bot.Revoke, msg.Sender, "")
			}
		}
	}

}
