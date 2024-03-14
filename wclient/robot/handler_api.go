package robot

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/opentdp/go-helper/request"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/dbase/setting"
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wclient"
)

func apiHandler() []*Handler {

	cmds := []*Handler{}

	if len(setting.ApiEndpoint) < 10 {
		return cmds
	}

	cmds = append(cmds, &Handler{
		Level:    0,
		Order:    200,
		Roomid:   "*",
		Command:  "/api",
		Describe: "调用查询接口",
		Callback: apiCallback,
	})

	return cmds

}

type apiCallbackData struct {
	Type string   `json:"type"` // 数据类型 [card, file, image, text, error]
	Card struct { // 当 type 为 card 时有效
		Name    string `json:"name"`    // 左下显示的名字，可选
		Account string `json:"account"` // 公众号 id，可显示对应的头像，可选
		Title   string `json:"title"`   // 标题，最多显示为两行
		Digest  string `json:"digest"`  // 摘要，最多显示为三行
		Link    string `json:"link"`    // 点击后跳转的链接
		Icon    string `json:"icon"`    // 右侧缩略图的链接，可选
	} `json:"card,omitempty"`
	Link string `json:"file,omitempty"` // 当 type 为 file 或 image 时有效
	Text string `json:"text,omitempty"` // 当 type 为 text 或 error 时有效
}

func apiCallback(msg *wcferry.WxMsg) string {

	self := wc.CmdClient.GetSelfInfo()

	// 组装参数
	cmd := []string{"help"}
	if msg.Content != "" {
		cmd = strings.SplitN(msg.Content, " ", 2)
		if len(cmd) > 1 {
			cmd[1] = url.QueryEscape(cmd[1])
		}
	}

	// 获取结果
	url := setting.ApiEndpoint + strings.Join(cmd, "/")
	res, err := request.TextGet(url, request.H{
		"Client-Uid": self.Wxid + "," + msg.Sender,
		"User-Agent": args.AppName + "/" + args.Version,
	})
	if err != nil {
		return err.Error()
	}

	// 解析失败
	data := &apiCallbackData{}
	if json.Unmarshal([]byte(res), data) != nil {
		return res
	}

	// 构造消息
	text := ""
	switch data.Type {
	case "card":
		if data.Card.Name == "" {
			data.Card.Name = self.Name
		}
		if data.Card.Account == "" {
			data.Card.Account = self.Wxid
		}
		if data.Card.Icon == "" {
			data.Card.Icon = setting.ApiEndpointIcon
		}
		text += "card\n"
		text += data.Card.Name + "\n"
		text += data.Card.Account + "\n"
		text += data.Card.Title + "\n"
		text += data.Card.Digest + "\n"
		text += data.Card.Link + "\n"
		text += data.Card.Icon
	case "file":
		text += data.Link
	case "image":
		text += data.Link
	default:
		text += data.Text
	}

	// 发送消息
	if wclient.SendFlexMsg(text, msg.Sender, msg.Roomid) == 0 {
		return ""
	}

	return text

}
