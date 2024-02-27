package robot

import (
	"fmt"
	"sort"
	"strings"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/dbase/chatroom"
	"github.com/opentdp/wechat-rest/dbase/llmodel"
	"github.com/opentdp/wechat-rest/dbase/profile"
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wclient/aichat"
)

func helpHandler() {

	handlers["/help"] = &Handler{
		Level:    0,
		Order:    900,
		ChatAble: true,
		RoomAble: true,
		Describe: "查看帮助信息",
		Callback: helpCallback,
		PreCheck: helpPreCheck,
	}

}

func helpCallback(msg *wcferry.WxMsg) string {

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

	// 排序后转为字符串
	sort.Strings(helper)
	text := strings.Join(helper, "\n") + "\n"
	if up.Level > 0 {
		text += fmt.Sprintf("级别 %d；", up.Level)
	}

	// 对话模型相关配置
	llmCount, _ := llmodel.Count(&llmodel.CountParam{})
	if llmCount > 0 {
		if up.AiArgot != "" && up.AiArgot != "-" {
			text += fmt.Sprintf("唤醒词 %s；", up.AiArgot)
		}
		text += fmt.Sprintf("对话模型 %s；", aichat.UserModel(msg.Sender, msg.Roomid).Family)
		text += fmt.Sprintf("上下文长度 %d/%d；", aichat.CountHistory(msg.Sender), args.LLM.HistoryNum)
	}

	return text + "祝你好运！"

}

func helpPreCheck(msg *wcferry.WxMsg) string {

	if args.Bot.WhiteLimit {
		if msg.IsGroup {
			room, _ := chatroom.Fetch(&chatroom.FetchParam{Roomid: msg.Roomid})
			if room.Level < 2 {
				return "-"
			}
		} else {
			up, _ := profile.Fetch(&profile.FetchParam{Wxid: msg.Sender})
			if up.Level < 2 {
				return "-"
			}
		}
	}

	return ""

}
