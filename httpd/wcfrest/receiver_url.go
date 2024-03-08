package wcfrest

import (
	"errors"

	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/request"

	"github.com/opentdp/wechat-rest/wcferry"
)

var urlReceiverKey = ""
var urlReceiverList = map[string]bool{}

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
				request.JsonPost(u, ret, request.H{})
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
