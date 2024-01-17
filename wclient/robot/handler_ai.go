package robot

import (
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wclient/model"
)

func aiHandler() {

	handlers["/ai"] = &Handler{
		Level:    0,
		ChatAble: true,
		RoomAble: true,
		Describe: "提问或交谈",
		Callback: func(msg *wcferry.WxMsg) string {
			if msg.Content == "" {
				return "请在指令后面输入问题"
			}
			return model.AiChat(msg.Sender, msg.Content)
		},
	}

}
