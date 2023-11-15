package wcf

import (
	"github.com/opentdp/go-helper/logman"
)

type MsgReceiver struct {
	callbacks []MsgCallback
	receiving bool // 正在接收消息
	pbSocket       // RPC 客户端
}

type MsgCallback func(msg *WxMsg)

func (c *MsgReceiver) Enroll(f MsgCallback) error {
	c.callbacks = append(c.callbacks, f)
	if !c.receiving {
		return c.looper()
	}
	return nil
}

func (c *MsgReceiver) Disable() error {
	c.callbacks = []MsgCallback{}
	if c.receiving {
		c.receiving = false
		return c.close()
	}
	return nil
}

func (c *MsgReceiver) looper() error {
	// 连接消息服务
	c.receiving = true
	if err := c.dial(); err != nil {
		logman.Warn("MsgReceiver", "error", err)
		c.receiving = false
		return err
	}
	// 开始接收消息
	for c.receiving {
		resp, err := c.recv()
		if err != nil {
			logman.Warn("MsgReceiver", "error", err)
		}
		for _, f := range c.callbacks {
			go f(resp.GetWxmsg())
		}
	}
	// 关闭连接
	return c.Disable()
}
