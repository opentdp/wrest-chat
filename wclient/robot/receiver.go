package robot

import (
	"encoding/xml"
	"regexp"
	"strconv"
	"strings"

	"github.com/opentdp/wechat-rest/dbase/chatroom"
	"github.com/opentdp/wechat-rest/dbase/message"
	"github.com/opentdp/wechat-rest/dbase/setting"
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wcferry/types"
)

func receiver(msg *wcferry.WxMsg) {

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

	if len(handlers) == 0 {
		setupHandlers()
	}

	// 处理聊天指令
	if msg.IsGroup || wcferry.ContactType(msg.Sender) == "好友" {
		output := applyHandlers(msg)
		if strings.Trim(output, "-") != "" {
			reply(msg, output)
		}
		return
	}

}

// 新朋友通知
func hook37(msg *wcferry.WxMsg) {

	// 自动接受新朋友
	if setting.FriendAccept {
		ret := &types.FriendRequestMsg{}
		err := xml.Unmarshal([]byte(msg.Content), ret)
		if err == nil && ret.FromUserName != "" {
			wc.CmdClient.AcceptNewFriend(ret.EncryptUserName, ret.Ticket, ret.Scene)
		}
	}

}

// 处理系统消息
func hook10000(msg *wcferry.WxMsg) {

	// 接受好友后响应
	if strings.Contains(msg.Content, "现在可以开始聊天了") {
		if len(setting.FriendHello) > 1 {
			wc.CmdClient.SendTxt(setting.FriendHello, msg.Sender, "")
		}
		return
	}

	// 有人进群时响应
	re := regexp.MustCompile(`邀请"(.+)"加入了群聊`)
	if matches := re.FindStringSubmatch(msg.Content); len(matches) > 1 {
		room, _ := chatroom.Fetch(&chatroom.FetchParam{Roomid: msg.Roomid})
		if strings.Trim(room.WelcomeMsg, "-") != "" {
			wc.CmdClient.SendTxt("@"+matches[1]+"\n"+room.WelcomeMsg, msg.Roomid, "")
		}
		return
	}

	// 自动回应拍一拍
	if strings.Contains(msg.Content, "拍了拍我") {
		if msg.IsGroup {
			room, _ := chatroom.Fetch(&chatroom.FetchParam{Roomid: msg.Roomid})
			if strings.Trim(room.PatReturn, "-") != "" {
				wc.CmdClient.SendPatMsg(msg.Roomid, msg.Sender)
			}
		} else if setting.PatReturn {
			wc.CmdClient.SendPatMsg(msg.Roomid, msg.Sender)
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
		output = setting.RevokeMsg
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
				reply(msg, output)
			}
		}
	}

}
