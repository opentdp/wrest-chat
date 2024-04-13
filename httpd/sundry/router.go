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

	plugin := Plugin{}
	rg.POST("plugin/cronjobs", plugin.cronjobs)
	rg.POST("plugin/keywords", plugin.keywords)

	system := System{}
	rg.POST("system/version", system.version)
	rg.POST("system/handlers", system.handlers)

}
