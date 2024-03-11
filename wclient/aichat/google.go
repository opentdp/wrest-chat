package aichat

import (
	"errors"
	"regexp"

	"github.com/opentdp/go-helper/googai"
)

func GoogleText(id, rid, ask string) (string, error) {

	llmc := UserModel(id, rid)

	// 初始化模型

	client := googai.NewClient(llmc.Secret)

	if len(llmc.Model) > 1 {
		client.Model = llmc.Model
	}

	if len(llmc.Endpoint) > 1 {
		client.ApiBaseUrl = llmc.Endpoint
	}

	req := &googai.RequestBody{
		Contents: []*googai.Content{},
	}

	// 设置上下文

	if llmc.RoleContext != "" {
		req.Contents = []*googai.Content{
			{Parts: []*googai.Part{{Text: llmc.RoleContext}}, Role: "user"},
			{Parts: []*googai.Part{{Text: "OK"}}, Role: "model"},
		}
	}

	for _, msg := range msgHistories[id] {
		role := msg.Role
		req.Contents = append(req.Contents, &googai.Content{
			Parts: []*googai.Part{{Text: msg.Content}}, Role: role,
		})
	}

	req.Contents = append(req.Contents, &googai.Content{
		Parts: []*googai.Part{{Text: ask}}, Role: googai.ChatMessageRoleUser,
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

func GoogleImage(id, rid, ask, img string) (string, error) {

	img, mime := ReadImage(img)

	imageTypeRegex := regexp.MustCompile(`^image/(png|jpeg|webp|heic|heif)$`)
	if !imageTypeRegex.MatchString(mime) {
		return "", errors.New("不支持此图片格式")
	}

	// 获取参数

	llmc := UserModel(id, rid)

	// 初始化模型

	client := googai.NewClient(llmc.Secret)

	if llmc.Endpoint != "" {
		client.ApiBaseUrl = llmc.Endpoint
	}

	req := &googai.RequestBody{
		Contents: []*googai.Content{
			{
				Parts: []*googai.Part{
					{Text: ask},
					{InlineData: &googai.InlineData{Data: img, MimeType: mime}},
				},
			},
		},
	}

	// 请求模型接口

	resp, err := client.CreateImageCompletion(req)
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
