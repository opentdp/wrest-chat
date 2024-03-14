package wcferry

import (
	"sync"

	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/strutil"
)

type MsgClient struct {
	*pbSocket             // RPC 客户端
	consumer  MapConsumer // 消费者
	mu        sync.Mutex  // 互斥锁
}

type MsgConsumer func(msg *WxMsg)
type MapConsumer map[string]MsgConsumer

// 关闭 RPC 连接
// param ks 消息接收器标识，空则关闭所有
// return error 错误信息
func (c *MsgClient) Destroy(ks ...string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.consumer) > 0 && len(ks) > 0 {
		for _, k := range ks {
			delete(c.consumer, k)
		}
		if len(c.consumer) > 0 {
			return nil
		}
	}
	// 关闭消息推送
	c.consumer = nil
	return c.close()
}

// 创建消息接收器
// param cb MsgConsumer 消息回调函数
// return string 接收器唯一标识
func (c *MsgClient) Register(cb MsgConsumer) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	k := strutil.Rand(16)
	if c.consumer == nil {
		if err := c.init(0); err != nil {
			logman.Error("msg consumer", "error", err)
			return "", err
		}
		c.consumer = MapConsumer{k: cb}
		go c.runner()
	} else {
		c.consumer[k] = cb
	}
	return k, nil
}

// 消息推送执行者
func (c *MsgClient) runner() {
	defer c.Destroy()
	// 接收消息
	for len(c.consumer) > 0 {
		if resp, err := c.recv(); err == nil {
			msg := resp.GetWxmsg()
			for _, f := range c.consumer {
				go f(msg) // 异步推送
			}
		} else {
			logman.Error("msg consumer", "error", err)
		}
	}
	// 连接断开
	logman.Warn("msg consumer stopped")
}
