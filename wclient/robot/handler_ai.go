package robot

import (
	"strings"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wclient/model"
)

func aiHandler() {

	handlers["/ai"] = &Handler{
		Level:    0,
		ChatAble: true,
		RoomAble: true,
		Describe: "提问或交谈",
		Callback: func(msg *wcferry.WxMsg) string {
			if msg.Content == "" {
				return "请在指令后面输入问题"
			}
			return model.AiChat(msg.Sender, msg.Content)
		},
	}

}

func aiMessagePrefix(msg *wcferry.WxMsg) string {

	if len(msg.Content) == 0 {
		return ""
	}

	if msg.Content[0:1] != "/" {
		if strings.Contains(msg.Content, "@"+self().Name) {
			msg.Content = "/ai " + msg.Content
		} else {
			wakeWord := args.GetMember(msg.Sender).AiArgot
			if wakeWord == "" {
				if !msg.IsGroup {
					msg.Content = "/ai " + msg.Content
				}
			} else if strings.HasPrefix(msg.Content, wakeWord) {
				msg.Content = strings.Replace(msg.Content, wakeWord, "/ai ", 1)
			}
		}
	}

	return ""

}
