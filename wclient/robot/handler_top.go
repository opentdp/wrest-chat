package robot

import (
	"fmt"
	"strconv"
	"time"

	"github.com/opentdp/wrest-chat/wcferry"
	"github.com/opentdp/wrest-chat/wclient"
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

	ago := int64(day)

	// 计算日期
	if day > 2 {
		ts := wclient.TodayUnix() - ago*86400
		sub = time.Unix(ts, 0).Format("2006年1月2日")
	} else if day == 1 {
		sub = "昨日"
	}

	// 群聊概况
	count := wclient.RoomCount(msg.Roomid, ago)
	if count.Talk > 0 {
		res += fmt.Sprintf(sub+"本群共发言 %d 次，图片 %d 张", count.Talk, count.Image)
		res += "\n---------------------\n"
	}

	// 聊天统计
	if items := wclient.TalkTop10(msg.Roomid, ago); len(items) > 0 {
		res += ">>>> 🏊 水王 Top10 🏊 <<<<\n"
		for _, v := range items {
			u := wc.CmdClient.GetAliasInChatRoom(v.Sender, msg.Roomid)
			res += fmt.Sprintf("%s:   %d 次\n", u, v.RecordCount)
		}
	}

	// 图片统计
	if items := wclient.ImageTop10(msg.Roomid, ago); len(items) > 0 {
		res += ">>>> 🌅 图王 Top10 🌅 <<<<\n"
		for _, v := range items {
			u := wc.CmdClient.GetAliasInChatRoom(v.Sender, msg.Roomid)
			res += fmt.Sprintf("%s:   %d 次\n", u, v.RecordCount)
		}
	}

	return res

}
