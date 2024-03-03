package cronjob

import (
	"strings"

	"github.com/opentdp/wechat-rest/wclient"
)

func MsgDeliver(deliver, message string) {

	deliver = strings.TrimSpace(deliver)
	message = strings.TrimSpace(message)

	args := strings.Split(deliver, ",")
	if len(args) < 2 {
		return
	}

	switch args[0] {
	case "wechat":
		wechatMessage(args[1:], message)
	}

}

// 将执行结果投递到微信

func wechatMessage(args []string, message string) int32 {

	wxid := args[0]

	roomid := ""
	if len(args) > 1 {
		roomid = args[1]
	}

	wc := wclient.Register()
	if wc == nil {
		logger.Error("cron:deliver", "error", "wclient is nil")
		return -1
	}

	return wc.SendMessage(wxid, roomid, message)

}
