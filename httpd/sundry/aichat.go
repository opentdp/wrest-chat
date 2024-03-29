package sundry

import (
	"github.com/gin-gonic/gin"

	"github.com/opentdp/wrest-chat/wclient/aichat"
)

type AiChatParam struct {
	Wxid    string `json:"wxid"`
	Message string `json:"message"`
}

// @Summary 获取模型配置
// @Produce json
// @Tags API::智能聊天
// @Param body body AiChatParam true "智能聊天参数"
// @Success 200 {object} aichat.UserConfig
// @Router /api/aichat/config [post]
func aiChatConfig(c *gin.Context) {

	var rq *AiChatParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	config := aichat.UserConfig(rq.Wxid, "")
	config.Secret = "***"

	c.Set("Payload", config)

}

// @Summary 发起文本聊天
// @Produce json
// @Tags API::智能聊天
// @Param body body AiChatParam true "智能聊天参数"
// @Success 200 {string} string
// @Router /api/aichat/text [post]
func aiChatText(c *gin.Context) {

	var rq *AiChatParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	text := aichat.Text(rq.Message, rq.Wxid, "")

	c.Set("Payload", text)

}
