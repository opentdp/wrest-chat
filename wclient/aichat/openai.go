package aichat

import (
	"context"

	"github.com/opentdp/wechat-rest/dbase/setting"
	"github.com/sashabaranov/go-openai"
)

func OpenaiText(id, rid, ask string) (string, error) {

	llmc := UserModel(id, rid)

	config := openai.DefaultConfig(llmc.Secret)
	if llmc.Endpoint != "" {
		config.BaseURL = llmc.Endpoint + "/v1"
	}

	client := openai.NewClientWithConfig(config)

	// 初始化模型

	req := openai.ChatCompletionRequest{
		Model:     llmc.Model,
		MaxTokens: 2048,
		Messages:  []openai.ChatCompletionMessage{},
	}

	// 设置上下文

	if setting.ModelContext != "" {
		req.Messages = []openai.ChatCompletionMessage{
			{Content: setting.ModelContext, Role: openai.ChatMessageRoleUser},
			{Content: "OK", Role: openai.ChatMessageRoleAssistant},
		}
	}

	for _, msg := range msgHistories[id] {
		role := msg.Role
		if role == "user" {
			role = openai.ChatMessageRoleUser
		}
		if role == "model" {
			role = openai.ChatMessageRoleAssistant
		}
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
	item2 := &MsgHistory{Content: resp.Choices[0].Message.Content, Role: "model"}

	AppendHistory(id, item1, item2)

	return item2.Content, nil

}
