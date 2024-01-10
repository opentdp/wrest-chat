package robot

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wclient/model"
)

var handlers = make(map[string]func(id, msg string) string)

func initHandler() {

	helper := []string{}

	helper = append(helper, "/ai 提问或交谈")
	handlers["/ai"] = model.AiChat

	helper = append(helper, "/new 重置上下文内容")
	handlers["/new"] = model.Clear

	for k, v := range args.LLM.Models {
		k, v := k, v // copy it
		cmdkey := "/m:" + v.Name
		helper = append(helper, cmdkey+" 切换为 "+v.Model+" 模型")
		handlers[cmdkey] = func(id, msg string) string {
			model.Clear(id, "")
			model.Models[id] = k
			return "对话模型已切换为 " + v.Name + " [" + v.Model + "]"
		}
	}

	for k, v := range args.Bot.InvitableRooms {
		k, v := k, v // copy it
		cmdkey := "/room:" + strconv.Itoa(k+1)
		helper = append(helper, cmdkey+" 加入群聊 "+v.Name)
		handlers[cmdkey] = func(id, msg string) string {
			resp := wc.CmdClient.InviteChatroomMembers(v.Id, id)
			if resp == 1 {
				return "已发送群邀请，稍后请点击进入"
			} else {
				return "发送群邀请失败"
			}
		}
	}

	helper = append(helper, "/help 查看帮助信息")
	handlers["/help"] = func(id, msg string) string {
		text := strings.Join(helper, "\n") + "\n"
		text += "当前对话模型 " + model.Model(id).Name + "，"
		text += "上下文长度 " + strconv.Itoa(len(model.History[id])) + "/" + strconv.Itoa(args.LLM.HistoryNum)
		return text
	}

}

func chatHandler(msg *wcferry.WxMsg) bool {

	re := regexp.MustCompile(`^(/[\w:-]{2,20})(\s|$)`)
	matches := re.FindStringSubmatch(msg.Content)
	if matches == nil || len(matches) < 2 {
		return false
	}

	output := ""
	command := matches[1]
	content := strings.TrimSpace(msg.Content[len(matches[0]):])

	if command == "/ai" && content == "" {
		command = "/help"
	}

	// 执行指令
	if fn, ok := handlers[command]; ok {
		output = fn(msg.Sender, content)
	} else {
		output = "指令或参数错误"
	}

	// 发送消息
	if msg.IsGroup {
		user := wc.CmdClient.GetInfoByWxid(msg.Sender)
		wc.CmdClient.SendTxt("@"+user.Name+"\n"+output, msg.Roomid, msg.Sender)
	} else {
		wc.CmdClient.SendTxt(output, msg.Sender, "")
	}

	return true

}
