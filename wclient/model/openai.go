package model

import (
	"context"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/wclient/cache"
	"github.com/sashabaranov/go-openai"
)

func OpenaiChat(id, msg string) (string, error) {

	config := openai.DefaultConfig(args.LLM.OpenAiKey)
	if args.LLM.OpenAiUrl != "" {
		config.BaseURL = args.LLM.OpenAiUrl
	}

	client := openai.NewClientWithConfig(config)

	// 构造请求参数

	req := openai.ChatCompletionRequest{
		Model:    cache.Models[id],
		Messages: []openai.ChatCompletionMessage{},
	}

	for _, msg := range cache.History[id] {
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

	item1 := cache.HistoryItem{
		Role:    "user",
		Content: msg,
	}

	item2 := cache.HistoryItem{
		Role:    resp.Choices[0].Message.Role,
		Content: resp.Choices[0].Message.Content,
	}

	cache.History[id] = append(cache.History[id], item1, item2)

	return item2.Content, nil

}
