package robot

import (
	"fmt"

	"github.com/opentdp/wrest-chat/dbase/setting"
	"github.com/opentdp/wrest-chat/dbase/webhook"
	"github.com/opentdp/wrest-chat/wcferry"
)

func webhookHandler() []*Handler {

	cmds := []*Handler{}

	if len(setting.ApiEndpoint) < 10 {
		return cmds
	}

	cmds = append(cmds, &Handler{
		Level:    7,
		Order:    210,
		Roomid:   "*",
		Command:  "/webhook",
		Describe: "创建 Webhook",
		Callback: func(msg *wcferry.WxMsg) string {
			target := msg.Sender
			if msg.IsGroup {
				target = msg.GetRoomid()
			}
			// 已存在
			item, err := webhook.Fetch(&webhook.FetchWebhookParam{
				TargetId: target,
			})
			if err == nil {
				return fmt.Sprintf("webhook 调用地址: /bot/webhook/%s/{type}", item.Token)
			}
			// 创建新的 Webhook
			token, err := webhook.Create(&webhook.CreateWebhookParam{
				TargetId: target,
				Remark:   fmt.Sprintf("由用户[%s]通过指令创建", msg.Sender),
			})
			if err != nil {
				return "webhook 创建失败，" + err.Error()
			}
			return fmt.Sprintf("webhook 已添加\nToken: %s\n调用地址: /bot/webhook/%s/{type}\ntype 为不同类型的应用发送的webhook(如github, gitea)\n自定义的请填写text直接原样发送body", token, token)
		},
	})

	cmds = append(cmds, &Handler{
		Level:    7,
		Order:    211,
		Roomid:   "*",
		Command:  "/webhook:rm",
		Describe: "删除 Webhook",
		Callback: func(msg *wcferry.WxMsg) string {
			target := msg.Sender
			if msg.IsGroup {
				target = msg.GetRoomid()
			}
			res := webhook.Delete(&webhook.DeleteWebhookParam{TargetId: target})
			if res != nil {
				return "webhook 删除失败，" + res.Error()
			}
			return "webhook 删除成功"
		},
	})

	return cmds

}
