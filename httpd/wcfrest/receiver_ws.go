package wcfrest

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/opentdp/go-helper/logman"

	"github.com/opentdp/wrest-chat/wcferry"
)

func socketReceiver(ws *websocket.Conn) wcferry.MsgConsumer {

	mu := sync.Mutex{}

	return func(msg *wcferry.WxMsg) {
		mu.Lock()
		defer mu.Unlock()
		ws.WriteJSON(wcferry.ParseWxMsg(msg))
	}

}

// @Summary 推送消息到Socket
// @Produce json
// @Tags WCF::消息推送
// @Success 101 {string} string "Switching Protocols 响应"
// @Router /wcf/socket_receiver [get]
func (wc *Controller) socketReceiver(c *gin.Context) {

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logman.Error("websocket upgrade", "error", err)
		c.Set("Error", err)
		return
	}

	defer ws.Close()

	logman.Warn("enable receiver", "socket", ws.RemoteAddr())
	key, err := wc.EnrollReceiver(true, socketReceiver(ws))
	if err != nil {
		logman.Error("enroll receiver", "error", err)
		c.Set("Error", err)
		return
	}

	defer wc.DisableReceiver(key)

	for {
		if _, _, err := ws.ReadMessage(); err != nil {
			break
		}
	}

}
