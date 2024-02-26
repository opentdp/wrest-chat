package robot

import (
	"github.com/opentdp/wechat-rest/wcferry"
)

func wgetHandler() {

	handlers["/wget"] = &Handler{
		Level:    7,
		Order:    90,
		ChatAble: true,
		RoomAble: true,
		Describe: "获取图片或文件",
		Callback: func(msg *wcferry.WxMsg) string {
			if fileReply(msg, msg.Content) == 0 {
				return ""
			}
			return "文件下载失败"
		},
	}

}
