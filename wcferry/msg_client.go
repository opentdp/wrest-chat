package wcferry

import (
	"errors"

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
func (c *MsgClient) Destroy(force bool) error {
	if !force && len(c.callbacks) > 0 {
		return errors.New("callbacks not empty")
	}
	c.callbacks = []MsgCallback{}
	c.receiving = false
	return c.close()
}

// 创建消息接收器
// param fn ...MsgCallback 消息回调函数
func (c *MsgClient) Register(fn ...MsgCallback) {
	c.callbacks = append(c.callbacks, fn...)
	if !c.receiving {
		c.receiving = true
		go c.listener()
	}
}

// 消息接收器循环
func (c *MsgClient) listener() {
	defer c.Destroy(true)
	// 连接消息服务
	if err := c.init(0); err != nil {
		logman.Error("msg receiver", "error", err)
		return
	}
	// 开始接收消息
	for c.receiving {
		if resp, err := c.recv(); err == nil {
			res := resp.GetWxmsg()
			for _, f := range c.callbacks {
				go f(res)
			}
		} else {
			logman.Error("msg receiver", "error", err)
		}
	}
	logman.Warn("msg receiver stopped")
}
