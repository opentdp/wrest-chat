package robot

import (
	"math/rand"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wclient/model"
)

func modelHandler() {

	if len(args.LLM.Models) == 0 {
		return
	}

	for _, v := range args.LLM.Models {
		v := v // copy it
		cmdkey := "/m:" + v.Name
		handlers[cmdkey] = &Handler{
			Level:    0,
			ChatAble: true,
			RoomAble: true,
			Describe: "切换为 " + v.Model + " 模型",
			Callback: func(msg *wcferry.WxMsg) string {
				model.GetUserConfig(msg.Sender).LLModel = v
				return "对话模型切换为 " + v.Name + " [" + v.Model + "]"
			},
		}
	}

	handlers["/mr"] = &Handler{
		Level:    0,
		ChatAble: true,
		RoomAble: true,
		Describe: "随机选择模型",
		Callback: func(msg *wcferry.WxMsg) string {
			l := len(args.LLM.Models)
			if l == 0 {
				return "没有可用模型"
			}
			v := args.LLM.Models[rand.Intn(l)]
			model.GetUserConfig(msg.Sender).LLModel = v
			return "对话模型切换为 " + v.Name + " [" + v.Model + "]"
		},
	}

}
