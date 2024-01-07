package wcferry

import (
	"errors"
	"strings"

	"github.com/clbanning/mxj"
	"github.com/opentdp/go-helper/logman"
)

type MsgClient struct {
	*pbSocket               // RPC 客户端
	receiving bool          // 接收消息中
	callbacks []MsgCallback // 消息回调函数
}

// 消息回调参数
type MsgPayload struct {
	*WxMsg      // 消息原始数据
	Content any `json:"content,omitempty"`
	Xml     any `json:"xml,omitempty"`
}

// 消息回调函数
type MsgCallback func(msg *MsgPayload)

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
func (c *MsgClient) Register(fn ...MsgCallback) error {
	if !c.receiving {
		// 连接消息服务
		if err := c.init(0); err != nil {
			logman.Error("msg receiver", "error", err)
			return err
		}
		// 开始接收消息
		c.receiving = true
		go func() {
			defer c.Destroy(true)
			for c.receiving {
				if resp, err := c.recv(); err == nil {
					msg := c.MsgXmlToMap(resp.GetWxmsg())
					for _, f := range c.callbacks {
						go f(msg)
					}
				} else {
					logman.Error("msg receiver", "error", err)
				}
			}
			logman.Warn("msg receiver stopped")
		}()
	}
	// 注册回调函数
	c.callbacks = append(c.callbacks, fn...)
	return nil
}

// 转换消息中的XML为Map
// param msg *WxMsg 消息
// return *MsgPayload 转换后的消息
func (c *MsgClient) MsgXmlToMap(msg *WxMsg) *MsgPayload {
	var str string
	var ret = &MsgPayload{msg, msg.Content, msg.Xml}
	// c.Xml
	str = strings.TrimSpace(msg.Xml)
	if strings.HasPrefix(str, "<") {
		mv, err := mxj.NewMapXml([]byte(str))
		if err == nil {
			ret.Xml = mv
		}
	}
	// c.Content
	str = strings.TrimSpace(msg.Content)
	xmlPrefixes := []string{"<?xml", "<sysmsg", "<msg"}
	for _, prefix := range xmlPrefixes {
		if strings.HasPrefix(str, prefix) {
			mv, err := mxj.NewMapXml([]byte(str))
			if err == nil {
				ret.Content = mv
			}
			break
		}
	}
	// return
	return ret
}
