package wcferry

import (
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/opentdp/go-helper/request"
)

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

// 获取联系人类型
// param wxid string 联系人wxid
// return string 类型
func ContactType(wxid string) string {
	notFriends := map[string]string{
		"fmessage":    "朋友推荐消息",
		"filehelper":  "文件传输助手",
		"floatbottle": "漂流瓶",
		"medianote":   "语音记事本",
		"mphelper":    "公众平台助手",
		"newsapp":     "新闻",
	}
	if notFriends[wxid] != "" {
		return notFriends[wxid]
	}
	if strings.HasSuffix(wxid, "@chatroom") {
		return "群聊"
	}
	if strings.HasSuffix(wxid, "@openim") {
		return "企业微信"
	}
	if strings.HasPrefix(wxid, "gh_") {
		return "公众号"
	}
	return "好友"
}

// 获取网络文件
// param str string 文件URL或路径
// return string 失败则返回空字符串
func DownloadFile(str string) string {
	u, err := url.Parse(str)
	if err == nil && u.Scheme == "http" || u.Scheme == "https" {
		target := path.Join(os.TempDir(), strings.Trim(path.Base(u.Path), "/"))
		tmp, err := request.Download(str, target, false)
		if err == nil {
			time.AfterFunc(15*time.Minute, func() {
				os.RemoveAll(tmp)
			})
			return tmp
		}
	}
	return ""
}

// 根据扩展名推测是否图片
// param text string 文件URL或路径
// return bool 是否为图片
func IsImageFile(str string) bool {
	list := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".bmp":  true,
		".webp": true,
		".tiff": true,
		".svg":  true,
	}
	return list[strings.ToLower(str)]
}
