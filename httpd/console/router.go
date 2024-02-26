package console

import (
	"github.com/opentdp/go-helper/httpd"

	"github.com/opentdp/wechat-rest/httpd/midware"
)

func Route() {

	rg := httpd.Group("/cpi")
	rg.Use(midware.OutputHandle, midware.ApiGuard)

	chatroom := Chatroom{}
	rg.POST("chatroom/list", chatroom.list)
	rg.POST("chatroom/create", chatroom.create)
	rg.POST("chatroom/detail", chatroom.detail)
	rg.POST("chatroom/update", chatroom.update)
	rg.POST("chatroom/delete", chatroom.delete)

	profile := Profile{}
	rg.POST("profile/list", profile.list)
	rg.POST("profile/create", profile.create)
	rg.POST("profile/detail", profile.detail)
	rg.POST("profile/update", profile.update)
	rg.POST("profile/delete", profile.delete)

	keyword := Keyword{}
	rg.POST("keyword/list", keyword.list)
	rg.POST("keyword/create", keyword.create)
	rg.POST("keyword/detail", keyword.detail)
	rg.POST("keyword/update", keyword.update)
	rg.POST("keyword/delete", keyword.delete)

}
