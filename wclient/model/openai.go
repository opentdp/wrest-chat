package model

import (
	"context"

	"github.com/opentdp/wechat-rest/args"
	"github.com/sashabaranov/go-openai"
)

func OpenaiChat(id, msg string) (string, error) {

	llmc := GetUserConfig(id).LLModel

	config := openai.DefaultConfig(llmc.Key)
	if llmc.Endpoint != "" {
		config.BaseURL = llmc.Endpoint + "/v1"
	}

	client := openai.NewClientWithConfig(config)

	// 初始化对话模型

	req := openai.ChatCompletionRequest{
		Model:    llmc.Model,
		Messages: []openai.ChatCompletionMessage{},
	}

	// 设置对话上下文

	if args.LLM.RoleContext != "" {
		req.Messages = []openai.ChatCompletionMessage{
			{Content: args.LLM.RoleContext, Role: "user"},
			{Content: "OK", Role: "assistant"},
		}
	}

	for _, msg := range msgHistoryMap[id] {
		role := msg.Role
		if role == "model" {
			role = "assistant"
		}
		req.Messages = append(req.Messages, openai.ChatCompletionMessage{
			Content: msg.Content, Role: role,
		})
	}

	req.Messages = append(req.Messages, openai.ChatCompletionMessage{
		Content: msg, Role: "user",
	})

	// 请求模型接口

	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", err
	}

	// 更新历史记录

	item1 := &MsgHistory{Content: msg, Role: "user"}
	item2 := &MsgHistory{Content: resp.Choices[0].Message.Content, Role: "model"}

	AppendHistory(id, item1, item2)

	return item2.Content, nil

}
