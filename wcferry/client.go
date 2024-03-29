package wcferry

import (
	"errors"
	"os/exec"
	"strconv"
	"time"

	"github.com/opentdp/go-helper/filer"
	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/onquit"
)

const Wcf_Version = "39.0.14"
const Wechat_Version = "3.9.2.23"

type Client struct {
	WcfBinary  string     // wcf.exe 路径
	ListenAddr string     // wcf 监听地址
	ListenPort int        // wcf 监听端口
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
// param cb MsgConsumer 消息回调函数，可选参数
// return string 接收器唯一标识
func (c *Client) EnrollReceiver(pyq bool, cb MsgConsumer) (string, error) {
	if c.MsgClient.consumer == nil {
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
	if c.MsgClient.consumer == nil {
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
	// 尝试在子目录查找
	if !filer.Exists(c.WcfBinary) {
		if !filer.Exists("wcferry/" + c.WcfBinary) {
			return errors.New(c.WcfBinary + " not found")
		}
		c.WcfBinary = "wcferry/" + c.WcfBinary
	}
	// 打开 wcf 服务程序
	port := strconv.Itoa(c.ListenPort)
	logman.Warn(c.WcfBinary + " start " + port)
	cmd := exec.Command(c.WcfBinary, "start", port)
	return cmd.Run()
}

// 关闭 wcf 服务
// return error 错误信息
func (c *Client) wxDestroySDK() error {
	if c.WcfBinary == "" {
		return nil
	}
	// 关闭 wcf 服务
	logman.Warn(c.WcfBinary + " stop")
	cmd := exec.Command(c.WcfBinary, "stop")
	return cmd.Run()
}
