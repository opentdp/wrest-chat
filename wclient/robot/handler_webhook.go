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
		Level:    -1,
		Order:    200,
		Roomid:   "*",
		Command:  "/webhook:add",
		Describe: "创建Webhook",
		Callback: func(msg *wcferry.WxMsg) string {
			target := msg.Sender
			if msg.IsGroup {
				target = msg.GetRoomid()
			}
			_, token, err := webhook.Create(&webhook.CreateWebhookParam{
				TargetId: target,
				Remark:   fmt.Sprintf("由用户[%s]通过指令创建", msg.Sender),
			})
			if err != nil {
				return "创建失败, 已经存在webhook，不可重复创建."
			}
			return fmt.Sprintf("webhook已添加\nToken: %s\n调用地址: /bot/webhook/%s/{type}\ntype 为不同类型的应用发送的webhook(如github, gitea)\n自定义的请填写text直接原样发送body", token, token)
		},
	})

	cmds = append(cmds, &Handler{
		Level:    -1,
		Order:    200,
		Roomid:   "*",
		Command:  "/webhook:rm",
		Describe: "删除Webhook",
		Callback: func(msg *wcferry.WxMsg) string {
			target := msg.Sender
			if msg.IsGroup {
				target = msg.GetRoomid()
			}
			if !webhook.DeleteByTargetId(target) {
				return "删除失败"
			}
			return "删除成功"
		},
	})

	return cmds

}
