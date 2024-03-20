package robot

import (
	"regexp"
	"strings"
	"time"

	"github.com/opentdp/wrest-chat/dbase/chatroom"
	"github.com/opentdp/wrest-chat/dbase/setting"
	"github.com/opentdp/wrest-chat/wcferry"
)

// 处理系统消息
func receiver10000(msg *wcferry.WxMsg) {

	if msg.IsGroup {
		receiver10000Public(msg)
	} else {
		receiver10000Private(msg)
	}

}

func receiver10000Public(msg *wcferry.WxMsg) {

	// 自动回应群聊拍一拍（私聊不支持）
	if strings.Contains(msg.Content, "拍了拍我") {
		room, _ := chatroom.Fetch(&chatroom.FetchParam{Roomid: msg.Roomid})
		if room.PatReturn == "true" {
			wc.CmdClient.SendPatMsg(msg.Roomid, msg.Sender)
		}
		return
	}

	// 邀请"xxx"加入了群聊
	r1 := regexp.MustCompile(`邀请"(.+)"加入了群聊`)
	if matches := r1.FindStringSubmatch(msg.Content); len(matches) > 1 {
		room, _ := chatroom.Fetch(&chatroom.FetchParam{Roomid: msg.Roomid})
		if len(room.WelcomeMsg) > 1 {
			time.Sleep(1 * time.Second) // 延迟1秒
			reply(msg, "@"+matches[1]+"\n"+room.WelcomeMsg)
		}
		return
	}

	// "xxx"通过扫描"xxx"分享的二维码加入群聊
	r2 := regexp.MustCompile(`"(.+)"通过扫描"(.+)"分享的二维码加入群聊`)
	if matches := r2.FindStringSubmatch(msg.Content); len(matches) > 1 {
		room, _ := chatroom.Fetch(&chatroom.FetchParam{Roomid: msg.Roomid})
		if len(room.WelcomeMsg) > 1 {
			time.Sleep(1 * time.Second) // 延迟1秒
			reply(msg, "@"+matches[1]+"\n"+room.WelcomeMsg)
		}
		return
	}

}

func receiver10000Private(msg *wcferry.WxMsg) {

	// 接受好友后响应
	if strings.Contains(msg.Content, "现在可以开始聊天了") {
		if len(setting.FriendHello) > 1 {
			reply(msg, setting.FriendHello)
		}
		return
	}

}
