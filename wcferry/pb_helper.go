package wcferry

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/clbanning/mxj"
	"github.com/opentdp/go-helper/request"
	"github.com/opentdp/go-helper/strutil"
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

// 解析数据库字段
// param field *DbField 字段
// return any 解析结果
func ParseDbField(field *DbField) any {
	str := string(field.Content)
	switch field.Type {
	case 1:
		n, _ := strconv.ParseInt(str, 10, 64)
		return n
	case 2:
		n, _ := strconv.ParseFloat(str, 64)
		return n
	case 4:
		return field.Content
	case 5:
		return nil
	default:
		return str
	}
}

// 联系人类型
// param wxid string 联系人wxid
// return string 类型
func ContactType(wxid string) string {
	notFriends := map[string]string{
		"mphelper":    "公众平台助手",
		"fmessage":    "朋友推荐消息",
		"medianote":   "语音记事本",
		"floatbottle": "漂流瓶",
		"filehelper":  "文件传输助手",
		"newsapp":     "新闻",
	}
	if notFriends[wxid] != "" {
		return notFriends[wxid]
	}
	if strings.HasPrefix(wxid, "gh_") {
		return "公众号"
	}
	if strings.HasSuffix(wxid, "@chatroom") {
		return "群聊"
	}
	if strings.HasSuffix(wxid, "@openim") {
		return "企业微信"
	}
	return "好友"
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

// 获取网络文件
// param str string 文件URL或路径
// return string 失败则返回空字符串
func DownloadFile(str string) string {
	if strings.HasPrefix(str, "http://") || strings.HasPrefix(str, "https://") {
		if tmp, err := request.Download(str, "", false); err == nil {
			time.AfterFunc(15*time.Minute, func() {
				os.RemoveAll(tmp)
			})
			return tmp
		}
	}
	return ""
}
