package aichat

import (
	"context"
	"errors"
	"io"
	"strings"

	"github.com/liudding/go-llm-api/xunfei"

	"github.com/opentdp/wechat-rest/args"
)

func XunfeiText(id, rid, ask string) (string, error) {

	llmc := UserModel(id, rid)

	keys := strings.Split(llmc.Secret, ",")
	if len(keys) != 3 {
		return "", errors.New("密钥格式错误")
	}

	client := xunfei.NewClient(keys[0], keys[1], keys[2])

	// 初始化模型

	req := xunfei.ChatCompletionRequest{
		MaxTokens: 2048,
		Messages:  []xunfei.ChatCompletionMessage{},
	}

	// 设置上下文

	if args.LLM.RoleContext != "" {
		req.Messages = []xunfei.ChatCompletionMessage{
			{Content: args.LLM.RoleContext, Role: xunfei.ChatMessageRoleUser},
			{Content: "OK", Role: xunfei.ChatMessageRoleAssistant},
		}
	}

	for _, msg := range msgHistories[id] {
		role := msg.Role
		if role == "user" {
			role = xunfei.ChatMessageRoleUser
		}
		if role == "model" {
			role = xunfei.ChatMessageRoleAssistant
		}
		req.Messages = append(req.Messages, xunfei.ChatCompletionMessage{
			Content: msg.Content, Role: role,
		})
	}

	req.Messages = append(req.Messages, xunfei.ChatCompletionMessage{
		Content: ask, Role: xunfei.ChatMessageRoleUser,
	})

	// 请求模型接口

	stream, err := client.CreateChatCompletionStream(context.Background(), req, llmc.Model)
	if err != nil {
		return "", err
	}

	defer stream.Close()

	reply := ""

	for {
		response, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return reply, err
		}
		if len(response.Payload.Choices.Text) > 0 {
			reply += response.Payload.Choices.Text[0].Content
		} else {
			reply += response.Payload.Choices.Content
		}
	}

	if reply == "" {
		return "", errors.New("未得到预期的结果")
	}

	// 更新历史记录

	item1 := &MsgHistory{Content: ask, Role: xunfei.ChatMessageRoleUser}
	item2 := &MsgHistory{Content: reply, Role: xunfei.ChatMessageRoleAssistant}

	AppendHistory(id, item1, item2)

	return item2.Content, nil

}
