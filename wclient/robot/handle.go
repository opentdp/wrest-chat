package robot

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wclient/cache"
	"github.com/opentdp/wechat-rest/wclient/model"
)

func chatHandler(msg *wcferry.WxMsg) bool {

	re := regexp.MustCompile(`^(/[\w:-]{2,20})(\s|$)`)
	matches := re.FindStringSubmatch(msg.Content)
	if matches == nil || len(matches) < 2 {
		return false
	}

	output := ""
	command := matches[1]
	content := msg.Content[len(matches[0]):]

	if fn, ok := cache.Handlers[command]; ok {
		output = fn(msg.Sender, content)
	} else {
		output = "指令或参数错误"
	}

	if msg.IsGroup {
		user := wc.CmdClient.GetInfoByWxid(msg.Sender)
		wc.CmdClient.SendTxt("@"+user.Name+"\n"+output, msg.Roomid, msg.Sender)
	} else {
		wc.CmdClient.SendTxt(output, msg.Sender, "")
	}

	return true

}

func initHandlers() {

	helper := []string{
		"/ai 提问或交谈",
		"/new 重置上下文内容",
		"/m:gpt35 切换为 Openai GPT-3.5 模型",
		"/m:gemini 切换为 Google Gemini 模型",
	}

	cache.Handlers["/ai"] = func(id, msg string) string {
		var err error
		var res string
		if _, exists := cache.Models[id]; !exists {
			cache.Models[id] = "gemini-pro"
		}
		if _, exists := cache.History[id]; !exists {
			cache.History[id] = []cache.HistoryItem{}
		}
		str := strings.TrimSpace(strings.TrimPrefix(msg, "/ai"))
		switch cache.Models[id] {
		case "gpt-3.5-turbo":
			res, err = model.OpenaiChat(id, str)
		case "gemini-pro":
			res, err = model.GeminiChat(id, str)
		}
		if err != nil {
			return err.Error()
		}
		return res
	}

	cache.Handlers["/new"] = func(id, msg string) string {
		cache.History[id] = []cache.HistoryItem{}
		return "已清空上下文"
	}

	cache.Handlers["/m:gpt35"] = func(id, msg string) string {
		cache.Models[id] = "gpt-3.5-turbo"
		cache.History[id] = []cache.HistoryItem{}
		return "对话模型已切换为 Openai GPT-3.5"
	}

	cache.Handlers["/m:gemini"] = func(id, msg string) string {
		cache.Models[id] = "gemini-pro"
		cache.History[id] = []cache.HistoryItem{}
		return "对话模型已切换为 Google Gemini"
	}

	cache.Handlers["/help"] = func(id, msg string) string {
		text := strings.Join(helper, "\n")
		text += "\n/help 显示此帮助信息"
		text += "\n当前对话模型 " + cache.Models[id] + "，上下文长度 " + strconv.Itoa(len(cache.History[id]))
		return text
	}

	for k, v := range args.Bot.InvitableRooms {
		cmdkey := "/room:" + strconv.Itoa(k+1)
		helper = append(helper, cmdkey+" 加入群聊 "+v.Name)
		cache.Handlers[cmdkey] = func(id, msg string) string {
			resp := wc.CmdClient.InviteChatroomMembers(v.Id, id)
			if resp == 1 {
				return "已发送群邀请，稍后请点击进入"
			} else {
				return "发送群邀请失败"
			}
		}
	}

}
