package robot

import (
	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wclient/model"
)

func modelHandler() {

	if len(args.LLM.Models) == 0 {
		return
	}

	for k, v := range args.LLM.Models {
		v := v // copy it
		cmdkey := "/m:" + k
		handlers[cmdkey] = &Handler{
			Level:    0,
			ChatAble: true,
			RoomAble: true,
			Describe: "切换为 " + v.Model + " 模型",
			Callback: func(msg *wcferry.WxMsg) string {
				model.GetUserConfig(msg.Sender).LLModel = v
				return "对话模型切换为 " + v.Family + " [" + v.Model + "]"
			},
		}
	}

	handlers["/mr"] = &Handler{
		Level:    0,
		ChatAble: true,
		RoomAble: true,
		Describe: "随机选择模型",
		Callback: func(msg *wcferry.WxMsg) string {
			for _, v := range args.LLM.Models {
				model.GetUserConfig(msg.Sender).LLModel = v
				return "对话模型切换为 " + v.Family + " [" + v.Model + "]"
			}
			return "没有可用的模型"
		},
	}

}
