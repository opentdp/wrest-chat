package wcferry

import (
	"fmt"
	"strings"

	"github.com/opentdp/go-helper/logman"
)

type MsgClient struct {
	*pbSocket      // RPC 客户端
	receiving bool // 接收消息中
	callbacks []MsgCallback
}

// 消息回调函数
type MsgCallback func(msg *WxMsg)

// 关闭 RPC 连接
// return error 错误信息
func (c *MsgClient) Close() error {
	c.callbacks = []MsgCallback{}
	c.receiving = false
	return c.close()
}

// 创建消息接收器
// param fn ...MsgCallback 消息回调函数
func (c *MsgClient) Register(fn ...MsgCallback) {
	c.callbacks = append(c.callbacks, fn...)
	if !c.receiving {
		go c.listener()
	}
}

// 消息接收器循环
// return error 错误信息
func (c *MsgClient) listener() error {
	// 连接消息服务
	c.receiving = true
	if err := c.dial(); err != nil {
		logman.Error("msg receiver", "error", err)
		c.receiving = false
		return err
	}
	c.deadline(2000)
	// 开始接收消息
	for c.receiving {
		resp, err := c.recv()
		if err != nil {
			logman.Error("msg receiver", "error", err)
		}
		for _, f := range c.callbacks {
			go f(resp.GetWxmsg())
		}
	}
	// 关闭连接
	return c.Close()
}

// 打印接收到的消息
// param msg *WxMsg 消息
func MsgPrinter(msg *WxMsg) {
	fmt.Print("=== New Message ===\n")
	if msg.Id > 0 {
		fmt.Printf("<<Id>> %d\n", msg.Id)
	}
	if msg.Type > 0 {
		fmt.Printf("<<Type>> %d\n", msg.Type)
	}
	if msg.Roomid != "" {
		fmt.Printf("<<Roomid>> %s\n", msg.Roomid)
	}
	if msg.Sender != "" {
		fmt.Printf("<<Sender>> %v\n", msg.Sender)
	}
	if msg.Content != "" {
		fmt.Printf("<<Content>> %s\n", msg.Content)
	}
	if msg.Extra != "" {
		fmt.Printf("<<Extra>> %s\n", strings.TrimSpace(msg.Extra))
	}
	fmt.Print("=== End Message ===\n")
}
