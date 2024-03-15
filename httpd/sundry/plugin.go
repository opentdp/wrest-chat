package sundry

import (
	"github.com/gin-gonic/gin"

	"github.com/opentdp/wechat-rest/wclient/plugin"
)

// @Summary 计划任务插件列表
// @Tags DRY::插件
// @Produce json
// @Success 200 {array} plugin.CronjobPlugin
// @Router /api/plugin/cronjobs [post]
func pluginCronjobs(c *gin.Context) {

	plugins, err := plugin.CronjobPluginSetup()

	if err != nil {
		c.Set("Error", err)
		return
	}

	c.Set("Payload", plugins)

}

// @Summary 外部指令插件列表
// @Tags DRY::插件
// @Produce json
// @Success 200 {array} plugin.KeywordPlugin
// @Router /api/plugin/keywords [post]
func pluginKeywords(c *gin.Context) {

	plugins, err := plugin.KeywordPluginSetup()

	if err != nil {
		c.Set("Error", err)
		return
	}

	c.Set("Payload", plugins)

}
