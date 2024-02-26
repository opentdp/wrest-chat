package robot

import (
	"fmt"
	"sort"
	"strings"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/dbase/chatroom"
	"github.com/opentdp/wechat-rest/dbase/profile"
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wclient/aichat"
)

func helpHandler() {

	handlers["/help"] = &Handler{
		Level:    0,
		ChatAble: true,
		RoomAble: true,
		Describe: "查看帮助信息",
		Callback: func(msg *wcferry.WxMsg) string {
			up, _ := profile.Fetch(&profile.FetchParam{Wxid: msg.Sender, Roomid: msg.Roomid})
			// 生成指令菜单
			helper := []string{}
			for k, v := range handlers {
				if v.Level > 0 {
					if up == nil || v.Level > up.Level {
						continue // 没有权限
					}
				}
				if msg.IsGroup {
					if v.RoomAble { // 群聊指令
						helper = append(helper, k+" "+v.Describe)
					}
				} else {
					if v.ChatAble { // 私聊指令
						helper = append(helper, k+" "+v.Describe)
					}
				}
			}
			sort.Strings(helper)
			text := strings.Join(helper, "\n") + "\n"
			if up.Level > 0 {
				text += fmt.Sprintf("级别 %d；", up.Level)
			}
			if up.AiArgot != "" {
				text += fmt.Sprintf("唤醒词 %s；", up.AiArgot)
			}
			if len(args.LLM.Models) > 0 {
				text += fmt.Sprintf("对话模型 %s；", aichat.UserModel(msg.Sender, msg.Roomid).Family)
				text += fmt.Sprintf("上下文长度 %d/%d；", aichat.CountHistory(msg.Sender), args.LLM.HistoryNum)
			}
			if msg.IsGroup {
				room, err := chatroom.Fetch(&chatroom.FetchParam{Roomid: msg.Roomid})
				if err == nil && room.Level > 0 {
					text += fmt.Sprintf("群级别 %d；", room.Level)
					if up.Level > 0 {
						text += fmt.Sprintf("群成员级别 %d；", up.Level)
					}
				}
			}
			return text + "祝你好运！"
		},
	}

}
