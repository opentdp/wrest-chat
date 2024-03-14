package robot

import (
	"fmt"
	"strconv"
	"time"

	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wclient"
)

func topHandler() []*Handler {

	cmds := []*Handler{}

	cmds = append(cmds, &Handler{
		Level:    7,
		Order:    330,
		Roomid:   "+",
		Command:  "/top",
		Describe: "获取群聊统计信息",
		Callback: topCallback,
	})

	return cmds

}

func topCallback(msg *wcferry.WxMsg) string {

	res := ""
	sub := "今日"
	day, _ := strconv.Atoi(msg.Content)

	// 计算日期
	if day > 2 {
		ts := wclient.TodayUnix() - int64(day)*86400
		sub = time.Unix(ts, 0).Format("2006年1月2日")
	} else if day == 1 {
		sub = "昨日"
	}

	// 聊天统计
	if items := wclient.TalkTop10(msg.Roomid, int64(day)); len(items) > 0 {
		res += "\n🏊 " + sub + "水王\n----------------\n"
		for _, v := range items {
			u := wc.CmdClient.GetAliasInChatRoom(v.Sender, msg.Roomid)
			res += fmt.Sprintf("%s:   %d 次\n", u, v.RecordCount)
		}
	}

	// 图片统计
	if items := wclient.ImageTop10(msg.Roomid, int64(day)); len(items) > 0 {
		res += "\n🌅 " + sub + "图王\n----------------\n"
		for _, v := range items {
			u := wc.CmdClient.GetAliasInChatRoom(v.Sender, msg.Roomid)
			res += fmt.Sprintf("%s:   %d 次\n", u, v.RecordCount)
		}
	}

	return res

}
