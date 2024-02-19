package wcferry

import (
	"errors"
	"os/exec"
	"strings"
	"time"

	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/onquit"
	"golang.org/x/sys/windows"
)

type Client struct {
	ListenAddr string     // wcf 监听地址
	ListenPort int        // wcf 监听端口
	SdkLibrary string     // sdk.dll 路径
	WeChatAuto bool       // 微信自动启停
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
		c.ListenPort = 10080
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

// 调用 sdk.dll 中的函数
func (c *Client) sdkCall(fn string, a ...uintptr) error {
	// 加载 sdk.dll
	sdk, err := windows.LoadDLL(c.SdkLibrary)
	if err != nil {
		logman.Info("failed to load sdk.dll", "error", err)
		return err
	}
	defer sdk.Release()
	// 查找 fn 函数
	proc, err := sdk.FindProc(fn)
	if err != nil {
		logman.Info("failed to call "+fn, "error", err)
		return err
	}
	// 执行 fn(a...)
	r1, r2, err := proc.Call(a...)
	logman.Warn("call dll:"+fn, "r1", r1, "r2", r2, "error", err)
	return err
}

// 启动 wcf 服务
// return error 错误信息
func (c *Client) wxInitSDK() error {
	if c.SdkLibrary == "" {
		return nil
	}
	// 尝试自动启动微信
	if c.WeChatAuto {
		out, _ := exec.Command("tasklist").Output()
		if strings.Contains(string(out), "WeChat.exe") {
			return errors.New("please close wechat")
		}
	}
	// 注入微信，打开 wcf 服务
	c.sdkCall("WxInitSDK", uintptr(0), uintptr(c.ListenPort))
	time.Sleep(5 * time.Second)
	return nil
}

// 关闭 wcf 服务
// return error 错误信息
func (c *Client) wxDestroySDK() error {
	if c.SdkLibrary == "" {
		return nil
	}
	// 关闭 wcf 服务
	err := c.sdkCall("WxDestroySDK", uintptr(0))
	logman.Warn("wcf destroyed", "error", err)
	// 尝试自动关闭微信
	if c.WeChatAuto {
		logman.Info("killing wechat process")
		cmd := exec.Command("taskkill", "/IM", "WeChat.exe", "/F")
		if err := cmd.Run(); err != nil {
			logman.Warn("failed to kill wechat", "error", err)
			return err
		}
	}
	return err
}
