package sundry

import (
	"github.com/gin-gonic/gin"

	"github.com/opentdp/wechat-rest/wclient/plugin"
)

// @Summary 计划任务插件列表
// @Tags API::插件管理
// @Produce json
// @Success 200 {array} plugin.CronjobPlugin
// @Router /api/plugin/cronjobs [post]
func pluginCronjobs(c *gin.Context) {

	plugins := plugin.CronjobPluginSetup()

	c.Set("Payload", plugins)

}

// @Summary 外部指令插件列表
// @Tags API::插件管理
// @Produce json
// @Success 200 {array} plugin.KeywordPlugin
// @Router /api/plugin/keywords [post]
func pluginKeywords(c *gin.Context) {

	plugins := plugin.KeywordPluginSetup()

	c.Set("Payload", plugins)

}
