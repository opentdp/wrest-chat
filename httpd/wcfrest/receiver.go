package wcfrest

import (
	"errors"

	"github.com/gorilla/websocket"
	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/request"

	"github.com/opentdp/wechat-rest/wcferry"
)

var urlReceiverKey = ""
var urlReceiverList = map[string]bool{}

var socketReceiverKey = ""
var socketReceiverList = map[*websocket.Conn]bool{}

func (wc *Controller) enableUrlReceiver(url string) error {

	logman.Warn("enable receiver", "url", url)

	if urlReceiverList[url] {
		return errors.New("url already exists")
	}

	if len(urlReceiverList) == 0 {
		key, err := wc.EnrollReceiver(true, func(msg *wcferry.WxMsg) {
			ret := wcferry.ParseWxMsg(msg)
			for u := range urlReceiverList {
				logman.Info("call receiver", "url", u, "Id", ret.Id)
				go request.JsonPost(u, ret, request.H{})
			}
		})
		if err != nil {
			return err
		}
		urlReceiverKey = key
	}

	urlReceiverList[url] = true
	return nil

}

func (wc *Controller) disableUrlReceiver(url string) error {

	logman.Warn("disable receiver", "url", url)

	if !urlReceiverList[url] {
		return errors.New("url not exists")
	}

	delete(urlReceiverList, url)
	if len(urlReceiverList) == 0 {
		return wc.DisableReceiver(urlReceiverKey)
	}

	return nil

}

func (wc *Controller) enableSocketReceiver(ws *websocket.Conn) error {

	logman.Warn("enable receiver", "socket", ws.RemoteAddr().String())

	if socketReceiverList[ws] {
		return errors.New("socket already exists")
	}

	if len(socketReceiverList) == 0 {
		key, err := wc.EnrollReceiver(true, func(msg *wcferry.WxMsg) {
			ret := wcferry.ParseWxMsg(msg)
			for s := range socketReceiverList {
				logman.Info("call receiver", "socket", s.RemoteAddr().String(), "Id", ret.Id)
				go s.WriteJSON(ret)
			}
		})
		if err != nil {
			return err
		}
		socketReceiverKey = key
	}

	socketReceiverList[ws] = true
	return nil

}

func (wc *Controller) disableSocketReceiver(ws *websocket.Conn) error {

	logman.Warn("disable receiver", "socket", ws.RemoteAddr().String())

	if !socketReceiverList[ws] {
		return errors.New("socket not exists")
	}

	delete(socketReceiverList, ws)
	if len(socketReceiverList) == 0 {
		return wc.DisableReceiver(socketReceiverKey)
	}

	return nil

}
