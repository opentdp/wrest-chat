package robot

import (
	"strings"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/wcferry"
)

func wakeHandler() {

	handlers["/wake"] = &Handler{
		Level:    0,
		ChatAble: true,
		RoomAble: true,
		Describe: "设置或禁用唤醒词",
		Callback: func(msg *wcferry.WxMsg) string {
			argot := strings.TrimSpace(msg.Content)
			if strings.Contains(argot, "@") || strings.Contains(argot, "/") {
				return "唤醒词不允许包含 @ 或 /"
			}
			if argot != "" {
				args.GetMember(msg.Sender).AiArgot = argot
				return "唤醒词设置为 " + argot
			}
			return "已禁用唤醒词"
		},
	}

}
