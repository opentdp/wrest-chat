package aichat

import (
	"errors"

	"github.com/opentdp/wechat-rest/library/google"
)

func GoogleText(id, rid, ask string) (string, error) {

	llmc := UserModel(id, rid)

	// 初始化模型

	client := google.NewClient(llmc.Secret)
	client.Model = llmc.Model

	if llmc.Endpoint != "" {
		client.ApiBaseUrl = llmc.Endpoint
	}

	req := &google.RequestBody{
		Contents: []*google.Content{},
	}

	// 设置上下文

	if llmc.RoleContext != "" {
		req.Contents = []*google.Content{
			{Parts: []*google.Part{{Text: llmc.RoleContext}}, Role: "user"},
			{Parts: []*google.Part{{Text: "OK"}}, Role: "model"},
		}
	}

	for _, msg := range msgHistories[id] {
		role := msg.Role
		req.Contents = append(req.Contents, &google.Content{
			Parts: []*google.Part{{Text: msg.Content}}, Role: role,
		})
	}

	req.Contents = append(req.Contents, &google.Content{
		Parts: []*google.Part{{Text: ask}}, Role: google.ChatMessageRoleUser,
	})

	// 请求模型接口

	resp, err := client.CreateChatCompletion(req)
	if err != nil {
		return "", err
	}

	if resp.Error != nil {
		return "", errors.New(resp.Error.Message)
	}

	if len(resp.Candidates) == 0 || resp.Candidates[0].Content == nil {
		if resp.PromptFeedback.BlockReason != "" {
			return "", errors.New("BlockReason:" + resp.PromptFeedback.BlockReason)
		}
		return "", errors.New("未得到预期的结果")
	}

	// 更新历史记录

	item1 := &MsgHistory{Content: ask, Role: "user"}
	item2 := &MsgHistory{Content: resp.Candidates[0].Content.Parts[0].Text, Role: "model"}

	AppendHistory(id, item1, item2)

	return item2.Content, nil

}

func GoogleImage(id, rid, ask, img, mime string) (string, error) {

	llmc := UserModel(id, rid)

	// 初始化模型

	client := google.NewClient(llmc.Secret)

	if llmc.Endpoint != "" {
		client.ApiBaseUrl = llmc.Endpoint
	}

	// 请求模型接口

	resp, err := client.CreateImageCompletion(ask, img, mime)
	if err != nil {
		return "", err
	}

	if resp.Error != nil {
		return "", errors.New(resp.Error.Message)
	}

	if len(resp.Candidates) == 0 || resp.Candidates[0].Content == nil {
		if resp.PromptFeedback.BlockReason != "" {
			return "", errors.New("BlockReason:" + resp.PromptFeedback.BlockReason)
		}
		return "", errors.New("未得到预期的结果")
	}

	// 返回结果

	return resp.Candidates[0].Content.Parts[0].Text, nil

}
