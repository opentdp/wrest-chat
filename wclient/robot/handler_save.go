package robot

import (
	"fmt"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/wcferry"
)

func saveHandler() {

	handlers["/save"] = &Handler{
		Level:    9,
		ChatAble: true,
		RoomAble: true,
		Describe: "保存配置信息",
		Callback: func(msg *wcferry.WxMsg) string {
			if err := args.SaveConfig(); err != nil {
				return fmt.Sprintf("写入配置错误：%s", err)
			}
			return "写入配置成功"
		},
	}

}
