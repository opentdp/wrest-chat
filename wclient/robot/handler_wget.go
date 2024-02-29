package robot

import (
	"net/url"
	"strings"

	"github.com/opentdp/wechat-rest/wcferry"
)

func wgetHandler() {

	handlers["/wget"] = &Handler{
		Level:    7,
		Order:    90,
		ChatAble: true,
		RoomAble: true,
		Describe: "获取图片或文件",
		Callback: func(msg *wcferry.WxMsg) string {
			if fileReply(msg, msg.Content) == 0 {
				return ""
			}
			return "文件下载失败"
		},
	}

}

func fileReply(msg *wcferry.WxMsg, text string) int32 {

	if text = strings.TrimSpace(text); text == "" {
		return -1
	}

	if u, err := url.Parse(text); err == nil {
		if u.Scheme == "http" || u.Scheme == "https" {
			if isImageFile(u.Path) {
				if msg.IsGroup {
					return wc.CmdClient.SendImg(text, msg.Roomid)
				} else {
					return wc.CmdClient.SendImg(text, msg.Sender)
				}
			} else {
				if msg.IsGroup {
					return wc.CmdClient.SendFile(text, msg.Roomid)
				} else {
					return wc.CmdClient.SendFile(text, msg.Sender)
				}
			}
		}
		return -1
	}

	return -2

}

func isImageFile(text string) bool {

	imageExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".bmp":  true,
		".webp": true,
		".tiff": true,
		".svg":  true,
	}

	return imageExtensions[strings.ToLower(text)]

}
