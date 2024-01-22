package robot

import (
	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/wcferry"
)

func modelHandler() {

	if len(args.LLM.Models) == 0 {
		return
	}

	for k, v := range args.LLM.Models {
		k, v := k, v // copy
		cmdkey := "/m:" + k
		handlers[cmdkey] = &Handler{
			Level:    0,
			ChatAble: true,
			RoomAble: true,
			Describe: "切换为 " + v.Model + " 模型",
			Callback: func(msg *wcferry.WxMsg) string {
				args.GetMember(msg.Sender).AiModel = k
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
			for k, v := range args.LLM.Models {
				args.GetMember(msg.Sender).AiModel = k
				return "对话模型切换为 " + v.Family + " [" + v.Model + "]"
			}
			return "没有可用的模型"
		},
	}

}
