package robot

import (
	"fmt"
	"strings"

	"github.com/opentdp/wrest-chat/dbase/keyword"
	"github.com/opentdp/wrest-chat/dbase/llmodel"
	"github.com/opentdp/wrest-chat/dbase/profile"
	"github.com/opentdp/wrest-chat/dbase/setting"
	"github.com/opentdp/wrest-chat/wcferry"
	"github.com/opentdp/wrest-chat/wclient/aichat"
)

func helpHandler() []*Handler {

	cmds := []*Handler{}

	cmds = append(cmds, &Handler{
		Level:    -1,
		Order:    900,
		Roomid:   "*",
		Command:  "/help",
		Describe: "查看帮助信息",
		Callback: helpCallback,
	})

	return cmds

}

func helpCallback(msg *wcferry.WxMsg) string {

	up, _ := profile.Fetch(&profile.FetchParam{Wxid: msg.Sender, Roomid: prid(msg)})

	// 别名映射表
	aliasMap := map[string]map[string]string{}
	keywords, err := keyword.FetchAll(&keyword.FetchAllParam{Group: "handler"})
	if err == nil {
		for _, v := range keywords {
			if aliasMap[v.Roomid] == nil {
				aliasMap[v.Roomid] = map[string]string{}
			}
			aliasMap[v.Roomid][v.Target] = v.Phrase
		}
	}

	// 生成指令菜单
	helper := []string{}
	for _, v := range handlers {
		cmd := v.Command
		if v.Level > 0 {
			if up == nil || v.Level > up.Level {
				continue // 没有权限
			}
		}
		if msg.IsGroup {
			if v.Roomid != "*" && v.Roomid != "+" && v.Roomid != msg.Roomid {
				continue // 没有权限
			}
			if aliasMap[msg.Roomid] != nil && aliasMap[msg.Roomid][v.Command] != "" {
				cmd = aliasMap[msg.Roomid][v.Command]
			} else if aliasMap["+"] != nil && aliasMap["+"][v.Command] != "" {
				cmd = aliasMap["+"][v.Command]
			} else if aliasMap["*"] != nil && aliasMap["*"][v.Command] != "" {
				cmd = aliasMap["*"][v.Command]
			}
		} else {
			if v.Roomid != "*" && v.Roomid != "-" {
				continue // 没有权限
			}
			if aliasMap["-"] != nil && aliasMap["-"][v.Command] != "" {
				cmd = aliasMap["-"][v.Command]
			} else if aliasMap["*"] != nil && aliasMap["*"][v.Command] != "" {
				cmd = aliasMap["*"][v.Command]
			}
		}
		helper = append(helper, fmt.Sprintf("【%s】%s", cmd, v.Describe))
	}

	// 数组转为字符串
	text := strings.Join(helper, "\n") + "\n"

	// 自定义帮助信息
	if len(setting.HelpAdditive) > 1 {
		text += setting.HelpAdditive + "\n"
	}

	// 分割线
	text += "----------------\n"

	// 当前用户状态信息
	if up.Level > 0 {
		text += fmt.Sprintf("级别 %d；", up.Level)
	}

	// 对话模型相关配置
	llmCount, _ := llmodel.Count(&llmodel.CountParam{})
	if llmCount > 0 {
		uc := aichat.UserConfig(msg.Sender, msg.Roomid)
		if len(uc.Family) > 1 {
			text += fmt.Sprintf("对话模型 %s；", uc.Family)
		}
		text += fmt.Sprintf("上下文长度 %d/%d；", len(uc.MsgHistorys), uc.MsgHistoryMax)
	}

	return text + "祝你好运！"

}
