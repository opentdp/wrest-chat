package wrobot

import (
	"github.com/gin-gonic/gin"

	"github.com/opentdp/wechat-rest/dbase/keyword"
)

type Keyword struct{}

// @Summary 关键字列表
// @Produce json
// @Tags BOT::关键字
// @Param body body keyword.FetchAllParam true "获取关键字列表参数"
// @Success 200 {object} []tables.Keyword
// @Router /bot/keyword/list [post]
func (*Keyword) list(c *gin.Context) {

	var rq *keyword.FetchAllParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if lst, err := keyword.FetchAll(rq); err == nil {
		c.Set("Payload", lst)
	} else {
		c.Set("Error", err)
	}

}

// @Summary 获取关键字
// @Produce json
// @Tags BOT::关键字
// @Param body body keyword.FetchParam true "获取关键字参数"
// @Success 200 {object} tables.Keyword
// @Router /bot/keyword/detail [post]
func (*Keyword) detail(c *gin.Context) {

	var rq *keyword.FetchParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if res, err := keyword.Fetch(rq); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// @Summary 添加关键字
// @Produce json
// @Tags BOT::关键字
// @Param body body keyword.CreateParam true "添加关键字参数"
// @Success 200
// @Router /bot/keyword/create [post]
func (*Keyword) create(c *gin.Context) {

	var rq *keyword.CreateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if id, err := keyword.Create(rq); err == nil {
		c.Set("Message", "添加成功")
		c.Set("Payload", id)
	} else {
		c.Set("Error", err)
	}

}

// @Summary 修改关键字
// @Produce json
// @Tags BOT::关键字
// @Param body body keyword.UpdateParam true "修改关键字参数"
// @Success 200
// @Router /bot/keyword/update [post]
func (*Keyword) update(c *gin.Context) {

	var rq *keyword.UpdateParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if err := keyword.Update(rq); err == nil {
		c.Set("Message", "更新成功")
	} else {
		c.Set("Error", err)
	}

}

// @Summary 删除关键字
// @Produce json
// @Tags BOT::关键字
// @Param body body keyword.DeleteParam true "删除关键字参数"
// @Success 200
// @Router /bot/keyword/delete [post]
func (*Keyword) delete(c *gin.Context) {

	var rq *keyword.DeleteParam

	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if err := keyword.Delete(rq); err == nil {
		c.Set("Message", "删除成功")
	} else {
		c.Set("Error", err)
	}

}
