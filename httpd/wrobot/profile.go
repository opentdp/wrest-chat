package wrobot

import (
	"github.com/gin-gonic/gin"

	"github.com/opentdp/wrest-chat/dbase/profile"
	"github.com/opentdp/wrest-chat/wclient/robot"
)

type Profile struct{}

// @Summary 用户配置列表
// @Produce json
// @Tags BOT::用户配置
// @Param body body profile.FetchAllParam true "获取用户配置列表参数"
// @Success 200 {array} tables.Profile
// @Router /bot/profile/list [post]
func (*Profile) list(c *gin.Context) {

	var rq *profile.FetchAllParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if lst, err := profile.FetchAll(rq); err == nil {
		c.Set("Payload", lst)
	} else {
		c.Set("Error", err)
	}

}

// @Summary 获取用户配置
// @Produce json
// @Tags BOT::用户配置
// @Param body body profile.FetchParam true "获取用户配置参数"
// @Success 200 {object} tables.Profile
// @Router /bot/profile/detail [post]
func (*Profile) detail(c *gin.Context) {

	var rq *profile.FetchParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if res, err := profile.Fetch(rq); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// @Summary 添加用户配置
// @Produce json
// @Tags BOT::用户配置
// @Param body body profile.CreateParam true "添加用户配置参数"
// @Success 200
// @Router /bot/profile/create [post]
func (*Profile) create(c *gin.Context) {

	var rq *profile.CreateParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if id, err := profile.Create(rq); err == nil {
		c.Set("Message", "添加成功")
		c.Set("Payload", id)
		robot.Reset()
	} else {
		c.Set("Error", err)
	}

}

// @Summary 修改用户配置
// @Produce json
// @Tags BOT::用户配置
// @Param body body profile.UpdateParam true "修改用户配置参数"
// @Success 200
// @Router /bot/profile/update [post]
func (*Profile) update(c *gin.Context) {

	var rq *profile.UpdateParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if err := profile.Update(rq); err == nil {
		c.Set("Message", "更新成功")
		robot.Reset()
	} else {
		c.Set("Error", err)
	}

}

// @Summary 删除用户配置
// @Produce json
// @Tags BOT::用户配置
// @Param body body profile.DeleteParam true "删除用户配置参数"
// @Success 200
// @Router /bot/profile/delete [post]
func (*Profile) delete(c *gin.Context) {

	var rq *profile.DeleteParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if err := profile.Delete(rq); err == nil {
		c.Set("Message", "删除成功")
		robot.Reset()
	} else {
		c.Set("Error", err)
	}

}
