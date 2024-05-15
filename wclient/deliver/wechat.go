package deliver

import (
	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/wrest-chat/wclient"
)

// 将执行结果投递到微信

func wechatMessage(args []string, message string) int32 {

	roomid := args[0]

	wxid := ""
	if len(args) > 1 {
		wxid = args[1]
	}

	wc := wclient.Register()
	if wc == nil {
		logman.Error("deliver", "error", "wclient is nil")
		return -1
	}

	return wclient.SendFlexMsg(message, wxid, roomid)

}
