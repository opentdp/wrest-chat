package wrobot

import (
	"github.com/opentdp/go-helper/httpd"

	"github.com/opentdp/wrest-chat/httpd/middle"
)

func Route() {

	rg := httpd.Group("/bot")
	rg.Use(middle.OutputHandle, middle.ApiGuard)

	chatroom := Chatroom{}
	rg.POST("chatroom/list", chatroom.list)
	rg.POST("chatroom/create", chatroom.create)
	rg.POST("chatroom/detail", chatroom.detail)
	rg.POST("chatroom/update", chatroom.update)
	rg.POST("chatroom/delete", chatroom.delete)

	cronjob := Cronjob{}
	rg.POST("cronjob/list", cronjob.list)
	rg.POST("cronjob/detail", cronjob.detail)
	rg.POST("cronjob/create", cronjob.create)
	rg.POST("cronjob/update", cronjob.update)
	rg.POST("cronjob/delete", cronjob.delete)
	rg.POST("cronjob/status", cronjob.status)
	rg.POST("cronjob/execute", cronjob.execute)

	keyword := Keyword{}
	rg.POST("keyword/list", keyword.list)
	rg.POST("keyword/create", keyword.create)
	rg.POST("keyword/detail", keyword.detail)
	rg.POST("keyword/update", keyword.update)
	rg.POST("keyword/delete", keyword.delete)

	llmodel := LLModel{}
	rg.POST("llmodel/list", llmodel.list)
	rg.POST("llmodel/create", llmodel.create)
	rg.POST("llmodel/detail", llmodel.detail)
	rg.POST("llmodel/update", llmodel.update)
	rg.POST("llmodel/delete", llmodel.delete)

	profile := Profile{}
	rg.POST("profile/list", profile.list)
	rg.POST("profile/create", profile.create)
	rg.POST("profile/detail", profile.detail)
	rg.POST("profile/update", profile.update)
	rg.POST("profile/delete", profile.delete)

	setting := Setting{}
	rg.POST("setting/list", setting.list)
	rg.POST("setting/detail", setting.detail)
	rg.POST("setting/create", setting.create)
	rg.POST("setting/update", setting.update)
	rg.POST("setting/delete", setting.delete)

	webhook := Webhook{}
	rg.POST("webhook/list", webhook.list)
	rg.POST("webhook/detail", webhook.detail)
	rg.POST("webhook/create", webhook.create)
	rg.POST("webhook/delete", webhook.delete)
	rg.POST("webhook/:token/:app", webhook.receive)

}
