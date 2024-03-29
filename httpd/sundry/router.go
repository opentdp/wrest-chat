package sundry

import (
	"github.com/opentdp/go-helper/httpd"

	"github.com/opentdp/wrest-chat/httpd/middle"
)

func Route() {

	rg := httpd.Group("/api")
	rg.Use(middle.OutputHandle, middle.ApiGuard)

	rg.POST("aichat/config", aiChatConfig)
	rg.POST("aichat/text", aiChatText)

	cronjob := Cronjob{}
	rg.POST("cronjob/list", cronjob.list)
	rg.POST("cronjob/detail", cronjob.detail)
	rg.POST("cronjob/create", cronjob.create)
	rg.POST("cronjob/update", cronjob.update)
	rg.POST("cronjob/delete", cronjob.delete)
	rg.POST("cronjob/status", cronjob.status)
	rg.POST("cronjob/execute", cronjob.execute)

	rg.POST("plugin/cronjobs", pluginCronjobs)
	rg.POST("plugin/keywords", pluginKeywords)

	system := System{}
	rg.POST("system/version", system.version)
	rg.POST("system/handlers", system.handlers)

}
