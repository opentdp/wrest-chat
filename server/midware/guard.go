package midware

import (
	"strings"

	"github.com/opentdp/wechat-rest/config"

	"github.com/gin-gonic/gin"
)

func AuthGuard(c *gin.Context) {

	token := ""

	authcode := c.GetHeader("Authorization")
	parts := strings.SplitN(authcode, " ", 2)
	if len(parts) == 2 && parts[0] == "Bearer" {
		token = parts[1]
	}

	if token != config.Httpd.Token {
		c.Set("Error", gin.H{"Code": 401, "Message": "未授权的操作"})
		c.Set("ExitCode", 401)
		c.Abort()
	}

}
