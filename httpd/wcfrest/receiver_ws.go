package wcfrest

import (
	"errors"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/recovery"

	"github.com/opentdp/wechat-rest/wcferry"
)

var socketReceiverKey = ""
var socketReceiverList = map[*websocket.Conn]*WebSocketMutex{}

type WebSocketMutex struct {
	*websocket.Conn
	mu sync.Mutex
}

func (w *WebSocketMutex) WriteJSON(v any) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.Conn.WriteJSON(v)
}

func (wc *Controller) enableSocketReceiver(ws *websocket.Conn) error {

	logman.Warn("enable receiver", "socket", ws.RemoteAddr().String())

	if socketReceiverList[ws] != nil {
		return errors.New("socket already exists")
	}

	wm := &WebSocketMutex{ws, sync.Mutex{}}

	if len(socketReceiverList) == 0 {
		key, err := wc.EnrollReceiver(true, func(msg *wcferry.WxMsg) {
			defer recovery.Handler()
			ret := wcferry.ParseWxMsg(msg)
			for s := range socketReceiverList {
				logman.Info("call receiver", "addr", s.RemoteAddr(), "Id", ret.Id)
				s.WriteJSON(ret)
			}
		})
		if err != nil {
			return err
		}
		socketReceiverKey = key
	}

	socketReceiverList[ws] = wm
	return nil

}

func (wc *Controller) disableSocketReceiver(ws *websocket.Conn) error {

	logman.Warn("disable receiver", "addr", ws.RemoteAddr())

	if socketReceiverList[ws] == nil {
		return errors.New("socket not exists")
	}

	delete(socketReceiverList, ws)
	if len(socketReceiverList) == 0 {
		return wc.DisableReceiver(socketReceiverKey)
	}

	return nil

}
