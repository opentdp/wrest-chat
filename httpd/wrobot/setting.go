package wrobot

import (
	"github.com/gin-gonic/gin"

	"github.com/opentdp/wrest-chat/dbase/setting"
)

type Setting struct{}

// @Summary 全局配置列表
// @Produce json
// @Tags BOT::全局配置
// @Param body body setting.FetchAllParam true "获取全局配置列表参数"
// @Success 200 {array} tables.Setting
// @Router /bot/setting/list [post]
func (*Setting) list(c *gin.Context) {

	var rq *setting.FetchAllParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if lst, err := setting.FetchAll(rq); err == nil {
		c.Set("Payload", lst)
	} else {
		c.Set("Error", err)
	}

}

// @Summary 获取全局配置
// @Produce json
// @Tags BOT::全局配置
// @Param body body setting.FetchParam true "获取全局配置参数"
// @Success 200 {object} tables.Setting
// @Router /bot/setting/detail [post]
func (*Setting) detail(c *gin.Context) {

	var rq *setting.FetchParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if res, err := setting.Fetch(rq); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// @Summary 添加全局配置
// @Produce json
// @Tags BOT::全局配置
// @Param body body setting.CreateParam true "添加全局配置参数"
// @Success 200
// @Router /bot/setting/create [post]
func (*Setting) create(c *gin.Context) {

	var rq *setting.CreateParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if id, err := setting.Create(rq); err == nil {
		c.Set("Message", "添加成功")
		c.Set("Payload", id)
		setting.Laod()
	} else {
		c.Set("Error", err)
	}

}

// @Summary 修改全局配置
// @Produce json
// @Tags BOT::全局配置
// @Param body body setting.UpdateParam true "修改全局配置参数"
// @Success 200
// @Router /bot/setting/update [post]
func (*Setting) update(c *gin.Context) {

	var rq *setting.UpdateParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if err := setting.Update(rq); err == nil {
		c.Set("Message", "更新成功")
		setting.Laod()
	} else {
		c.Set("Error", err)
	}

}

// @Summary 删除全局配置
// @Produce json
// @Tags BOT::全局配置
// @Param body body setting.DeleteParam true "删除全局配置参数"
// @Success 200
// @Router /bot/setting/delete [post]
func (*Setting) delete(c *gin.Context) {

	var rq *setting.DeleteParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if err := setting.Delete(rq); err == nil {
		c.Set("Message", "删除成功")
		setting.Laod()
	} else {
		c.Set("Error", err)
	}

}
