package wrobot

import (
	"github.com/opentdp/go-helper/httpd"

	"github.com/opentdp/wechat-rest/httpd/midware"
)

func Route() {

	rg := httpd.Group("/bot")
	rg.Use(midware.OutputHandle, midware.ApiGuard)

	chatroom := Chatroom{}
	rg.POST("chatroom/list", chatroom.list)
	rg.POST("chatroom/create", chatroom.create)
	rg.POST("chatroom/detail", chatroom.detail)
	rg.POST("chatroom/update", chatroom.update)
	rg.POST("chatroom/delete", chatroom.delete)

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

}
