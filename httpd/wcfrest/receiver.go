package wcfrest

import (
	"errors"

	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/request"
	"golang.org/x/net/websocket"

	"github.com/opentdp/wechat-rest/wcferry"
)

var urlReceiverList = map[string]bool{}
var wsReceiverList = map[*websocket.Conn]bool{}

func (wc *Controller) enableUrlReceiver(url string) error {

	logman.Info("enable receiver", "url", url)

	if len(urlReceiverList) == 0 {
		err := wc.EnrollReceiver(true, func(msg *wcferry.WxMsg) {
			ret := wcferry.ParseWxMsg(msg)
			for u := range urlReceiverList {
				logman.Info("call receiver", "url", u, "Id", ret.Id)
				go request.JsonPost(u, ret, request.H{})
			}
		})
		if err != nil {
			return err
		}
	}

	if _, ok := urlReceiverList[url]; ok {
		return errors.New("url already exists")
	}

	urlReceiverList[url] = true
	return nil

}

func (wc *Controller) disableUrlReceiver(url string) error {

	logman.Info("disable receiver", "url", url)

	if _, ok := urlReceiverList[url]; !ok {
		return errors.New("url not exists")
	}

	delete(urlReceiverList, url)

	if len(urlReceiverList) == 0 && len(wsReceiverList) == 0 {
		if err := wc.DisableReceiver(false); err != nil {
			return err
		}
	}

	return nil

}

func (wc *Controller) enableWsReceiver(ws *websocket.Conn) error {

	logman.Info("enable receiver", "addr", ws.RemoteAddr().String())

	if len(wsReceiverList) == 0 {
		err := wc.EnrollReceiver(true, func(msg *wcferry.WxMsg) {
			ret := wcferry.ParseWxMsg(msg)
			for w := range wsReceiverList {
				logman.Info("call receiver", "addr", ws.RemoteAddr().String(), "Id", ret.Id)
				go websocket.JSON.Send(w, ret)
			}
		})
		if err != nil {
			return err
		}
	}

	if _, ok := wsReceiverList[ws]; ok {
		return errors.New("ws already exists")
	}

	wsReceiverList[ws] = true
	return nil

}

func (wc *Controller) disableWsReceiver(ws *websocket.Conn) error {

	logman.Info("disable receiver", "addr", ws.RemoteAddr().String())

	if _, ok := wsReceiverList[ws]; !ok {
		return errors.New("ws not exists")
	}

	delete(wsReceiverList, ws)

	if len(wsReceiverList) == 0 && len(urlReceiverList) == 0 {
		if err := wc.DisableReceiver(false); err != nil {
			return err
		}
	}

	return nil

}
