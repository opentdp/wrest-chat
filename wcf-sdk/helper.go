package wcf

import (
	"bytes"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/onquit"
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
	// 解析地址
	parts := strings.Split(l.Address, ":")
	port, _ := strconv.Atoi(parts[1])

	// 启动 wcf 服务
	if l.Wcfexe != "" {
		l.injectWechat(port)
		time.Sleep(3 * time.Second)
	}

	// 连接 wcf 服务
	l.client = &Client{
		CmdServer: "tcp://" + l.Address,
		MsgServer: "tcp://" + parts[0] + ":" + strconv.Itoa(port+1),
	}
	logman.Info("wcf connect", "server", l.Address)
	return l.client, l.client.dial()
}

func (l *Launcher) AutoDestory() {
	onquit.Register(func() {
		// 关闭 wcf 连接
		l.client.Close()

		// 关闭 wcf 服务
		if l.Wcfexe != "" {
			logman.Info("killing wechat process")
			cmd := exec.Command("taskkill", "/IM", "WeChat.exe", "/F")
			if err := cmd.Run(); err != nil {
				logman.Warn("failed to kill wechat", "error", err)
			}
		}
	})
}

func (l *Launcher) injectWechat(port int) {
	var cmd *exec.Cmd

	// 检查 wechat 是否已经启动
	var out bytes.Buffer
	cmd = exec.Command("tasklist")
	cmd.Stdout = &out
	if strings.Contains(out.String(), "WeChat") {
		logman.Fatal("please close wechat first")
	}

	// 启动 wcf 并注入 wechat
	logman.Info("start wcf", "bin", l.Wcfexe, "port", port)
	cmd = exec.Command(l.Wcfexe, "start", strconv.Itoa(port))
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	if err := cmd.Start(); err != nil {
		logman.Fatal("failed to inject wecaht", err)
	}
}
