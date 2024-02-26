package robot

import (
	"strings"

	"github.com/opentdp/wechat-rest/dbase/profile"
	"github.com/opentdp/wechat-rest/wcferry"
)

func wakeHandler() {

	handlers["/wake"] = &Handler{
		Level:    0,
		ChatAble: true,
		RoomAble: true,
		Describe: "设置唤醒词",
		Callback: func(msg *wcferry.WxMsg) string {
			argot := strings.TrimSpace(msg.Content)
			if argot == "" {
				return "唤醒词不允许为空"
			}
			if strings.Contains(argot, "@") || strings.Contains(argot, "/") {
				return "唤醒词不允许包含 @ 或 /"
			}
			profile.Migrate(&profile.MigrateParam{Wxid: msg.Sender, Roomid: msg.Roomid, AiArgot: argot})
			return "唤醒词设置为 " + argot
		},
	}

}
