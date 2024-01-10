package model

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GoogleChat(id, msg string) (string, error) {

	llmc := GetUserModel(id)

	opts := []option.ClientOption{
		option.WithAPIKey(llmc.Key),
	}
	if llmc.Endpoint != "" {
		opts = append(opts, option.WithEndpoint(llmc.Endpoint))
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, opts...)
	if err != nil {
		return "", err
	}

	defer client.Close()

	// 构造请求参数

	model := client.GenerativeModel(llmc.Model)

	cs := model.StartChat()
	cs.History = []*genai.Content{}

	for _, msg := range MsgHistory[id] {
		role := "model"
		if msg.Role == "user" {
			role = "user"
		}
		cs.History = append(cs.History, &genai.Content{
			Parts: []genai.Part{genai.Text(msg.Content)},
			Role:  role,
		})
	}

	// 请求模型接口

	resp, err := cs.SendMessage(ctx, genai.Text(msg))
	if err != nil {
		return "", err
	}

	if len(resp.Candidates) == 0 || resp.Candidates[0].Content == nil {
		return "", fmt.Errorf("未得到预期的结果")
	}

	// 更新历史记录

	item1 := &HistoryItem{
		Content: msg,
		Role:    "user",
	}

	item2 := &HistoryItem{
		Content: fmt.Sprintf("%s", resp.Candidates[0].Content.Parts[0]),
		Role:    "model",
	}

	AppendHistory(id, item1, item2)

	return item2.Content, nil

}
