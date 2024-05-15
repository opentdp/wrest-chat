package robot

import (
	"github.com/opentdp/go-helper/command"
	"github.com/opentdp/go-helper/logman"

	"github.com/opentdp/wrest-chat/dbase/keyword"
	"github.com/opentdp/wrest-chat/wcferry"
	"github.com/opentdp/wrest-chat/wclient"
)

func cmddHandler() []*Handler {

	cmds := []*Handler{}

	keywords, err := keyword.FetchAll(&keyword.FetchAllParam{Group: "command"})
	if err != nil || len(keywords) == 0 {
		return cmds
	}

	for k, v := range keywords {
		v := v // copy
		if v.Remark == "" {
			v.Remark = "神秘指令"
		}
		cmds = append(cmds, &Handler{
			Level:    v.Level,
			Order:    410 + int32(k),
			Roomid:   v.Roomid,
			Command:  v.Phrase,
			Describe: v.Remark,
			Callback: func(msg *wcferry.WxMsg) string {
				exec := v.Target + " " + msg.Content
				output, err := command.Exec(&command.ExecPayload{
					Name:          "Handler:" + v.Phrase,
					CommandType:   "EXEC",
					WorkDirectory: ".",
					Content:       exec,
				})
				if err != nil {
					logman.Error("cmd: "+v.Phrase, "error", err)
				}
				wclient.SendFlexMsg(output, msg.Sender, msg.Roomid)
				return ""
			},
		})
	}

	return cmds

}
