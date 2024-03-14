package robot

import (
	"github.com/opentdp/go-helper/command"
	"github.com/opentdp/go-helper/logman"

	"github.com/opentdp/wechat-rest/dbase/keyword"
	"github.com/opentdp/wechat-rest/wcferry"
)

func cmddHandler() []*Handler {

	cmds := []*Handler{}

	keywords, err := keyword.FetchAll(&keyword.FetchAllParam{Group: "command"})
	if err != nil || len(keywords) == 0 {
		return cmds
	}

	for k, v := range keywords {
		v := v // copy
		cmds = append(cmds, &Handler{
			Level:    v.Level,
			Order:    400 + int32(k),
			Roomid:   v.Roomid,
			Command:  v.Phrase,
			Describe: "执行命令 " + v.Phrase,
			Callback: func(msg *wcferry.WxMsg) string {
				exec := v.Target + " " + msg.Content
				output, err := command.Exec(&command.ExecPayload{
					Name:        "Handler:" + v.Phrase,
					CommandType: "EXEC",
					Content:     exec,
				})
				if err != nil {
					logman.Error("cmd: "+v.Phrase, "error", err)
				}
				return output
			},
		})
	}

	return cmds

}
