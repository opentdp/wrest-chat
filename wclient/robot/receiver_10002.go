package robot

import (
	"encoding/xml"
	"strconv"
	"strings"

	"github.com/opentdp/wechat-rest/dbase/chatroom"
	"github.com/opentdp/wechat-rest/dbase/message"
	"github.com/opentdp/wechat-rest/dbase/setting"
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wcferry/types"
)

// 处理撤回消息
func receiver10002(msg *wcferry.WxMsg) {

	var output string

	// 获取撤回提示
	if msg.IsGroup {
		room, _ := chatroom.Fetch(&chatroom.FetchParam{Roomid: msg.Roomid})
		output = room.RevokeMsg
	} else {
		output = setting.RevokeMsg
	}

	// 防撤回提示已关闭
	if len(output) < 2 {
		return
	}

	// 解析已撤回的消息
	revoke := types.MsgContent10002{}
	err := xml.Unmarshal([]byte(msg.Content), &revoke)
	if err != nil || revoke.RevokeMsg.NewMsgID == "" {
		return
	}

	// 获取已撤回消息的 Id
	id, err := strconv.Atoi(revoke.RevokeMsg.NewMsgID)
	if err != nil || id == 0 {
		return
	}

	// 取回已撤回的消息内容
	origin, err := message.Fetch(&message.FetchParam{Id: uint64(id)})
	if err != nil || origin.Content == "" {
		return
	}

	// 提示已撤回的消息内容
	str := strings.TrimSpace(origin.Content)
	xmlPrefixes := []string{"<?xml", "<sysmsg", "<msg"}
	for _, prefix := range xmlPrefixes {
		if strings.HasPrefix(str, prefix) {
			str = ""
		}
	}

	if str != "" {
		output += "\n-------\n" + str
		reply(msg, output)
		return
	}

	if origin.Type == 3 {
		if origin.Remark != "" {
			if origin.IsGroup {
				wc.CmdClient.SendImg(origin.Remark, origin.Roomid)
			} else {
				wc.CmdClient.SendImg(origin.Remark, origin.Sender)
			}
		}
		output += "\n-------\n一张不可描述的图片"
		reply(msg, output)
		return
	}

	if origin.Type == 47 {
		output += "\n-------\n一个震惊四座的表情"
		reply(msg, output)
		return
	}

	if origin.Type == 49 {
		appmsg := types.MsgContent49{}
		err := xml.Unmarshal([]byte(origin.Content), &appmsg)
		if err == nil {
			switch appmsg.AppMsg.Type {
			case 6:
				output += "\n-------\n一份暗藏机密的文件"
			case 19:
				output += "\n-------\n多条来自异界的消息"
			case 57:
				output += "\n-------\n" + appmsg.AppMsg.Title
			default:
				output += "\n-------\n暂不支持回显的消息类型"
			}
			reply(msg, output)
			return
		}
	}

	output += "\n-------\n暂不支持回显的消息类型"
	reply(msg, output)

}
