//go:build !windows

package wcferry

import (
	"github.com/opentdp/go-helper/logman"
)

// 调用 sdk.dll 中的函数
// return error 错误信息
func (c *Client) sdkCall(fn string, a ...uintptr) error {
	logman.Warn("skip to load sdk.dll", "fn", fn, "a", a)
	return nil
}
