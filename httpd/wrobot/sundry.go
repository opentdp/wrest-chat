package wrobot

import (
	"github.com/gin-gonic/gin"
	"github.com/opentdp/wechat-rest/wclient/robot"
)

// @Summary 机器人指令集
// @Tags BOT::杂项
// @Success 200 {object} []robot.Handler
// @Router /bot/handlers [post]
func handlers(c *gin.Context) {

	c.Set("Payload", robot.Handlers)

}
