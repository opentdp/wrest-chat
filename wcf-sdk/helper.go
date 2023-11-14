package wcf

import (
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/opentdp/go-helper/logman"
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
		logman.Info("wcf start", "bin", l.Wcfexe, "port", port)
		cmd := exec.Command(l.Wcfexe, "start", strconv.Itoa(port))
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
		if err := cmd.Start(); err != nil {
			return nil, err
		}
		time.Sleep(5 * time.Second)
	}

	l.client = &Client{
		CmdServer: "tcp://" + l.Address,
		MsgServer: "tcp://" + parts[0] + ":" + strconv.Itoa(port+1),
	}

	logman.Info("wcf connect", "CmdServer", l.client.CmdServer)
	return l.client, l.client.dial()
}

func (l *Launcher) AutoStop() error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if l.Wcfexe != "" {
		logman.Info("wcf stop", "bin", l.Wcfexe)
		cmd := exec.Command(l.Wcfexe, "stop")
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
		cmd.Run()
	}

	logman.Info("wcf disconnect", "CmdServer", l.client.CmdServer)
	return l.client.Close()
}
