package robot

import (
	"os"
	"path/filepath"

	"github.com/opentdp/go-helper/filer"
	"github.com/opentdp/go-helper/logman"

	"github.com/opentdp/wrest-chat/args"
	"github.com/opentdp/wrest-chat/dbase/message"
	"github.com/opentdp/wrest-chat/dbase/setting"
	"github.com/opentdp/wrest-chat/wcferry"
)

// 自动保存图片
func receiver3(msg *wcferry.WxMsg) {

	if setting.AutoSaveImage && msg.Extra != "" {
		msgImage(msg.Id, msg.Extra)
	}

}

func msgImage(id uint64, extra string) string {

	// 从数据库获取
	if args.Wcf.MsgStore && extra == "" {
		res, _ := message.Fetch(&message.FetchParam{Id: id})
		if res.Remark != "" {
			return res.Remark
		}
		extra = res.Extra
	}

	// 获取存储路径
	target, err := filepath.Abs(args.Web.Storage)
	if err != nil {
		if self, err := os.Executable(); err == nil {
			target = filepath.Dir(self)
		}
	}
	target = filepath.Join(target, "chat-images")
	if !filer.Exists(target) {
		os.MkdirAll(target, 0755)
	}

	// 从消息中获取
	fp, err := wc.CmdClient.DownloadImage(id, extra, target, 15)
	if err != nil || fp == "" {
		logman.Error("image save failed", "err", err)
		return ""
	}

	// 保存到数据库
	logman.Info("image saved", "path", fp)
	if args.Wcf.MsgStore {
		message.Update(&message.UpdateParam{Id: id, Remark: fp})
	}

	return fp

}
