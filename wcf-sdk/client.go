package wcf

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/onquit"
)

type Client struct {
	WcfPath   string     // sdk.dll 路径
	WcfAddr   string     // wcf 监听地址
	WcfPort   int        // wcf 监听端口
	CmdClient *CmdClient // 命令客户端
	MsgClient *MsgClient // 消息客户端
}

// 启动 wcf 服务
// return error 错误信息
func (c *Client) Connect() error {
	// 设置默认值
	if c.WcfAddr == "" {
		c.WcfAddr = "127.0.0.1"
	}
	if c.WcfPort == 0 {
		c.WcfPort = 10080
	}
	// 启动 wcf 服务
	if c.WcfPath != "" {
		if err := c.injectWechat(c.WcfPort); err != nil {
			return err
		}
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
// param pyq bool 是否接收朋友圈消息
// param fn ...MsgCallback 消息回调函数
// return error 错误信息
func (c *Client) EnrollReceiver(pyq bool, fn ...MsgCallback) error {
	if c.CmdClient.EnableMsgServer(true) != 0 {
		return errors.New("failed to enable msg server")
	}
	time.Sleep(1 * time.Second)
	c.MsgClient.Register(fn...)
	return nil
}

// 关闭消息接收器
// return error 错误信息
func (c *Client) DisableReceiver() error {
	if c.CmdClient.DisableMsgServer() != 0 {
		return errors.New("failed to disable msg server")
	}
	return c.MsgClient.Close()
}

// 构建地址
// param ip string IP地址
// param port int 端口
// return string IP地址和端口
func (c *Client) buildAddr(ip string, port int) string {
	if strings.Contains(ip, ":") {
		return fmt.Sprintf("tcp://[%s]:%d", ip, port)
	} else {
		return fmt.Sprintf("tcp://%s:%d", ip, port)
	}
}

// 启动 wcf 服务并注入 wechat
// param port int wcf 服务端口
// return error 错误信息
func (c *Client) injectWechat(port int) error {
	// 检查 wechat 状态
	var out bytes.Buffer
	cmd := exec.Command("tasklist")
	cmd.Stdout = &out
	if strings.Contains(out.String(), "WeChat") {
		return errors.New("please close wechat first")
	}
	// 加载 sdk.dll 动态库
	sdk, err := syscall.LoadDLL(c.WcfPath)
	if err != nil {
		logman.Info("load sdk.dll failed", "error", err)
		return err
	}
	defer sdk.Release()
	// 查找 WxInitSDK 函数
	wxInitSDK, err := sdk.FindProc("WxInitSDK")
	if err != nil {
		logman.Info("call wxInitSDK failed", "error", err)
		return err
	}
	// 初始化 WeChatFerry 服务
	logman.Info("init Wcf service", "listen", port)
	r1, r2, err := wxInitSDK.Call(uintptr(0), uintptr(port))
	logman.Warn("wxInitSDK", "r1", r1, "r2", r2, "error", err)
	return nil
}
