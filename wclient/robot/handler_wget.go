package robot

import (
	"net/url"
	"strings"

	"github.com/opentdp/wechat-rest/wcferry"
)

func wgetHandler() {

	handlers["/wget"] = &Handler{
		Level:    7,
		ChatAble: true,
		RoomAble: true,
		Describe: "获取图片或文件",
		Callback: func(msg *wcferry.WxMsg) string {
			str := strings.TrimSpace(msg.Content)
			u, err := url.Parse(str)
			if err == nil && u.Scheme == "http" || u.Scheme == "https" {
				var status int32
				if isImageFile(u.Path) {
					if msg.IsGroup {
						status = wc.CmdClient.SendImg(str, msg.Roomid)
					} else {
						status = wc.CmdClient.SendImg(str, msg.Sender)
					}
				} else {
					if msg.IsGroup {
						status = wc.CmdClient.SendFile(str, msg.Roomid)
					} else {
						status = wc.CmdClient.SendFile(str, msg.Sender)
					}
				}
				if status == 0 {
					return ""
				}
				return "文件发送失败"
			}
			return "不支持文件协议"
		},
	}

}

func isImageFile(f string) bool {
	ext := strings.ToLower(f)
	imageExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp", ".tiff", ".svg"}
	for _, imageExt := range imageExtensions {
		if ext == imageExt {
			return true
		}
	}
	return false
}
