package sundry

import (
	"github.com/opentdp/go-helper/httpd"

	"github.com/opentdp/wrest-chat/httpd/middle"
)

func Route() {

	rg := httpd.Group("/api")
	rg.Use(middle.OutputHandle, middle.ApiGuard)

	aichat := AiChat{}
	rg.POST("aichat/config", aichat.config)
	rg.POST("aichat/text", aichat.text)

	cronjob := Cronjob{}
	rg.POST("cronjob/list", cronjob.list)
	rg.POST("cronjob/detail", cronjob.detail)
	rg.POST("cronjob/create", cronjob.create)
	rg.POST("cronjob/update", cronjob.update)
	rg.POST("cronjob/delete", cronjob.delete)
	rg.POST("cronjob/status", cronjob.status)
	rg.POST("cronjob/execute", cronjob.execute)

	plugin := Plugin{}
	rg.POST("plugin/cronjobs", plugin.cronjobs)
	rg.POST("plugin/keywords", plugin.keywords)

	system := System{}
	rg.POST("system/version", system.version)
	rg.POST("system/handlers", system.handlers)

}
