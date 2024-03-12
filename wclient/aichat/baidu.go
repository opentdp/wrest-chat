package aichat

import (
	"context"
	"errors"
	"strings"

	"github.com/rehiy/one-llm/baidu"
)

func BaiDuText(id, rid, ask string) (string, error) {

	llmc := UserModel(id, rid)

	keys := strings.Split(llmc.Secret, ",")
	if len(keys) != 2 {
		return "", errors.New("密钥格式错误")
	}

	// 初始化模型

	config := baidu.DefaultConfig(keys[0], keys[1], true)

	if len(llmc.Endpoint) > 1 {
		config.BaseURL = llmc.Endpoint
	}

	client := baidu.NewClientWithConfig(config)

	req := baidu.ChatCompletionRequest{
		Messages: []baidu.ChatCompletionMessage{},
	}

	// 设置上下文

	if llmc.RoleContext != "" {
		req.Messages = []baidu.ChatCompletionMessage{
			{Content: llmc.RoleContext, Role: baidu.ChatMessageRoleUser},
			{Content: "OK", Role: baidu.ChatMessageRoleAssistant},
		}
	}

	for _, msg := range GetHistory(id, rid) {
		role := msg.Role
		req.Messages = append(req.Messages, baidu.ChatCompletionMessage{
			Content: msg.Content, Role: role,
		})
	}

	req.Messages = append(req.Messages, baidu.ChatCompletionMessage{
		Content: ask, Role: baidu.ChatMessageRoleUser,
	})

	// 请求模型接口

	res, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", err
	}

	if res.Result == "" {
		return "", errors.New("未得到预期的结果")
	}

	// 更新历史记录

	item1 := &MsgHistory{Content: ask, Role: "user"}
	item2 := &MsgHistory{Content: res.Result, Role: "assistant"}

	AddHistory(id, rid, item1, item2)

	return item2.Content, nil

}
