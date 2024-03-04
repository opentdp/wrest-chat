package aichat

import (
	"context"
	"errors"
	"strings"

	"github.com/liudding/go-llm-api/baidu"
)

func BaiDuText(id, rid, ask string) (string, error) {

	llmc := UserModel(id, rid)

	keys := strings.Split(llmc.Secret, ",")
	if len(keys) != 2 {
		return "", errors.New("密钥格式错误")
	}

	client := baidu.NewClient(keys[0], keys[1], true)

	// 初始化模型

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

	for _, msg := range msgHistories[id] {
		role := msg.Role
		if role == "user" {
			role = baidu.ChatMessageRoleUser
		}
		if role == "model" {
			role = baidu.ChatMessageRoleAssistant
		}
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

	reply := res.Result

	if reply == "" {
		return "", errors.New("未得到预期的结果")
	}

	// 更新历史记录

	item1 := &MsgHistory{Content: ask, Role: baidu.ChatMessageRoleUser}
	item2 := &MsgHistory{Content: reply, Role: baidu.ChatMessageRoleAssistant}

	AppendHistory(id, item1, item2)

	return item2.Content, nil

}
