package wrobot

import (
	"github.com/gin-gonic/gin"

	"github.com/opentdp/wrest-chat/dbase/chatroom"
	"github.com/opentdp/wrest-chat/wclient/robot"
)

type Chatroom struct{}

// @Summary 群聊配置列表
// @Produce json
// @Tags BOT::群聊配置
// @Param body body chatroom.FetchAllParam true "获取群聊配置列表参数"
// @Success 200 {array} tables.Chatroom
// @Router /bot/chatroom/list [post]
func (*Chatroom) list(c *gin.Context) {

	var rq *chatroom.FetchAllParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if lst, err := chatroom.FetchAll(rq); err == nil {
		c.Set("Payload", lst)
	} else {
		c.Set("Error", err)
	}

}

// @Summary 获取群聊配置
// @Produce json
// @Tags BOT::群聊配置
// @Param body body chatroom.FetchParam true "获取群聊配置参数"
// @Success 200 {object} tables.Chatroom
// @Router /bot/chatroom/detail [post]
func (*Chatroom) detail(c *gin.Context) {

	var rq *chatroom.FetchParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if res, err := chatroom.Fetch(rq); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// @Summary 添加群聊配置
// @Produce json
// @Tags BOT::群聊配置
// @Param body body chatroom.CreateParam true "添加群聊配置参数"
// @Success 200
// @Router /bot/chatroom/create [post]
func (*Chatroom) create(c *gin.Context) {

	var rq *chatroom.CreateParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if id, err := chatroom.Create(rq); err == nil {
		c.Set("Message", "添加成功")
		c.Set("Payload", id)
		robot.Reset()
	} else {
		c.Set("Error", err)
	}

}

// @Summary 修改群聊配置
// @Produce json
// @Tags BOT::群聊配置
// @Param body body chatroom.UpdateParam true "修改群聊配置参数"
// @Success 200
// @Router /bot/chatroom/update [post]
func (*Chatroom) update(c *gin.Context) {

	var rq *chatroom.UpdateParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if err := chatroom.Update(rq); err == nil {
		c.Set("Message", "更新成功")
		robot.Reset()
	} else {
		c.Set("Error", err)
	}

}

// @Summary 删除群聊配置
// @Produce json
// @Tags BOT::群聊配置
// @Param body body chatroom.DeleteParam true "删除群聊配置参数"
// @Success 200
// @Router /bot/chatroom/delete [post]
func (*Chatroom) delete(c *gin.Context) {

	var rq *chatroom.DeleteParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if err := chatroom.Delete(rq); err == nil {
		c.Set("Message", "删除成功")
		robot.Reset()
	} else {
		c.Set("Error", err)
	}

}
