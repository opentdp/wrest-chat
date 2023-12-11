package wcferry

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"go.nanomsg.org/mangos"
	"go.nanomsg.org/mangos/v3/protocol"
	"go.nanomsg.org/mangos/v3/protocol/pair1"
	"go.nanomsg.org/mangos/v3/transport/all"
	"google.golang.org/protobuf/proto"

	"github.com/opentdp/go-helper/logman"
)

type pbSocket struct {
	server string // 服务端
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
	if c.socket, err = pair1.NewSocket(); err == nil {
		return c.socket.Dial(c.server)
	}
	return err
}

// 读写超时
func (c *pbSocket) deadline(d uint) {
	if c.socket != nil {
		t := time.Duration(d) * time.Second
		c.socket.SetOption(mangos.OptionRecvDeadline, t)
		c.socket.SetOption(mangos.OptionSendDeadline, t)
	}
}

// 调用接口
func (c *CmdClient) call(req *Request) *Response {
	if err := c.send(req); err != nil {
		logman.Error(err.Error())
		return &Response{}
	}
	if resp, err := c.recv(); err != nil {
		logman.Error(err.Error())
		return &Response{}
	} else {
		return resp
	}
}

// 发送数据
func (c *pbSocket) send(req *Request) error {
	if c.socket == nil {
		return errors.New("socket is nil")
	}
	data, err := proto.Marshal(req)
	if err == nil {
		return c.socket.Send(data)
	}
	return err
}

// 接收数据
func (c *pbSocket) recv() (*Response, error) {
	resp := &Response{}
	if c.socket == nil {
		return resp, errors.New("socket is nil")
	}
	data, err := c.socket.Recv()
	if err == nil {
		err = proto.Unmarshal(data, resp)
	}
	return resp, err
}

// 关闭连接
func (c *pbSocket) close() error {
	if c.socket == nil {
		return errors.New("socket is nil")
	}
	return c.socket.Close()
}
