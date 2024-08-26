//go:build windows

package wcferry

import (
	"errors"
	"syscall"

	"github.com/opentdp/go-helper/filer"
	"github.com/opentdp/go-helper/logman"
)

// 调用 sdk.dll 中的函数
// return error 错误信息
func (c *Client) sdkCall(fn string, a ...uintptr) error {
	if c.SdkLibrary == "" {
		logman.Warn("skip to load sdk.dll")
		return nil
	}
	// 查找 sdk.dll
	dll := c.SdkLibrary
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
