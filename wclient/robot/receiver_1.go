package robot

import (
	"strings"

	"github.com/opentdp/wrest-chat/wcferry"
)

// 处理新消息
func receiver1(msg *wcferry.WxMsg) {

	// 处理聊天指令
	if msg.IsGroup || wcferry.ContactType(msg.Sender) == "好友" {
		output := ApplyHandlers(msg)
		if strings.Trim(output, "-") != "" {
			reply(msg, output)
		}
		return
	}

}
