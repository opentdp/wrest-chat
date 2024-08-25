package wclient

import (
	"encoding/json"
	"net"
	"net/url"
	"strings"

	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/request"
	"github.com/opentdp/go-helper/strutil"

	"github.com/opentdp/wrest-chat/args"
	"github.com/opentdp/wrest-chat/dbase/message"
	"github.com/opentdp/wrest-chat/dbase/setting"
	"github.com/opentdp/wrest-chat/wcferry"
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
		SdkLibrary: args.Wcf.SdkLibrary,
	}

	// 初始化连接
	logman.Warn("wcf starting ...")
	if err := wc.Connect(); err != nil {
		logman.Fatal("wcf start failed", "error", err)
	}

	// 打印收到的消息
	if args.Wcf.MsgPrint {
		wc.EnrollReceiver(true, wcferry.WxMsgPrinter)
	}

	// 存储收到的消息
	if args.Wcf.MsgStore {
		wc.EnrollReceiver(true, msgToDatabase)
		if (args.Wcf.MsgStoreDays) > 0 {
			message.Shrink(args.Wcf.MsgStoreDays)
		}
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

	// 验证参数
	msg = strings.TrimSpace(strings.Trim(msg, "-"))
	wxid = strings.TrimSpace(strings.Trim(wxid, "-"))
	roomid = strings.TrimSpace(strings.Trim(roomid, "-"))
	if msg == "" || (wxid == "" && roomid == "") {
		return -1
	}

	// 重组参数
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
		if p := strings.Split(msg, "\n")[1:]; len(p) > 5 {
			return wc.CmdClient.SendRichText(p[0], p[1], p[2], p[3], p[4], p[5], receiver)
		}
	}

	// 发送文本信息
	if ater != "" {
		if u := wc.CmdClient.GetInfoByWxid(ater); u != nil && u.Name != "" {
			msg = "@" + u.Name + "\n" + msg
		}
	}
	return wc.CmdClient.SendTxt(msg, receiver, ater)

}

// 使用接口回复消息
// param url string 待请求的接口地址
// param wxid string 消息接收人，如果 roomid 存在则为 at 此人
// param roomid string 消息接收群，空则为私聊
// return string 成功返回空，失败返回错误信息或结果
func ApiRequestMsg(url, wxid, roomid string) int32 {

	self := wc.CmdClient.GetSelfInfo()

	// 验证参数
	if self == nil {
		logman.Error("ApiRequestMsg::GetSelfInfo failed")
		return -1
	}

	// 获取结果
	res, err := request.TextGet(url, request.H{
		"User-Agent": args.AppName + "/" + args.Version,
		"Client-Uid": self.Wxid + "," + wxid,
	})

	if err != nil {
		if res == "" {
			res = err.Error()
		}
		return SendFlexMsg(res, wxid, roomid)
	}

	// 解析结果
	data := &ApiResponse{}
	if json.Unmarshal([]byte(res), data) != nil {
		return SendFlexMsg(res, wxid, roomid)
	}

	// 构造消息
	text := ""
	switch data.Type {
	case "card":
		if data.Card.Name == "" {
			data.Card.Name = self.Name
		}
		if data.Card.Account == "" {
			data.Card.Account = self.Wxid
		}
		if data.Card.Icon == "" {
			data.Card.Icon = setting.ApiEndpointIcon
		}
		text += "card\n"
		text += data.Card.Name + "\n"
		text += data.Card.Account + "\n"
		text += data.Card.Title + "\n"
		text += data.Card.Digest + "\n"
		text += data.Card.Link + "\n"
		text += data.Card.Icon
	case "file":
		text += data.Link
	case "image":
		text += data.Link
	default:
		text += data.Text
	}

	// 发送消息
	return SendFlexMsg(text, wxid, roomid)

}

// 接口响应结果
type ApiResponse struct {
	Type string   `json:"type"` // 数据类型 [card, file, image, text, error]
	Card struct { // 当 type 为 card 时有效
		Name    string `json:"name"`    // 左下显示的名字，可选
		Account string `json:"account"` // 公众号 id，可显示对应的头像，可选
		Title   string `json:"title"`   // 标题，最多显示为两行
		Digest  string `json:"digest"`  // 摘要，最多显示为三行
		Link    string `json:"link"`    // 点击后跳转的链接
		Icon    string `json:"icon"`    // 右侧缩略图的链接，可选
	} `json:"card,omitempty"`
	Link string `json:"link,omitempty"` // 当 type 为 file 或 image 时有效
	Text string `json:"text,omitempty"` // 当 type 为 text 或 error 时有效
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
