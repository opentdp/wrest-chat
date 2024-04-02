package sundry

import (
	"github.com/gin-gonic/gin"

	"github.com/opentdp/wrest-chat/wclient/plugin"
)

type Plugin struct{}

// @Summary 计划任务插件列表
// @Tags API::插件管理
// @Produce json
// @Success 200 {array} plugin.CronjobPlugin
// @Router /api/plugin/cronjobs [post]
func (*Plugin) cronjobs(c *gin.Context) {

	plugins := plugin.CronjobPluginSetup()

	c.Set("Payload", plugins)

}

// @Summary 外部指令插件列表
// @Tags API::插件管理
// @Produce json
// @Success 200 {array} plugin.KeywordPlugin
// @Router /api/plugin/keywords [post]
func (*Plugin) keywords(c *gin.Context) {

	plugins := plugin.KeywordPluginSetup()

	c.Set("Payload", plugins)

}
