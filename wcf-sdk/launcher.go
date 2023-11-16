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
)

type Launcher struct {
	Address string  // RPC 监听地址
	Wcfexe  string  // wcf.exe 路径
	client  *Client // wcf 客户端
}

// 启动 wcf 服务
//
// Returns:
//
// *Client: wcf 客户端
// error: 错误信息
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
		time.Sleep(5 * time.Second)
	}
	// 连接 wcf 服务
	l.client = &Client{
		pbSocket: pbSocket{
			Server: "tcp://" + l.Address,
		},
		Receiver: &MsgReceiver{
			pbSocket: pbSocket{
				Server: "tcp://" + parts[0] + ":" + strconv.Itoa(port+1),
			},
		},
	}
	return l.client, l.client.dial()
}

// 自动销毁 wcf 服务
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

// 启动 wcf 服务并注入 wechat
//
// Args:
//
// port int: wcf 服务端口
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
