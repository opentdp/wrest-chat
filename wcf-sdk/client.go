package wcf

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/onquit"
)

type Client struct {
	WcfPath   string     // wcf.exe 路径
	WcfAddr   string     // wcf 监听地址
	WcfPort   int        // wcf 监听端口
	CmdClient *CmdClient // 命令客户端
	MsgClient *MsgClient // 消息客户端
}

// 启动 wcf 服务
//
// Returns:
//
// *CmdClient: wcf 客户端
// error: 错误信息
func (c *Client) Start() error {
	if c.WcfAddr == "" {
		c.WcfAddr = "127.0.0.1"
	}
	if c.WcfPort == 0 {
		c.WcfPort = 10080
	}
	// 启动 wcf 服务
	if c.WcfPath != "" {
		c.injectWechat(c.WcfPort)
		time.Sleep(5 * time.Second)
	}
	// 连接 wcf 服务
	c.CmdClient = &CmdClient{
		pbSocket: pbSocket{server: c.buildAddr(c.WcfAddr, c.WcfPort)},
	}
	c.MsgClient = &MsgClient{
		pbSocket: pbSocket{server: c.buildAddr(c.WcfAddr, c.WcfPort+1)},
	}
	return c.CmdClient.dial()
}

// 自动销毁 wcf 服务
func (c *Client) AutoDestory() {
	onquit.Register(func() {
		// 关闭 wcf 连接
		c.MsgClient.Close()
		c.CmdClient.Close()
		// 关闭 wcf 服务
		if c.WcfPath != "" {
			logman.Info("killing wechat process")
			cmd := exec.Command("taskkill", "/IM", "WeChat.exe", "/F")
			if err := cmd.Run(); err != nil {
				logman.Warn("failed to kill wechat", "error", err)
			}
		}
	})
}

// 启动消息接收器
//
// Args:
//
// pyq bool: 是否接收朋友圈消息
// fn ...MsgCallback: 消息回调函数
//
// Returns:
//
// int32: 0 为成功，其他失败
func (c *Client) EnrollReceiver(pyq bool, fn ...MsgCallback) int32 {
	status := c.CmdClient.EnableMsgServer(true)
	c.MsgClient.Register(fn...)
	return status
}

// 关闭消息接收器
//
// Returns:
//
// int32: 0 为成功，其他失败
func (c *Client) DisableReceiver() int32 {
	status := c.CmdClient.DisableMsgServer()
	c.MsgClient.Close()
	return status
}

// 构建地址
//
// Args:
//
// ip string: IP地址
// port int: 端口
//
// Returns:
//
// string: IP地址和端口
func (c *Client) buildAddr(ip string, port int) string {
	if strings.Contains(ip, ":") {
		return fmt.Sprintf("tcp://[%s]:%d", ip, port)
	} else {
		return fmt.Sprintf("tcp://%s:%d", ip, port)
	}
}

// 启动 wcf 服务并注入 wechat
//
// Args:
//
// port int: wcf 服务端口
func (c *Client) injectWechat(port int) {
	var cmd *exec.Cmd
	// 检查 wechat 是否已经启动
	var out bytes.Buffer
	cmd = exec.Command("tasklist")
	cmd.Stdout = &out
	if strings.Contains(out.String(), "WeChat") {
		logman.Fatal("please close wechat first")
	}
	// 启动 wcf 并注入 wechat
	logman.Info("start wcf", "bin", c.WcfPath, "port", port)
	cmd = exec.Command(c.WcfPath, "start", strconv.Itoa(port))
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	if err := cmd.Start(); err != nil {
		logman.Fatal("failed to inject wecaht", err)
	}
}
