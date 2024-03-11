package wcferry

import (
	"errors"
	"net"
	"strconv"
	"time"

	"go.nanomsg.org/mangos/v3"
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
// param ip string 服务器地址
// param port int 服务器端口
// return *pbSocket 客户端
func newPbSocket(ip string, port int) *pbSocket {
	addr := net.JoinHostPort(ip, strconv.Itoa(port))
	return &pbSocket{server: "tcp://" + addr}
}

// 连接服务器
// param d uint 读写超时时间(s)
// return error 错误信息
func (c *pbSocket) init(d uint) (err error) {
	// 创建连接
	all.AddTransports(nil)
	if c.socket, err = pair1.NewSocket(); err != nil {
		return err
	}
	// 设置参数
	if d > 0 {
		t := time.Duration(d) * time.Second
		c.socket.SetOption(mangos.OptionRecvDeadline, t)
		c.socket.SetOption(mangos.OptionSendDeadline, t)
	}
	c.socket.SetOption(mangos.OptionMaxRecvSize, 16*1024*1024)
	// 连接服务器
	logman.Warn("pbSocket dial", "server", c.server)
	return c.socket.Dial(c.server)
}

// 调用接口
// param req *Request 请求参数
// return *Response 响应参数
func (c *pbSocket) call(req *Request) *Response {
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
// param req *Request 请求参数
// return error 错误信息
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
// return *Response 响应参数
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
