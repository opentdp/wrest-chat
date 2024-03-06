package robot

import (
	"fmt"
	"strings"

	"github.com/opentdp/wechat-rest/dbase/chatroom"
	"github.com/opentdp/wechat-rest/dbase/llmodel"
	"github.com/opentdp/wechat-rest/dbase/profile"
	"github.com/opentdp/wechat-rest/dbase/setting"
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

	up, _ := profile.Fetch(&profile.FetchParam{Wxid: msg.Sender, Roomid: prid(msg)})

	// 生成指令菜单
	helper := []string{}
	for _, k := range orderHandlers() {
		v := handlers[k]
		if v.Level > 0 {
			if up == nil || v.Level > up.Level {
				continue // 没有权限
			}
		}
		if (msg.IsGroup && v.RoomAble) || (!msg.IsGroup && v.ChatAble) {
			o := fmt.Sprintf("%s %s", k, v.Describe)
			helper = append(helper, o)
		}
	}

	// 数组转为字符串
	text := strings.Join(helper, "\n") + "\n"

	// 自定义帮助信息
	if len(setting.HelpAdditive) > 1 {
		text += setting.HelpAdditive + "\n"
	}

	// 当前用户状态信息
	if up.Level > 0 {
		text += fmt.Sprintf("级别 %d；", up.Level)
	}

	// 对话模型相关配置
	llmCount, _ := llmodel.Count(&llmodel.CountParam{})
	if llmCount > 0 {
		if strings.Trim(up.AiArgot, "-") != "" {
			text += fmt.Sprintf("唤醒词 %s；", up.AiArgot)
		}
		text += fmt.Sprintf("对话模型 %s；", aichat.UserModel(msg.Sender, msg.Roomid).Family)
		text += fmt.Sprintf("上下文长度 %d/%d；", aichat.CountHistory(msg.Sender), setting.ModelHistory)
	}

	return text + "祝你好运！"

}

func helpPreCheck(msg *wcferry.WxMsg) string {

	if setting.WhiteLimit {
		// 管理豁免
		up, _ := profile.Fetch(&profile.FetchParam{Wxid: msg.Sender, Roomid: prid(msg)})
		if up.Level >= 7 {
			return ""
		}
		// 权限检查
		if msg.IsGroup {
			room, _ := chatroom.Fetch(&chatroom.FetchParam{Roomid: msg.Roomid})
			if room.Level < 2 {
				return "-"
			}
		} else if up.Level < 2 {
			return "-"
		}
	}

	return ""

}
