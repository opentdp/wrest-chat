package model

import (
	"context"
	"errors"

	"github.com/sashabaranov/go-openai"
)

func OpenaiChat(id, msg string) (string, error) {

	llmc := Model(id)
	if llmc == nil {
		return "", errors.New("未配置模型")
	}

	config := openai.DefaultConfig(llmc.Key)
	if llmc.Endpoint != "" {
		config.BaseURL = llmc.Endpoint + "/v1"
	}

	client := openai.NewClientWithConfig(config)

	// 构造请求参数

	req := openai.ChatCompletionRequest{
		Model:    llmc.Model,
		Messages: []openai.ChatCompletionMessage{},
	}

	for _, msg := range History[id] {
		role := openai.ChatMessageRoleAssistant
		if msg.Role == "user" {
			role = openai.ChatMessageRoleUser
		}
		req.Messages = append(req.Messages, openai.ChatCompletionMessage{
			Role:    role,
			Content: msg.Content,
		})
	}

	req.Messages = append(req.Messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: msg,
	})

	// 请求模型接口

	resp, err := client.CreateChatCompletion(
		context.Background(), req,
	)
	if err != nil {
		return "", err
	}

	// 更新历史记录

	item1 := HistoryItem{
		Role:    "user",
		Content: msg,
	}

	item2 := HistoryItem{
		Role:    resp.Choices[0].Message.Role,
		Content: resp.Choices[0].Message.Content,
	}

	History[id] = append(History[id], item1, item2)

	return item2.Content, nil

}
