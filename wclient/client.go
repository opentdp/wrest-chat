package wclient

import (
	"net"
	"net/url"
	"strings"

	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/strutil"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/dbase/message"
	"github.com/opentdp/wechat-rest/wcferry"
)

var wc *wcferry.Client

// 注册客户端（单例模式）
func Register() *wcferry.Client {

	if wc != nil {
		return wc
	}

	// 检查参数
	host, port, err := net.SplitHostPort(args.Wcf.Address)
	if err != nil {
		logman.Fatal("invalid address", "error", err)
	}

	// 创建客户端
	wc = &wcferry.Client{
		ListenAddr: host,
		ListenPort: strutil.ToInt(port),
		WcfBinary:  args.Wcf.WcfBinary,
	}

	// 初始化连接
	logman.Warn("wcf starting ...")
	if err := wc.Connect(); err != nil {
		logman.Fatal("wcf start failed", "error", err)
	}

	// 存储收到的消息
	if args.Wcf.MsgStore {
		wc.EnrollReceiver(true, msgToDatabase)
	}

	// 打印收到的消息
	if args.Wcf.MsgPrint {
		wc.EnrollReceiver(true, wcferry.WxMsgPrinter)
	}

	return wc

}

// 发送弹性消息（文本、卡片、网络图片或文件）
// 网络文件格式：以 http:// 或 https:// 开头
// 卡片消息格式：card\n{name}\n{account}\n{title}\n{digest}\n{url}\n{thumburl}
// param msg string 要发送的消息
// param wxid string 消息接收人，如果 roomid 存在则为 at 此人
// param roomid string 消息接收群，空则为私聊
// return int32 0 为成功，其他失败
func SendFlexMsg(msg, wxid, roomid string) int32 {
	wxid = strings.Trim(wxid, "-")
	roomid = strings.Trim(roomid, "-")
	receiver, ater := wxid, ""
	if roomid != "" {
		receiver, ater = roomid, wxid
	}
	// 发送网络文件
	if strings.HasPrefix(msg, "http://") || strings.HasPrefix(msg, "https://") {
		if u, err := url.Parse(msg); err == nil {
			if wcferry.IsImageFile(u.Path) {
				return wc.CmdClient.SendImg(msg, receiver)
			}
			return wc.CmdClient.SendFile(msg, receiver)
		}
	}
	// 发送卡片信息
	if strings.HasPrefix(msg, "card\n") {
		if args := strings.Split(msg, "\n")[1:]; len(args) >= 6 {
			return wc.CmdClient.SendRichText(args[0], args[1], args[2], args[3], args[4], args[5], receiver)
		}
	}
	// 发送文本信息
	if ater != "" {
		user := wc.CmdClient.GetInfoByWxid(ater)
		if user != nil && user.Name != "" {
			msg = "@" + user.Name + "\n" + msg
		}
	}
	return wc.CmdClient.SendTxt(msg, receiver, ater)
}

// 存储收到的消息
// param msg *wcferry.WxMsg 收到的消息
func msgToDatabase(msg *wcferry.WxMsg) {

	rq := message.CreateParam{
		Id:      msg.Id,
		IsSelf:  msg.IsSelf,
		IsGroup: msg.IsGroup,
		Type:    msg.Type,
		Ts:      msg.Ts,
		Roomid:  msg.Roomid,
		Content: msg.Content,
		Sender:  msg.Sender,
		Sign:    msg.Sign,
		Thumb:   msg.Thumb,
		Extra:   msg.Extra,
		Xml:     msg.Xml,
	}

	message.Create(&rq)

}
