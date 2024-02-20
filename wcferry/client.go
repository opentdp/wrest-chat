package wcferry

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/opentdp/go-helper/filer"
	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/onquit"
)

type Client struct {
	WcfBinary  string     // wcf.exe 路径
	ListenAddr string     // wcf 监听地址
	ListenPort int        // wcf 监听端口
	WeChatAuto bool       // 自动启停微信
	CmdClient  *CmdClient // 命令客户端
	MsgClient  *MsgClient // 消息客户端
}

// 启动 wcf 服务
// return error 错误信息
func (c *Client) Connect() error {
	if c.ListenAddr == "" {
		c.ListenAddr = "127.0.0.1"
	}
	if c.ListenPort == 0 {
		c.ListenPort = 10086
	}
	// 注册 wcf 服务
	c.CmdClient = &CmdClient{
		pbSocket: newPbSocket(c.ListenAddr, c.ListenPort),
	}
	c.MsgClient = &MsgClient{
		pbSocket: newPbSocket(c.ListenAddr, c.ListenPort+1),
	}
	// 启动 wcf 服务
	if err := c.wxInitSDK(); err != nil {
		return err
	}
	// 自动注销 wcf
	onquit.Register(func() {
		c.CmdClient.Destroy()
		c.MsgClient.Destroy()
		c.wxDestroySDK()
	})
	// 返回连接结果
	return c.CmdClient.init(5)
}

// 启动消息接收器
// param pyq bool 是否接收朋友圈消息
// param cb MsgCallback 消息回调函数，可选参数
// return string 接收器唯一标识
func (c *Client) EnrollReceiver(pyq bool, cb MsgCallback) (string, error) {
	if c.MsgClient.callbacks == nil {
		if c.CmdClient.EnableMsgReciver(true) != 0 {
			return "", errors.New("failed to enable msg server")
		}
	}
	time.Sleep(1 * time.Second)
	return c.MsgClient.Register(cb)
}

// 关闭消息接收器
// param sk 消息接收器标识，为空则关闭所有
// return error 错误信息
func (c *Client) DisableReceiver(ks ...string) error {
	err := c.MsgClient.Destroy(ks...)
	if c.MsgClient.callbacks == nil {
		if c.CmdClient.DisableMsgReciver() != 0 {
			return errors.New("failed to disable msg server")
		}
	}
	return err
}

// 启动 wcf 服务
// return error 错误信息
func (c *Client) wxInitSDK() error {
	if c.WcfBinary == "" {
		return nil
	}
	// 尝试自动启动微信
	if c.WeChatAuto {
		out, _ := exec.Command("tasklist").Output()
		if strings.Contains(string(out), "WeChat.exe") {
			return errors.New("please close wechat")
		}
	}
	// 查找 wcf.exe 路径
	if !filer.Exists(c.WcfBinary) {
		if filer.Exists("wcferry/wcf.exe") {
			c.WcfBinary = "wcferry/wcf.exe"
		} else if filer.Exists("wcferry/bin/wcf.exe") {
			c.WcfBinary = "wcferry/bin/wcf.exe"
		} else {
			return errors.New("wcf.exe not found")
		}
	}
	// 注入微信，打开 wcf 服务
	logman.Info(c.WcfBinary + " start " + strconv.Itoa(c.ListenPort))
	cmd := exec.Command(c.WcfBinary, "start", strconv.Itoa(c.ListenPort))
	return cmd.Run()
}

// 关闭 wcf 服务
// return error 错误信息
func (c *Client) wxDestroySDK() error {
	if c.WcfBinary == "" {
		return nil
	}
	// 关闭 wcf 服务
	logman.Info(c.WcfBinary + " stop")
	cmd := exec.Command(c.WcfBinary, "stop")
	err := cmd.Run()
	// 尝试自动关闭微信
	if err == nil && c.WeChatAuto {
		logman.Info("killing wechat process")
		cmd := exec.Command("taskkill", "/IM", "WeChat.exe", "/F")
		if err := cmd.Run(); err != nil {
			logman.Warn("failed to kill wechat", "error", err)
			return err
		}
	}
	return err
}
