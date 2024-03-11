package robot

import (
	"regexp"
	"strings"
	"time"

	"github.com/opentdp/wechat-rest/dbase/chatroom"
	"github.com/opentdp/wechat-rest/dbase/setting"
	"github.com/opentdp/wechat-rest/wcferry"
)

// 处理系统消息
func receiver10000(msg *wcferry.WxMsg) {

	// 接受好友后响应
	if strings.Contains(msg.Content, "现在可以开始聊天了") {
		if len(setting.FriendHello) > 1 {
			time.Sleep(2 * time.Second) // 延迟 2 秒
			wc.CmdClient.SendTxt(setting.FriendHello, msg.Sender, "")
		}
		return
	}

	// 邀请"xxx"加入了群聊
	r1 := regexp.MustCompile(`邀请"(.+)"加入了群聊`)
	if matches := r1.FindStringSubmatch(msg.Content); len(matches) > 1 {
		room, _ := chatroom.Fetch(&chatroom.FetchParam{Roomid: msg.Roomid})
		if strings.Trim(room.WelcomeMsg, "-") != "" {
			time.Sleep(3 * time.Second) // 延迟 3 秒
			wc.CmdClient.SendTxt("@"+matches[1]+"\n"+room.WelcomeMsg, msg.Roomid, "")
		}
		return
	}

	// "xxx"通过扫描"xxx"分享的二维码加入群聊
	r2 := regexp.MustCompile(`"(.+)"通过扫描"(.+)"分享的二维码加入群聊`)
	if matches := r2.FindStringSubmatch(msg.Content); len(matches) > 1 {
		room, _ := chatroom.Fetch(&chatroom.FetchParam{Roomid: msg.Roomid})
		if strings.Trim(room.WelcomeMsg, "-") != "" {
			time.Sleep(3 * time.Second) // 延迟 3 秒
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
