package robot

import (
	"regexp"
	"strings"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/wcferry"
)

var models = make(map[string]string)
var history = make(map[string][]string)
var handlers = make(map[string]commandHandler)

type commandHandler func(id, msg string) string

func initHandlers() {

	helper := []string{
		"/ai 提问或交谈",
		"/new 重置上下文内容",
		"/m:gpt35 切换为 Openai GPT-3.5 模型",
		"/m:gemini 切换为 Google Gemini 模型",
	}

	handlers["/ai"] = func(id, msg string) string {
		if _, exists := history[id]; !exists {
			history[id] = []string{}
		}
		return strings.TrimSpace(strings.TrimPrefix(msg, "/ai"))
	}

	handlers["/new"] = func(id, msg string) string {
		history[id] = []string{}
		return "已清空上下文"
	}

	handlers["/m:gpt35"] = func(id, msg string) string {
		models[id] = "gpt35"
		history[id] = []string{}
		return "你的对话模型已切换为 openai gpt-3.5"
	}

	handlers["/m:gemini"] = func(id, msg string) string {
		models[id] = "gemini"
		history[id] = []string{}
		return "你的对话模型已切换为 google gemini"
	}

	handlers["/help"] = func(id, msg string) string {
		text := strings.Join(helper, "\n")
		text += "\n/help 显示此帮助信息"
		text += "\n当前对话模型 " + models[id] + "，上下文长度 " + string(rune(len(history[id])))
		return text
	}

	for _, v := range args.Bot.RoomAddList {
		handlers["/room:"+v.Id] = func(id, msg string) string {
			resp := wc.CmdClient.InviteChatroomMembers(v.Id, id)
			if resp == 1 {
				return "已发送群邀请，稍后请点击进入"
			} else {
				return "发送群邀请失败"
			}
		}
	}

}

func chatCommand(msg *wcferry.WxMsg) bool {

	re := regexp.MustCompile(`^/([\w:-]{2,20})(\s|$)`)
	matches := re.FindStringSubmatch(msg.Content)
	if matches == nil || len(matches) < 2 {
		return false
	}

	output := ""
	command := matches[1]
	content := msg.Content[len(matches[0]):]

	if len(handlers) == 0 {
		initHandlers()
	}

	if fn, ok := handlers[command]; ok {
		output = fn(msg.Sender, content)
	} else {
		output = "指令或参数错误"
	}

	if msg.IsGroup {
		user := wc.CmdClient.GetInfoByWxid(msg.Sender)
		wc.CmdClient.SendTxt(msg.Roomid, "@"+user.Name+"\n"+output, msg.Sender)
	} else {
		wc.CmdClient.SendTxt(msg.Sender, output, "")
	}

	return true

}
