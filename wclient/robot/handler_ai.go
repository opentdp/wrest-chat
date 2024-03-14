package robot

import (
	"fmt"
	"math/rand"

	"github.com/opentdp/wechat-rest/dbase/llmodel"
	"github.com/opentdp/wechat-rest/dbase/message"
	"github.com/opentdp/wechat-rest/dbase/profile"
	"github.com/opentdp/wechat-rest/wcferry"
	"github.com/opentdp/wechat-rest/wclient/aichat"
)

func aiHandler() []*Handler {

	cmds := []*Handler{}

	models, err := llmodel.FetchAll(&llmodel.FetchAllParam{})
	if err != nil || len(models) == 0 {
		return cmds
	}

	cmds = append(cmds, &Handler{
		Level:    0,
		Order:    100,
		Roomid:   "*",
		Command:  "/ai",
		Describe: "提问或交谈",
		Callback: aiCallback,
	})

	cmds = append(cmds, &Handler{
		Level:    0,
		Order:    101,
		Roomid:   "*",
		Command:  "/ai:new",
		Describe: "重置上下文内容",
		Callback: func(msg *wcferry.WxMsg) string {
			aichat.ResetHistory(msg.Sender, msg.Roomid)
			return "已重置上下文"
		},
	})

	if len(models) > 3 {
		cmds = append(cmds, &Handler{
			Level:    0,
			Order:    103,
			Roomid:   "*",
			Command:  "/ai:rand",
			Describe: "随机选择模型",
			Callback: func(msg *wcferry.WxMsg) string {
				up, _ := profile.Fetch(&profile.FetchParam{Wxid: msg.Sender, Roomid: prid(msg)})
				ks := []int{}
				for k, v := range models {
					if v.Level <= up.Level {
						ks = append(ks, k)
					}
				}
				if len(ks) > 0 {
					v := models[ks[rand.Intn(len(ks))]]
					profile.Replace(&profile.ReplaceParam{Wxid: msg.Sender, Roomid: prid(msg), AiModel: v.Mid})
					return "对话模型切换为 " + v.Family + " [" + v.Model + "]"
				}
				return fmt.Sprintf("没有可用的模型（Level ≤ %d）", up.Level)
			},
		})
	}

	for k, v := range models {
		v := v // copy
		cmdkey := "/cm:" + v.Mid
		cmds = append(cmds, &Handler{
			Level:    v.Level,
			Order:    110 + int32(k),
			Roomid:   "*",
			Command:  cmdkey,
			Describe: "换模型 " + v.Family,
			Callback: func(msg *wcferry.WxMsg) string {
				profile.Replace(&profile.ReplaceParam{Wxid: msg.Sender, Roomid: prid(msg), AiModel: v.Mid})
				return "对话模型切换为 " + v.Family + " [" + v.Model + "]"
			},
		})
	}

	return cmds

}

func aiCallback(msg *wcferry.WxMsg) string {

	if msg.Content == "" {
		return "请在指令后输入问题"
	}

	// 处理引用的消息
	if msg.Sign == "refer-msg" {
		ref, err := message.Fetch(&message.FetchParam{Id: msg.Id})
		if err != nil { //TODO: 此处无法提取机器人发的消息
			ref.Content = msg.Extra
		}
		switch msg.Type {
		// 文本
		case 1:
			if ref.Content != "" {
				msg.Content += "\n内容如下:\n" + ref.Content
				return aichat.Text(msg.Sender, msg.Roomid, msg.Content)
			}
		// 图片
		case 3:
			if ref.Remark == "" {
				ref.Remark = msgImage(ref.Id, ref.Extra)
				if ref.Remark == "" {
					return "提取消息图片失败"
				}
			}
			return aichat.Vison(msg.Sender, msg.Roomid, msg.Content, ref.Remark)
		// 混合类消息
		case 49:
			if ref.Content != "" {
				msg.Content += "\nXML数据如下:\n" + ref.Content
				return aichat.Text(msg.Sender, msg.Roomid, msg.Content)
			}
		// 默认提示
		default:
			return "暂不支持处理此类消息"
		}
	}

	return aichat.Text(msg.Sender, msg.Roomid, msg.Content)

}
