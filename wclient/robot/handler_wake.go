package robot

import (
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
			args.GetMember(msg.Sender).AiArgot = msg.Content
			if msg.Content != "" {
				return "唤醒词设置为 " + msg.Content
			}
			return "已禁用唤醒词"
		},
	}

}
