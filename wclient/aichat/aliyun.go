package aichat

import (
	"errors"

	"github.com/opentdp/go-helper/aliqwen"
)

func AliyunText(id, rid, ask string) (string, error) {

	llmc := UserModel(id, rid)

	// 初始化模型

	client := aliqwen.NewClient(llmc.Secret)

	if len(llmc.Model) > 1 {
		client.Model = llmc.Model
	}

	if len(llmc.Endpoint) > 1 {
		client.ApiBaseUrl = llmc.Endpoint
	}

	req := []*aliqwen.Messages{}

	// 设置上下文

	if llmc.RoleContext != "" {
		req = []*aliqwen.Messages{
			{Content: llmc.RoleContext, Role: aliqwen.ChatMessageRoleSystem},
		}
	}

	for _, msg := range GetHistory(id, rid) {
		role := msg.Role
		req = append(req, &aliqwen.Messages{
			Content: msg.Content, Role: role,
		})
	}

	req = append(req, &aliqwen.Messages{
		Content: ask, Role: aliqwen.ChatMessageRoleUser,
	})

	// 请求模型接口

	resp, err := client.CreateChatCompletion(req)
	if err != nil {
		return "", err
	}

	if resp.Message != "" {
		return "", errors.New(resp.Message)
	}

	if resp.Output.Text == "" {
		return "", errors.New(resp.Output.FinishReason)
	}

	// 更新历史记录

	item1 := &MsgHistory{Content: ask, Role: "user"}
	item2 := &MsgHistory{Content: resp.Output.Text, Role: "assistant"}

	AddHistory(id, rid, item1, item2)

	return item2.Content, nil

}
