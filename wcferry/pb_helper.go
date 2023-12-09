package wcferry

import (
	"errors"
	"fmt"
	"strings"

	"go.nanomsg.org/mangos"
	"go.nanomsg.org/mangos/v3/protocol"
	"go.nanomsg.org/mangos/v3/protocol/pair1"
	"go.nanomsg.org/mangos/v3/transport/all"
	"google.golang.org/protobuf/proto"

	"github.com/opentdp/go-helper/logman"
)

// 通用消息

type cmdMsg struct {
	*Request
}

// 生成消息
func (c *cmdMsg) build() []byte {
	marshal, _ := proto.Marshal(c)
	return marshal
}

// 生成请求消息
func genFunReq(fun Functions) *cmdMsg {
	return &cmdMsg{
		&Request{Func: fun, Msg: nil},
	}
}

// RPC 客户端

type pbSocket struct {
	server string // 接口地址
	socket protocol.Socket
}

// 创建客户端
func newPbSocket(ip string, port int) *pbSocket {
	var addr string
	if strings.Contains(ip, ":") {
		addr = fmt.Sprintf("tcp://[%s]:%d", ip, port)
	} else {
		addr = fmt.Sprintf("tcp://%s:%d", ip, port)
	}
	return &pbSocket{server: addr}
}

// 连接服务器
func (c *pbSocket) dial() (err error) {
	all.AddTransports(nil) // 注册所有传输协议
	logman.Info("pbSocket dial", "server", c.server)
	c.socket, err = pair1.NewSocket()
	if err == nil {
		return c.socket.Dial(c.server)
	}
	return err
}

// 设置超时时间
func (c *pbSocket) deadline(d int) {
	c.socket.SetOption(mangos.OptionRecvDeadline, d)
	c.socket.SetOption(mangos.OptionSendDeadline, d)
}

// 调用接口
func (c *CmdClient) call(data []byte) *Response {
	if err := c.send(data); err != nil {
		logman.Error(err.Error())
	}
	recv, err := c.recv()
	if err != nil {
		logman.Error(err.Error())
	}
	return recv
}

// 接收数据
func (c *pbSocket) recv() (*Response, error) {
	resp := &Response{}
	if c.socket == nil {
		return resp, errors.New("socket is nil")
	}
	recv, err := c.socket.Recv()
	if err == nil {
		err = proto.Unmarshal(recv, resp)
	}
	return resp, err
}

// 发送数据
func (c *pbSocket) send(data []byte) error {
	if c.socket == nil {
		return errors.New("socket is nil")
	}
	return c.socket.Send(data)
}

// 关闭连接
func (c *pbSocket) close() error {
	if c.socket == nil {
		return errors.New("socket is nil")
	}
	return c.socket.Close()
}
