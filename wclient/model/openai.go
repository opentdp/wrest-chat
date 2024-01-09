package model

import (
	"context"

	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/wclient/cache"
	"github.com/sashabaranov/go-openai"
)

func OpenaiChat(uid, msg string) (string, error) {

	config := openai.DefaultConfig(args.LLM.OpenAiKey)
	if args.LLM.OpenAiUrl != "" {
		config.BaseURL = args.LLM.OpenAiUrl
	}

	client := openai.NewClientWithConfig(config)

	// 构造请求参数

	req := openai.ChatCompletionRequest{
		Model:    cache.Models[uid],
		Messages: []openai.ChatCompletionMessage{},
	}

	for _, msg := range cache.History[uid] {
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

	cache.History[uid] = append(cache.History[uid], item1, item2)

	return item2.Content, nil

}
