package wcferry

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/opentdp/go-helper/request"
)

// 打印接收到的消息
// param msg *MsgPayload 消息
func MsgPrinter(msg *MsgPayload) {
	rs := "\n=== New Message ===\n"
	re := regexp.MustCompile(`(?m)^\s*|\n`)
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
		rs += fmt.Sprintf("::Content:: %s\n", re.ReplaceAllString(msg.Content, ""))
	}
	if msg.Extra != "" {
		rs += fmt.Sprintf("::Extra:: %s\n", re.ReplaceAllString(msg.Extra, ""))
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
