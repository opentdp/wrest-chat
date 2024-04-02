package robot

import (
	"os"
	"path/filepath"

	"github.com/opentdp/go-helper/command"
	"github.com/opentdp/go-helper/filer"
	"github.com/opentdp/go-helper/logman"

	"github.com/opentdp/wrest-chat/args"
	"github.com/opentdp/wrest-chat/dbase/keyword"
	"github.com/opentdp/wrest-chat/dbase/message"
	"github.com/opentdp/wrest-chat/dbase/setting"
	"github.com/opentdp/wrest-chat/wcferry"
	"github.com/opentdp/wrest-chat/wclient"
)

func receiver3(msg *wcferry.WxMsg) {

	// 自动保存图片
	if setting.AutoSaveImage && msg.Extra != "" {
		msgImage(msg.Id, msg.Extra)
	}

	// 外部图片处理插件
	keywords, err := keyword.FetchAll(&keyword.FetchAllParam{Group: "imagefn"})
	if err == nil && len(keywords) > 0 {
		img := msgImage(msg.Id, msg.Extra)
		for _, v := range keywords {
			if groupLimit(msg, v.Level, v.Roomid) {
				continue
			}
			output, err := command.Exec(&command.ExecPayload{
				Name:          "Imager:" + v.Phrase,
				CommandType:   "EXEC",
				WorkDirectory: ".",
				Content:       v.Target + " " + img,
			})
			if err != nil {
				logman.Error("cmd: "+v.Phrase, "error", err)
			}
			wclient.SendFlexMsg(output, msg.Sender, msg.Roomid)
		}
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
