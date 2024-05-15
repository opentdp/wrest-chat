package robot

import (
	"net/url"
	"strings"

	"github.com/opentdp/wrest-chat/dbase/setting"
	"github.com/opentdp/wrest-chat/wcferry"
	"github.com/opentdp/wrest-chat/wclient"
)

func apiHandler() []*Handler {

	cmds := []*Handler{}

	if len(setting.ApiEndpoint) < 10 {
		return cmds
	}

	cmds = append(cmds, &Handler{
		Level:    -1,
		Order:    200,
		Roomid:   "*",
		Command:  "/api",
		Describe: "调用查询接口",
		Callback: func(msg *wcferry.WxMsg) string {
			query := []string{"help"}
			if msg.Content != "" {
				query = strings.SplitN(msg.Content, " ", 2)
				if len(query) > 1 {
					query[1] = url.QueryEscape(query[1])
				}
			}
			url := setting.ApiEndpoint + strings.Join(query, "/")
			wclient.ApiRequestMsg(url, msg.Sender, msg.Roomid)
			return ""
		},
	})

	return cmds

}
