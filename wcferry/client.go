package wcferry

import (
	"errors"
	"runtime"
	"syscall"
	"time"

	"github.com/opentdp/go-helper/filer"
	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/onquit"
)

const Wcf_Version = "39.2.4.0"
const Wechat_Version = "3.9.10.27"

type Client struct {
	SdkLibrary string     // sdk.dll 路径
	ListenAddr string     // wcf 监听地址
	ListenPort int        // wcf 监听端口
	CmdClient  *CmdClient // 命令客户端
	MsgClient  *MsgClient // 消息客户端
}

// 注册消息服务
// return error 错误信息
func (c *Client) Connect() error {
	if c.ListenAddr == "" {
		c.ListenAddr = "127.0.0.1"
	}
	if c.ListenPort == 0 {
		c.ListenPort = 10086
	}
	// 启动 rpc
	if err := c.wxInitSDK(); err != nil {
		return err
	}
	// 配置客户端
	c.CmdClient = &CmdClient{
		pbSocket: newPbSocket(c.ListenAddr, c.ListenPort),
	}
	c.MsgClient = &MsgClient{
		pbSocket: newPbSocket(c.ListenAddr, c.ListenPort+1),
	}
	// 退出时注销
	onquit.Register(func() {
		c.MsgClient.Destroy()
		c.CmdClient.Destroy()
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
	err := c.sdkCall("WxInitSDK", uintptr(0), uintptr(c.ListenPort))
	time.Sleep(5 * time.Second)
	return err
}

// 关闭 wcf 服务
// return error 错误信息
func (c *Client) wxDestroySDK() error {
	return c.sdkCall("WxDestroySDK")
}

// 调用 sdk.dll 中的函数
// return error 错误信息
func (c *Client) sdkCall(fn string, a ...uintptr) error {
	dll := c.SdkLibrary
	if dll == "" || runtime.GOOS != "windows" {
		logman.Warn("skip to load sdk.dll")
		return nil
	}
	// 查找 sdk.dll
	if !filer.Exists(dll) {
		dll = "wcferry/" + dll
		if !filer.Exists(dll) {
			return errors.New(dll + " not found")
		}
	}
	// 加载 sdk.dll
	sdk, err := syscall.LoadDLL(dll)
	if err != nil {
		logman.Warn("failed to load sdk.dll", "error", err)
		return err
	}
	defer sdk.Release()
	// 查找 fn 函数
	proc, err := sdk.FindProc(fn)
	if err != nil {
		logman.Warn("failed to call "+fn, "error", err)
		return err
	}
	// 执行 fn(a...)
	r1, r2, err := proc.Call(a...)
	logman.Warn("call dll:"+fn, "r1", r1, "r2", r2, "error", err)
	if err.Error() == "Attempt to access invalid address." {
		err = nil // 忽略已知问题
	}
	return err
}
