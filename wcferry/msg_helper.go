package wcferry

import (
	"encoding/xml"
	"fmt"
	"html"
	"strings"

	"github.com/clbanning/mxj"
	"github.com/opentdp/go-helper/strutil"

	"github.com/opentdp/wrest-chat/wcferry/types"
)

type FlexWxMsg struct {
	*WxMsg      // 消息原始数据
	Content any `json:"content,omitempty"`
	Xml     any `json:"xml,omitempty"`
}

// 解析消息数据
// param msg *WxMsg 消息
// return *FlexWxMsg 转换后的消息
func ParseWxMsg(msg *WxMsg) *FlexWxMsg {
	ret := &FlexWxMsg{msg, msg.Content, msg.Xml}
	// preset
	str := ""
	mxj.SetAttrPrefix("")
	// c.Content
	str = strings.TrimSpace(msg.Content)
	xmlPrefixes := []string{"<?xml", "<sysmsg", "<msg"}
	for _, prefix := range xmlPrefixes {
		if strings.HasPrefix(str, prefix) {
			mv, err := mxj.NewMapXml([]byte(str))
			if err == nil {
				ret.Content = mv
			}
			break
		}
	}
	// c.Xml
	str = strings.TrimSpace(msg.Xml)
	if strings.HasPrefix(str, "<") {
		mv, err := mxj.NewMapXml([]byte(str))
		if err == nil {
			ret.Xml = mv
		}
	}
	// return
	return ret
}

// 解析聊天记录
// param str string 消息内容
// return *types.RecordInfo 聊天记录
func ParseWxMsgRecord(str string) (*types.RecordInfo, error) {
	// 解析消息内容
	content := types.MsgContent49{}
	err := xml.Unmarshal([]byte(str), &content)
	if err != nil || content.AppMsg.Type != 19 {
		return nil, fmt.Errorf("不支持的记录格式")
	}
	// 解析聊天记录
	record := types.RecordInfo{}
	itemXml := html.UnescapeString(content.AppMsg.RecordItem)
	err = xml.Unmarshal([]byte(itemXml), &record)
	return &record, err
}

// 打印接收到的消息
// param msg *FlexWxMsg 消息
func WxMsgPrinter(msg *WxMsg) {
	rs := "\n=== New Message ===\n"
	if msg.Id > 0 {
		rs += fmt.Sprintf("::Id:: %d\n", msg.Id)
	}
	if msg.Type > 0 {
		rs += fmt.Sprintf("::Type:: %d\n", msg.Type)
	}
	if msg.Roomid != "" {
		rs += fmt.Sprintf("::Roomid:: %s\n", msg.Roomid)
	}
	if msg.Sender != "" {
		rs += fmt.Sprintf("::Sender:: %v\n", msg.Sender)
	}
	if msg.Content != "" {
		rs += fmt.Sprintf("::Content::\n%s\n", strutil.Dedent(msg.Content))
	}
	if msg.Xml != "" {
		rs += fmt.Sprintf("::Xml::\n%s\n", strutil.Dedent(msg.Xml))
	}
	if msg.Extra != "" {
		rs += fmt.Sprintf("::Extra:: %s\n", msg.Extra)
	}
	fmt.Print(rs, "=== End Message ===\n")
}
