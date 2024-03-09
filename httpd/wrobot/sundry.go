package wrobot

import (
	"github.com/gin-gonic/gin"
	"github.com/opentdp/wechat-rest/wclient/robot"
)

// @Summary 机器人指令集
// @Tags BOT::杂项
// @Success 200 {array} RobotHandler
// @Router /bot/handlers [post]
func robotHandlers(c *gin.Context) {

	items := []RobotHandler{}

	for _, v := range robot.Handlers {
		items = append(items, RobotHandler{
			Level:    v.Level,
			Order:    v.Order,
			ChatAble: v.ChatAble,
			RoomAble: v.RoomAble,
			Command:  v.Command,
			Describe: v.Describe,
		})
	}

	c.Set("Payload", items)

}

type RobotHandler struct {
	Level    int32  `json:"level"`     // 0:不限制 7:群管理 9:创始人
	Order    int32  `json:"order"`     // 排序，越小越靠前
	ChatAble bool   `json:"chat_able"` // 是否允许在私聊使用
	RoomAble bool   `json:"room_able"` // 是否允许在群聊使用
	Command  string `json:"command"`   // 指令
	Describe string `json:"describe"`  // 指令的描述信息
}
