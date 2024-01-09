package wcferry

import (
	"fmt"
	"strings"

	"github.com/clbanning/mxj"
	"github.com/opentdp/go-helper/request"
	"github.com/opentdp/go-helper/strutil"
)

// 消息解析结果
type WxMsgParsed struct {
	*WxMsg      // 消息原始数据
	Content any `json:"content,omitempty"`
	Xml     any `json:"xml,omitempty"`
}

// 解析消息数据
// param msg *WxMsg 消息
// return *WxMsgParsed 转换后的消息
func WxMsgParser(msg *WxMsg) *WxMsgParsed {
	var ret = &WxMsgParsed{msg, msg.Content, msg.Xml}
	// preset
	var str string
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

// 打印接收到的消息
// param msg *WxMsgParsed 消息
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

// 获取网络文件
// param url string 文件URL或路径
// return string 失败则返回空字符串
func DownloadFile(url string) string {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		if tmp, err := request.Download(url, "", false); err == nil {
			return tmp
		}
	}
	return ""
}
