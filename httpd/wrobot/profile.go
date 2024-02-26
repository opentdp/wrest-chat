package wrobot

import (
	"github.com/gin-gonic/gin"

	"github.com/opentdp/wechat-rest/dbase/profile"
)

type Profile struct{}

// @Summary 配置列表
// @Produce json
// @Tags BOT::用户配置
// @Param body body profile.FetchAllParam true "获取配置列表参数"
// @Success 200 {object} []tables.Profile
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

// @Summary 获取配置
// @Produce json
// @Tags BOT::用户配置
// @Param body body profile.FetchParam true "获取配置参数"
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

// @Summary 添加配置
// @Produce json
// @Tags BOT::用户配置
// @Param body body profile.CreateParam true "添加配置参数"
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
	} else {
		c.Set("Error", err)
	}

}

// @Summary 修改配置
// @Produce json
// @Tags BOT::用户配置
// @Param body body profile.UpdateParam true "修改配置参数"
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
	} else {
		c.Set("Error", err)
	}

}

// @Summary 删除配置
// @Produce json
// @Tags BOT::用户配置
// @Param body body profile.DeleteParam true "删除配置参数"
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
	} else {
		c.Set("Error", err)
	}

}
