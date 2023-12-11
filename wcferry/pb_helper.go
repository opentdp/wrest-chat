package wcferry

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/opentdp/go-helper/request"
)

// 打印接收到的消息
// param msg *WxMsg 消息
func MsgPrinter(msg *WxMsg) {
	re := regexp.MustCompile(`(?m)^\s*|\n`)
	fmt.Print("\n=== New Message ===\n")
	if msg.Id > 0 {
		fmt.Printf("::Id:: %d\n", msg.Id)
	}
	if msg.Type > 0 {
		fmt.Printf("::Type:: %d\n", msg.Type)
	}
	if msg.Roomid != "" {
		fmt.Printf("::Roomid:: %s\n", msg.Roomid)
	}
	if msg.Sender != "" {
		fmt.Printf("::Sender:: %v\n", msg.Sender)
	}
	if msg.Content != "" {
		fmt.Printf("::Content:: %s\n", re.ReplaceAllString(msg.Content, ""))
	}
	if msg.Extra != "" {
		fmt.Printf("::Extra:: %s\n", re.ReplaceAllString(msg.Extra, ""))
	}
	fmt.Print("=== End Message ===\n")
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
