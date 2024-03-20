package wcfrest

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/request"

	"github.com/opentdp/wrest-chat/wcferry"
)

var urlReceiverList = map[string]string{}

func urlReciever(url string) wcferry.MsgConsumer {

	return func(msg *wcferry.WxMsg) {
		ret := wcferry.ParseWxMsg(msg)
		request.JsonPost(url, ret, request.H{})
	}

}

// @Summary 开启推送消息到URL
// @Produce json
// @Tags WCF::消息推送
// @Param body body ReceiverRequest true "推送消息到URL参数"
// @Success 200 {object} CommonPayload
// @Router /wcf/enable_receiver [post]
func (wc *Controller) enabledReceiver(c *gin.Context) {

	var req ReceiverRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	if !strings.HasPrefix(req.Url, "http") {
		c.Set("Error", "url must start with http(s)://")
		return
	}

	if urlReceiverList[req.Url] != "" {
		c.Set("Error", "url already exists")
		return
	}

	logman.Warn("enable receiver", "url", req.Url)
	key, err := wc.EnrollReceiver(true, urlReciever(req.Url))
	if err != nil {
		c.Set("Error", err)
		return
	}

	urlReceiverList[req.Url] = key

	c.Set("Payload", CommonPayload{
		Success: err == nil,
		Result:  key,
		Error:   err,
	})

}

type ReceiverRequest struct {
	// 接收推送消息的 url
	Url string `json:"url"`
}

// @Summary 关闭推送消息到URL
// @Produce json
// @Tags WCF::消息推送
// @Param body body ReceiverRequest true "推送消息到URL参数"
// @Success 200 {object} CommonPayload
// @Router /wcf/disable_receiver [post]
func (wc *Controller) disableReceiver(c *gin.Context) {

	var req ReceiverRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("Error", err)
		return
	}

	logman.Warn("disable receiver", "url", req.Url)
	if urlReceiverList[req.Url] == "" {
		c.Set("Error", "url not exists")
		return
	}

	err := wc.DisableReceiver(urlReceiverList[req.Url])
	delete(urlReceiverList, req.Url)

	c.Set("Payload", CommonPayload{
		Success: err == nil,
		Error:   err,
	})

}
