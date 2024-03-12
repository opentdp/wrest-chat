package aichat

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

func OpenaiText(id, rid, ask string) (string, error) {

	llmc := UserModel(id, rid)

	// 初始化模型

	config := openai.DefaultConfig(llmc.Secret)

	if len(llmc.Endpoint) > 1 {
		config.BaseURL = llmc.Endpoint + "/v1"
	}

	client := openai.NewClientWithConfig(config)

	req := openai.ChatCompletionRequest{
		Model:     "gpt-3.5-turbo",
		MaxTokens: 2048,
		Messages:  []openai.ChatCompletionMessage{},
	}

	if len(llmc.Model) > 1 {
		req.Model = llmc.Model
	}

	// 设置上下文

	if llmc.RoleContext != "" {
		req.Messages = []openai.ChatCompletionMessage{
			{Content: llmc.RoleContext, Role: openai.ChatMessageRoleSystem},
		}
	}

	for _, msg := range msgHistories[id] {
		role := msg.Role
		req.Messages = append(req.Messages, openai.ChatCompletionMessage{
			Content: msg.Content, Role: role,
		})
	}

	req.Messages = append(req.Messages, openai.ChatCompletionMessage{
		Content: ask, Role: "user",
	})

	// 请求模型接口

	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", err
	}

	// 更新历史记录

	item1 := &MsgHistory{Content: ask, Role: "user"}
	item2 := &MsgHistory{Content: resp.Choices[0].Message.Content, Role: "assistant"}

	AppendHistory(id, item1, item2)

	return item2.Content, nil

}
