package sundry

import (
	"github.com/gin-gonic/gin"

	"github.com/opentdp/wrest-chat/args"
	"github.com/opentdp/wrest-chat/wcferry"
	"github.com/opentdp/wrest-chat/wclient/robot"
)

type System struct{}

// @Summary 获取系统版本
// @Produce json
// @Tags API::系统
// @Success 200 {object} SystemVersion
// @Router /api/system/version [post]
func (s *System) version(c *gin.Context) {

	c.Set("Payload", SystemVersion{
		Version:       args.Version,
		BuildVersion:  args.BuildVersion,
		WcfVersion:    wcferry.Wcf_Version,
		WechatVersion: wcferry.Wechat_Version,
	})

}

type SystemVersion struct {
	Version       string `json:"version"`        // 系统版本
	BuildVersion  string `json:"build_version"`  // 系统编译版本
	WcfVersion    string `json:"wcf_version"`    // wcferry 版本
	WechatVersion string `json:"wechat_version"` // wechat 版本
}

// @Summary 获取可用指令
// @Tags API::系统
// @Param body body HandlersParam true "获取所有可用指令参数"
// @Success 200 {array} Handler
// @Router /api/system/handlers [post]
func (s *System) handlers(c *gin.Context) {

	var rq *HandlersParam
	if err := c.ShouldBind(&rq); err != nil {
		c.Set("Error", err)
		return
	}

	if rq.Reset {
		robot.Reset()
	}

	items := []Handler{}
	for _, v := range robot.GetHandlers() {
		items = append(items, Handler{
			Level:    v.Level,
			Order:    v.Order,
			Roomid:   v.Roomid,
			Command:  v.Command,
			Describe: v.Describe,
		})
	}

	c.Set("Payload", items)

}

type Handler struct {
	Level    int32  `json:"level"`    // 0:不限制 7:群管理 9:创始人
	Order    int32  `json:"order"`    // 排序，越小越靠前
	Roomid   string `json:"roomid"`   // 使用场景 [*:所有,-:私聊,+:群聊,其他:群聊]
	Command  string `json:"command"`  // 指令
	Describe string `json:"describe"` // 指令的描述信息
}

type HandlersParam struct {
	Reset bool `json:"reset"` // 是否重置机器人指令
}
