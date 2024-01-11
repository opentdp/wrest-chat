package wcferry

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/clbanning/mxj"
	"github.com/opentdp/go-helper/request"
	"github.com/opentdp/go-helper/strutil"
)

type FlexWxMsg struct {
	*WxMsg      // 消息原始数据
	Content any `json:"content,omitempty"`
	Xml     any `json:"xml,omitempty"`
}

type FlexDbField struct {
	Type    int32  `json:"type,omitempty"`    // 字段类型
	Column  string `json:"column,omitempty"`  // 字段名称
	Content any    `json:"content,omitempty"` // 字段内容
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

// 解析数据库字段
// param fieldType int 字段类型
// param content []byte 字段内容
// return any 解析结果
func ParseDbField(fieldType int32, content []byte) (any, error) {
	str := string(content)
	switch fieldType {
	case 1:
		return strconv.ParseInt(str, 10, 64)
	case 2:
		return strconv.ParseFloat(str, 64)
	case 4:
		return content, nil
	case 5:
		return nil, nil
	default:
		return str, nil
	}
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
