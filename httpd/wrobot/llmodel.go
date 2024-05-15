package wrobot

import (
	"github.com/gin-gonic/gin"

	"github.com/opentdp/wrest-chat/dbase/llmodel"
	"github.com/opentdp/wrest-chat/wclient/robot"
)

type LLModel struct{}

// @Summary 模型列表
// @Produce json
// @Tags BOT::大语言模型
// @Param body body llmodel.FetchAllParam true "获取模型列表参数"
// @Success 200 {array} tables.LLModel
// @Router /bot/llmodel/list [post]
func (*LLModel) list(c *gin.Context) {

	var rq *llmodel.FetchAllParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if lst, err := llmodel.FetchAll(rq); err == nil {
		c.Set("Payload", lst)
	} else {
		c.Set("Error", err)
	}

}

// @Summary 获取模型
// @Produce json
// @Tags BOT::大语言模型
// @Param body body llmodel.FetchParam true "获取模型参数"
// @Success 200 {object} tables.LLModel
// @Router /bot/llmodel/detail [post]
func (*LLModel) detail(c *gin.Context) {

	var rq *llmodel.FetchParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if res, err := llmodel.Fetch(rq); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// @Summary 添加模型
// @Produce json
// @Tags BOT::大语言模型
// @Param body body llmodel.CreateParam true "添加模型参数"
// @Success 200
// @Router /bot/llmodel/create [post]
func (*LLModel) create(c *gin.Context) {

	var rq *llmodel.CreateParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if id, err := llmodel.Create(rq); err == nil {
		c.Set("Message", "添加成功")
		c.Set("Payload", id)
		robot.Reset()
	} else {
		c.Set("Error", err)
	}

}

// @Summary 修改模型
// @Produce json
// @Tags BOT::大语言模型
// @Param body body llmodel.UpdateParam true "修改模型参数"
// @Success 200
// @Router /bot/llmodel/update [post]
func (*LLModel) update(c *gin.Context) {

	var rq *llmodel.UpdateParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if err := llmodel.Update(rq); err == nil {
		c.Set("Message", "更新成功")
		robot.Reset()
	} else {
		c.Set("Error", err)
	}

}

// @Summary 删除模型
// @Produce json
// @Tags BOT::大语言模型
// @Param body body llmodel.DeleteParam true "删除模型参数"
// @Success 200
// @Router /bot/llmodel/delete [post]
func (*LLModel) delete(c *gin.Context) {

	var rq *llmodel.DeleteParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if err := llmodel.Delete(rq); err == nil {
		c.Set("Message", "删除成功")
		robot.Reset()
	} else {
		c.Set("Error", err)
	}

}
