package wrobot

import (
	"github.com/gin-gonic/gin"

	"github.com/opentdp/wrest-chat/dbase/cronjob"
	"github.com/opentdp/wrest-chat/wclient/crond"
)

type Cronjob struct{}

// @Summary 计划任务列表
// @Produce json
// @Tags BOT::计划任务
// @Param body body cronjob.FetchAllParam true "获取计划任务列表参数"
// @Success 200 {object} []tables.Cronjob
// @Router /bot/cronjob/list [post]
func (*Cronjob) list(c *gin.Context) {

	var rq *cronjob.FetchAllParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if lst, err := cronjob.FetchAll(rq); err == nil {
		c.Set("Payload", lst)
	} else {
		c.Set("Error", err)
	}

}

// @Summary 获取计划任务
// @Produce json
// @Tags BOT::计划任务
// @Param body body cronjob.FetchParam true "获取计划任务参数"
// @Success 200 {object} tables.Cronjob
// @Router /bot/cronjob/detail [post]
func (*Cronjob) detail(c *gin.Context) {

	var rq *cronjob.FetchParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if res, err := cronjob.Fetch(rq); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

// @Summary 添加计划任务
// @Produce json
// @Tags BOT::计划任务
// @Param body body cronjob.CreateParam true "添加计划任务参数"
// @Success 200
// @Router /bot/cronjob/create [post]
func (*Cronjob) create(c *gin.Context) {

	var rq *cronjob.CreateParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if id, err := cronjob.Create(rq); err == nil {
		c.Set("Message", "添加成功")
		c.Set("Payload", id)
		crond.NewById(id)
	} else {
		c.Set("Error", err)
	}

}

// @Summary 修改计划任务
// @Produce json
// @Tags BOT::计划任务
// @Param body body cronjob.UpdateParam true "修改计划任务参数"
// @Success 200
// @Router /bot/cronjob/update [post]
func (*Cronjob) update(c *gin.Context) {

	var rq *cronjob.UpdateParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if err := cronjob.Update(rq); err == nil {
		c.Set("Message", "更新成功")
		crond.RedoById(rq.Rd)
	} else {
		c.Set("Error", err)
	}

}

// @Summary 删除计划任务
// @Produce json
// @Tags BOT::计划任务
// @Param body body cronjob.DeleteParam true "删除计划任务参数"
// @Success 200
// @Router /bot/cronjob/delete [post]
func (*Cronjob) delete(c *gin.Context) {

	var rq *cronjob.DeleteParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	crond.UndoById(rq.Rd) // 先从计划列表中删除

	if err := cronjob.Delete(rq); err == nil {
		c.Set("Message", "删除成功")
	} else {
		c.Set("Error", err)
	}

}

// @Summary 触发计划任务
// @Produce json
// @Tags BOT::计划任务
// @Param body body cronjob.FetchParam true "触发计划任务参数"
// @Success 200
// @Router /bot/cronjob/execute [post]
func (*Cronjob) execute(c *gin.Context) {

	var rq *cronjob.FetchParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	} else if rq.Rd == 0 {
		c.Set("Error", "参数错误")
		return
	}

	if err := crond.Execute(rq.Rd); err == nil {
		c.Set("Message", "触发完成")
	} else {
		c.Set("Error", err)
	}

}

// @Summary 计划任务状态
// @Produce json
// @Tags BOT::计划任务
// @Success 200
// @Router /bot/cronjob/status [post]
func (*Cronjob) status(c *gin.Context) {

	c.Set("Payload", crond.GetEntries())

}
