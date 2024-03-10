package robot

import (
	"time"

	"github.com/opentdp/go-helper/logman"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/dbase/message"
	"github.com/opentdp/wechat-rest/dbase/setting"
	"github.com/opentdp/wechat-rest/wcferry"
)

// 自动保存图片
func receiver3(msg *wcferry.WxMsg) {

	if !setting.AutoSaveImage || msg.Extra == "" {
		return
	}

	time.Sleep(2 * time.Second) // 等待数据落库

	p, err := wc.CmdClient.DownloadImage(msg.Id, msg.Extra, "", 5)
	if err != nil || p == "" {
		logman.Error("image save failed", "err", err)
		return
	}

	logman.Info("image saved", "path", p)
	if args.Wcf.MsgStore {
		message.Update(&message.UpdateParam{Id: msg.Id, Remark: p})
	}

}
