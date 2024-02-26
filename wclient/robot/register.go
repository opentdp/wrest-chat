package robot

import (
	"encoding/xml"
	"regexp"
	"strconv"
	"strings"

	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/dbase/chatroom"
	"github.com/opentdp/wechat-rest/dbase/message"
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wcferry/types"
	"github.com/opentdp/wechat-rest/wclient"
)

var wc *wcferry.Client
var selfInfo *wcferry.UserInfo

func Register() {

	if !args.Bot.Enable {
		logman.Warn("robot disabled")
		return
	}

	wc = wclient.Register()

	if len(handlers) == 0 {
		setupHandlers()
	}

	wc.EnrollReceiver(true, reciver)

}

func self() *wcferry.UserInfo {

	if selfInfo == nil {
		selfInfo = wc.CmdClient.GetSelfInfo()
	}

	return selfInfo

}

func reciver(msg *wcferry.WxMsg) {

	switch msg.Type {
	case 1: // 新消息
		hook1(msg)
	case 37: // 好友请求
		hook37(msg)
	case 10000: // 系统消息
		hook10000(msg)
	case 10002: // 撤回消息
		hook10002(msg)
	}

}

// 处理新消息
func hook1(msg *wcferry.WxMsg) {

	// 处理聊天指令
	if msg.IsGroup || wcferry.ContactType(msg.Sender) == "好友" {
		output := applyHandlers(msg)
		if output != "" && output != "-" {
			textReply(msg, output)
		}
		return
	}

}

// 新朋友通知
func hook37(msg *wcferry.WxMsg) {

	// 自动接受新朋友
	if args.Bot.AutoFriend {
		ret := &types.FriendRequestMsg{}
		err := xml.Unmarshal([]byte(msg.Content), ret)
		if err == nil && ret.FromUserName != "" {
			wc.CmdClient.AcceptNewFriend(ret.EncryptUserName, ret.Ticket, ret.Scene)
		}
	}

}

// 处理系统消息
func hook10000(msg *wcferry.WxMsg) {

	// 自动回应拍一拍
	if strings.Contains(msg.Content, "拍了拍我") {
		if args.Bot.AutoPat {
			wc.CmdClient.SendPatMsg(msg.Roomid, msg.Sender)
		}
		return
	}

	// 接受好友后响应
	if strings.Contains(msg.Content, "现在可以开始聊天了") {
		if len(args.Bot.WelcomeMsg) > 1 {
			wc.CmdClient.SendTxt(args.Bot.WelcomeMsg, msg.Sender, "")
		}
		return
	}

	// 有人进群时响应
	re := regexp.MustCompile(`邀请"(.+)"加入了群聊`)
	if matches := re.FindStringSubmatch(msg.Content); len(matches) > 1 {
		room, err := chatroom.Fetch(&chatroom.FetchParam{Roomid: msg.Roomid})
		if err == nil && len(room.WelcomeMsg) > 1 {
			wc.CmdClient.SendTxt("@"+matches[1]+"\n"+room.WelcomeMsg, msg.Roomid, "")
		}
		return
	}

}

// 处理撤回消息
func hook10002(msg *wcferry.WxMsg) {

	var output string

	if msg.IsGroup {
		room, _ := chatroom.Fetch(&chatroom.FetchParam{Roomid: msg.Roomid})
		output = room.RevokeMsg
	} else {
		output = args.Bot.RevokeMsg
	}

	if len(output) < 2 {
		return // 防撤回提示过短则忽略
	}

	ret := &types.SysMsg{}
	err := xml.Unmarshal([]byte(msg.Content), ret)

	if err == nil && ret.RevokeMsg.NewMsgID != "" {
		if id, _ := strconv.Atoi(ret.RevokeMsg.NewMsgID); id > 0 {
			revoke, err := message.Fetch(&message.FetchParam{Id: uint64(id)})
			if err == nil && revoke.Content != "" {
				str := strings.TrimSpace(revoke.Content)
				xmlPrefixes := []string{"<?xml", "<sysmsg", "<msg"}
				for _, prefix := range xmlPrefixes {
					if strings.HasPrefix(str, prefix) {
						str = ""
					}
				}
				if str != "" {
					output += "\n-------\n" + str
				} else {
					output += "\n-------\n暂不支持回显的消息类型"
				}
				textReply(msg, output)
			}
		}
	}

}
