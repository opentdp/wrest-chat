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
	wc.EnrollReceiver(true, reciver)

	selfInfo = wc.CmdClient.GetSelfInfo()

}

func reciver(msg *wcferry.WxMsg) {

	switch msg.Type {
	case 1:
		// 忽略公号消息
		if strings.HasPrefix(msg.Sender, "gh_") {
			return
		}
		// 处理聊天指令
		if output := applyHandlers(msg); output != "" {
			if msg.IsGroup {
				user := wc.CmdClient.GetInfoByWxid(msg.Sender)
				wc.CmdClient.SendTxt("@"+user.Name+"\n"+output, msg.Roomid, msg.Sender)
			} else {
				wc.CmdClient.SendTxt(output, msg.Sender, "")
			}
			return
		}
		return
	case 37:
		// 自动确认好友请求
		ret := &types.FriendRequestMsg{}
		err := xml.Unmarshal([]byte(msg.Content), ret)
		if err == nil && ret.FromUserName != "" {
			wc.CmdClient.AcceptNewFriend(ret.EncryptUserName, ret.Ticket, ret.Scene)
		}
		return
	case 10000:
		// 自动回应拍一拍
		if strings.Contains(msg.Content, "拍了拍我") {
			wc.CmdClient.SendPatMsg(msg.Roomid, msg.Sender)
			return
		}
		// 添加好友后响应
		if strings.Contains(msg.Content, "现在可以开始聊天了") {
			if args.Bot.Welcome != "" {
				wc.CmdClient.SendTxt(args.Bot.Welcome, msg.Sender, "")
			}
			return
		}
		// 有人进群时响应
		re := regexp.MustCompile(`邀请"(.+)"加入了群聊`)
		if matches := re.FindStringSubmatch(msg.Content); len(matches) > 1 {
			welcome := args.Bot.Welcome
			for _, room := range args.Bot.HostedRooms {
				if room.RoomId == msg.Roomid {
					welcome = room.Welcome
					break
				}
			}
			if welcome != "" {
				wc.CmdClient.SendTxt("@"+matches[1]+"\n"+welcome, msg.Roomid, "")
			}
			return
		}
		return
	case 10002:
		// 撤回消息时响应
		ret := &types.SysMsg{}
		err := xml.Unmarshal([]byte(msg.Content), ret)
		if err == nil && ret.RevokeMsg.MsgID != "" && args.Bot.Revoke != "" {
			if msg.IsGroup {
				user := wc.CmdClient.GetInfoByWxid(msg.Sender)
				wc.CmdClient.SendTxt("@"+user.Name+" 撤回了寂寞？", msg.Roomid, msg.Sender)
			} else {
				wc.CmdClient.SendTxt("撤回了寂寞？", msg.Sender, "")
			}
		}

	}

}
