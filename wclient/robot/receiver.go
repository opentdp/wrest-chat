package robot

import (
	"encoding/xml"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/opentdp/go-helper/logman"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/dbase/chatroom"
	"github.com/opentdp/wechat-rest/dbase/message"
	"github.com/opentdp/wechat-rest/dbase/setting"
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wcferry/types"
)

func receiver(msg *wcferry.WxMsg) {

	switch msg.Type {
	case 1: //文字
		hook1(msg)
	case 3: //图片
		hook3(msg)
	case 37: //好友确认
		hook37(msg)
	case 10000: //红包、系统消息
		hook10000(msg)
	case 10002: //撤回消息
		hook10002(msg)
	}

}

// 处理新消息
func hook1(msg *wcferry.WxMsg) {

	// 注册指令函数
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

// 自动保存图片
func hook3(msg *wcferry.WxMsg) {

	if !setting.AutoSaveImage || msg.Extra == "" {
		return
	}

	dir := path.Dir(msg.Extra)
	p, err := wc.CmdClient.DownloadImage(msg.Id, msg.Extra, dir, 5)
	if err != nil || p == "" {
		logman.Error("download image error", "err", err)
		return
	}

	logman.Info("downloaded image", "path", p)
	if args.Wcf.MsgStore {
		message.Update(&message.UpdateParam{Id: msg.Id, Remark: p})
	}

}

// 处理新朋友通知
func hook37(msg *wcferry.WxMsg) {

	// 自动接受新朋友
	if setting.FriendAccept {
		ret := &types.FriendRequestContent{}
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

// 处理撤回消息
func hook10002(msg *wcferry.WxMsg) {

	var output string

	// 获取撤回提示
	if msg.IsGroup {
		room, _ := chatroom.Fetch(&chatroom.FetchParam{Roomid: msg.Roomid})
		output = room.RevokeMsg
	} else {
		output = setting.RevokeMsg
	}

	// 防撤回提示已关闭
	if len(output) < 2 {
		return
	}

	// 解析已撤回的消息
	ret := &types.RevokeContent{}
	err := xml.Unmarshal([]byte(msg.Content), ret)
	if err != nil || ret.RevokeMsg.NewMsgID == "" {
		return
	}

	// 获取已撤回消息的 Id
	id, err := strconv.Atoi(ret.RevokeMsg.NewMsgID)
	if err != nil || id == 0 {
		return
	}

	// 取回已撤回的消息内容
	revoke, err := message.Fetch(&message.FetchParam{Id: uint64(id)})
	if err != nil || revoke.Content == "" {
		return
	}

	// 提示已撤回的消息内容
	str := strings.TrimSpace(revoke.Content)
	xmlPrefixes := []string{"<?xml", "<sysmsg", "<msg"}
	for _, prefix := range xmlPrefixes {
		if strings.HasPrefix(str, prefix) {
			str = ""
		}
	}

	if str != "" {
		output += "\n-------\n" + str
		reply(msg, output)
		return
	}

	if revoke.Type == 3 {
		if revoke.Remark != "" {
			if revoke.IsGroup {
				wc.CmdClient.SendImg(revoke.Remark, revoke.Roomid)
			} else {
				wc.CmdClient.SendImg(revoke.Remark, revoke.Sender)
			}
		}
		output += "\n-------\n一张不可描述的图片"
		reply(msg, output)
		return
	}

	if revoke.Type == 47 {
		output += "\n-------\n一个浪漫无敌的表情"
		reply(msg, output)
		return
	}

	output += "\n-------\n暂不支持回显的消息类型"
	reply(msg, output)

}
