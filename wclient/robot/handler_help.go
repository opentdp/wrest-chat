package robot

import (
	"fmt"
	"sort"
	"strings"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wclient/model"
)

func helpHandler() {

	helper1 := []string{} // 私聊指令
	helper2 := []string{} // 群聊指令

	handlers["/help"] = &Handler{
		Level:    0,
		ChatAble: true,
		RoomAble: true,
		Describe: "查看帮助信息",
		Callback: func(msg *wcferry.WxMsg) string {
			text := ""
			// 场景帮助
			if msg.IsGroup {
				text += strings.Join(helper2, "\n") + "\n"
			} else {
				text += strings.Join(helper1, "\n") + "\n"
			}
			// 模型运行时信息
			if model.GetUserConfig(msg.Sender).WakeWord != "" {
				text += "唤醒词 " + model.GetUserConfig(msg.Sender).WakeWord + "，"
			}
			if len(args.LLM.Models) > 0 {
				text += "对话模型 " + model.GetUserConfig(msg.Sender).LLModel.Name + "，"
				text += fmt.Sprintf("上下文长度 %d/%d", model.CountHistory(msg.Sender), args.LLM.HistoryNum)
			}
			return text
		},
	}

	// 生成指令菜单
	for k, v := range handlers {
		if v.ChatAble {
			helper1 = append(helper1, k+" "+v.Describe)
		}
		if v.RoomAble {
			helper2 = append(helper2, k+" "+v.Describe)
		}
	}

	// 按字母排序
	sort.Strings(helper1)
	sort.Strings(helper2)

}
