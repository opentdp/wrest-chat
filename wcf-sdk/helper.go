package wcf

import (
	"os"
	"os/exec"
	"strconv"
	"strings"

	"google.golang.org/protobuf/proto"
)

// 通用消息

type cmdMsg struct {
	*Request
}

func (c *cmdMsg) build() []byte {
	marshal, _ := proto.Marshal(c)
	return marshal
}

func genFunReq(fun Functions) *cmdMsg {
	return &cmdMsg{
		&Request{Func: fun, Msg: nil},
	}
}

// 服务启动器

type Launcher struct {
	client *Client

	Address string // RPC 监听地址
	Wcfexe  string // wcf.exe 路径
}

func (l *Launcher) Start() (*Client, error) {
	if l.Address == "" {
		l.Address = "127.0.0.1:10080"
	}

	parts := strings.Split(l.Address, ":")
	port, _ := strconv.Atoi(parts[1])

	if l.Wcfexe != "" {
		cmd := exec.Command(l.Wcfexe, "start", strconv.Itoa(port))
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
		if err := cmd.Start(); err != nil {
			return nil, err
		}
	}

	l.client = &Client{
		CmdServer: "tcp://" + l.Address,
		MsgServer: "tcp://" + parts[0] + ":" + strconv.Itoa(port+1),
	}
	return l.client, l.client.dial()
}

func (l *Launcher) Stop() error {
	if l.Wcfexe != "" {
		cmd := exec.Command(l.Wcfexe, "stop")
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
		cmd.Run()
	}

	return l.client.Close()
}
