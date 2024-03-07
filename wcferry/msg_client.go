package wcferry

import (
	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/strutil"
)

type MsgClient struct {
	*pbSocket                        // RPC 客户端
	callbacks map[string]MsgCallback // 消息回调列表
}

// 消息回调函数
type MsgCallback func(msg *WxMsg)

// 关闭 RPC 连接
// param ks 消息接收器标识，空则关闭所有
// return error 错误信息
func (c *MsgClient) Destroy(ks ...string) error {
	if len(c.callbacks) > 0 && len(ks) > 0 {
		for _, k := range ks {
			delete(c.callbacks, k)
		}
		if len(c.callbacks) > 0 {
			return nil
		}
	}
	// 关闭消息推送
	c.callbacks = nil
	return c.close()
}

// 创建消息接收器
// param cb MsgCallback 消息回调函数
// return string 接收器唯一标识
func (c *MsgClient) Register(cb MsgCallback) (string, error) {
	k := strutil.Rand(16)
	if c.callbacks == nil {
		if err := c.init(0); err != nil {
			logman.Error("msg receiver", "error", err)
			return "", err
		}
		c.callbacks = map[string]MsgCallback{
			k: cb,
		}
		go func() {
			defer c.Destroy()
			for len(c.callbacks) > 0 {
				if resp, err := c.recv(); err == nil {
					msg := resp.GetWxmsg()
					for _, f := range c.callbacks {
						go f(msg) // 异步处理
					}
				} else {
					logman.Error("msg receiver", "error", err)
				}
			}
			logman.Warn("msg receiver stopped")
		}()
	} else {
		c.callbacks[k] = cb
	}
	return k, nil
}
