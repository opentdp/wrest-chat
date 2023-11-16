package wcf

import (
	"github.com/opentdp/go-helper/logman"
)

type MsgClient struct {
	pbSocket       // RPC 客户端
	receiving bool // 接收消息中
	callbacks []MsgCallback
}

// 消息回调函数
type MsgCallback func(msg *WxMsg)

// 关闭 RPC 连接
// @return error 错误信息
func (c *MsgClient) Close() error {
	c.callbacks = []MsgCallback{}
	c.receiving = false
	return c.close()
}

// 创建消息接收器
// @param fn ...MsgCallback 消息回调函数
func (c *MsgClient) Register(fn ...MsgCallback) {
	c.callbacks = append(c.callbacks, fn...)
	if !c.receiving {
		go c.listener()
	}
}

// 消息接收器循环
// @return error 错误信息
func (c *MsgClient) listener() error {
	// 连接消息服务
	c.receiving = true
	if err := c.dial(); err != nil {
		logman.Warn("msg receiver", "error", err)
		c.receiving = false
		return err
	}
	// 开始接收消息
	for c.receiving {
		resp, err := c.recv()
		if err != nil {
			logman.Warn("msg receiver", "error", err)
		}
		for _, f := range c.callbacks {
			go f(resp.GetWxmsg())
		}
	}
	// 关闭连接
	return c.Close()
}
