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
		// 自动接受好友请求
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
		// 添加好友后主动回复
		if strings.Contains(msg.Content, "现在可以开始聊天了") {
			wc.CmdClient.SendTxt(args.Bot.Welcome, msg.Sender, "")
			return
		}
		// 添加群友后主动回复
		re := regexp.MustCompile(`邀请"(.+)"加入了群聊`)
		if matches := re.FindStringSubmatch(msg.Content); len(matches) > 1 {
			wc.CmdClient.SendTxt("欢迎 @"+matches[1]+"，"+args.Bot.Welcome, msg.Roomid, "")
			return
		}
		return
	}

}
