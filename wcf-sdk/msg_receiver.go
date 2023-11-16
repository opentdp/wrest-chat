package wcf

import (
	"github.com/opentdp/go-helper/logman"
)

type MsgReceiver struct {
	callbacks []MsgCallback
	receiving bool // 正在接收消息
	pbSocket       // RPC 客户端
}

// 消息回调函数
type MsgCallback func(msg *WxMsg)

// 创建消息接收器
//
// Args:
//
// fn ...MsgCallback: 消息回调函数
//
// Returns:
//
// error: 错误信息
func (c *MsgReceiver) Enroll(fn ...MsgCallback) error {
	c.callbacks = append(c.callbacks, fn...)
	if !c.receiving {
		return c.looper()
	}
	return nil
}

// 关闭消息接收器
//
// Returns:
//
// error: 错误信息
func (c *MsgReceiver) Disable() error {
	c.callbacks = []MsgCallback{}
	if c.receiving {
		c.receiving = false
		return c.close()
	}
	return nil
}

// 消息接收器循环
//
// Returns:
//
// error: 错误信息
func (c *MsgReceiver) looper() error {
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
	return c.Disable()
}
