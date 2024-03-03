package cronjob

import (
	"strings"

	"github.com/opentdp/wechat-rest/wclient"
)

func MsgDeliver(deliver, message string) {

	deliver = strings.TrimSpace(deliver)
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

	if len(args) == 0 {
		return -1
	}

	if len(args) == 1 {
		args[1] = ""
	}

	return wclient.Register().SendMessage(args[0], args[1], message)

}
