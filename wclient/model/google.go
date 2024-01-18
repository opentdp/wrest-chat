package model

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"github.com/opentdp/wechat-rest/args"
	"google.golang.org/api/option"
)

func GoogleChat(id, ask string) (string, error) {

	llmc := GetUserConfig(id).LLModel

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

	// 初始化模型

	model := client.GenerativeModel(llmc.Model)

	req := model.StartChat()
	req.History = []*genai.Content{}

	// 设置上下文

	if args.LLM.RoleContext != "" {
		req.History = []*genai.Content{
			{Parts: []genai.Part{genai.Text(args.LLM.RoleContext)}, Role: "user"},
			{Parts: []genai.Part{genai.Text("OK")}, Role: "model"},
		}
	}

	for _, msg := range msgHistoryMap[id] {
		role := msg.Role
		req.History = append(req.History, &genai.Content{
			Parts: []genai.Part{genai.Text(msg.Content)}, Role: role,
		})
	}

	// 请求模型接口

	resp, err := req.SendMessage(ctx, genai.Text(ask))
	if err != nil {
		return "", err
	}

	if len(resp.Candidates) == 0 || resp.Candidates[0].Content == nil {
		return "", errors.New("未得到预期的结果")
	}

	// 更新历史记录

	item1 := &MsgHistory{Content: ask, Role: "user"}
	item2 := &MsgHistory{Content: fmt.Sprintf("%s", resp.Candidates[0].Content.Parts[0]), Role: "model"}

	AppendHistory(id, item1, item2)

	return item2.Content, nil

}
