package robot

import (
	"strings"

	"github.com/opentdp/go-helper/request"
	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/wcferry"
)

func apiHandler() {

	handlers["/api"] = &Handler{
		Level:    0,
		ChatAble: true,
		RoomAble: true,
		Describe: "调用接口，查看帮助 /api help",
		Callback: func(msg *wcferry.WxMsg) string {
			str := strings.TrimSpace(msg.Content)
			res, err := request.TextGet("https://api.rehi.org/"+str, request.H{
				"User-Agent": args.AppName + "/" + args.Version,
			})
			if err != nil {
				return err.Error()
			}
			return res
		},
	}

}
