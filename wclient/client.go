package wclient

import (
	"net"

	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/strutil"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/dbase/message"
	"github.com/opentdp/wechat-rest/wcferry"
)

var wc *wcferry.Client

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
		WeChatAuto: args.Wcf.WeChatAuto,
	}

	// 初始化连接
	logman.Info("wcf starting ...")
	if err := wc.Connect(); err != nil {
		logman.Fatal("failed to start wcf", "error", err)
	}

	// 存储收到的消息
	if args.Wcf.MsgBackup {
		wc.EnrollReceiver(true, msgToDatabase)
	}

	// 打印收到的消息
	if args.Wcf.MsgPrinter {
		wc.EnrollReceiver(true, wcferry.WxMsgPrinter)
	}

	return wc

}

func msgToDatabase(msg *wcferry.WxMsg) {

	rq := message.CreateParam{
		Rd:      0,
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
