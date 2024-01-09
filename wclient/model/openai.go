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
		req.Messages = append(req.Messages, openai.ChatCompletionMessage{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}

	// 调用 OpenAI 接口

	resp, err := client.CreateChatCompletion(
		context.Background(), req,
	)

	if err != nil {
		return "", err
	}

	// 解析返回的数据

	data := cache.HistoryItem{
		Role:    resp.Choices[0].Message.Role,
		Content: resp.Choices[0].Message.Content,
	}
	cache.History[uid] = append(cache.History[uid], data)

	return data.Content, nil

}
