package aichat

import (
	"errors"
	"regexp"

	"github.com/rehiy/one-llm/google"
)

func GoogleText(id, rid, ask string) (string, error) {

	llmc := UserModel(id, rid)

	// 初始化模型

	client := google.NewClient(llmc.Secret)

	if len(llmc.Model) > 1 {
		client.Model = llmc.Model
	}

	if len(llmc.Endpoint) > 1 {
		client.ApiBaseUrl = llmc.Endpoint
	}

	req := []google.ChatCompletionMessage{}

	// 设置上下文

	if llmc.RoleContext != "" {
		req = []google.ChatCompletionMessage{
			{Parts: []google.ChatCompletionMessagePart{{Text: llmc.RoleContext}}, Role: google.ChatMessageRoleUser},
			{Parts: []google.ChatCompletionMessagePart{{Text: "OK"}}, Role: google.ChatMessageRoleAssistant},
		}
	}

	for _, msg := range GetHistory(id, rid) {
		role := msg.Role
		if role == "assistant" {
			role = google.ChatMessageRoleAssistant
		}
		req = append(req, google.ChatCompletionMessage{
			Parts: []google.ChatCompletionMessagePart{{Text: msg.Content}}, Role: role,
		})
	}

	req = append(req, google.ChatCompletionMessage{
		Parts: []google.ChatCompletionMessagePart{{Text: ask}}, Role: google.ChatMessageRoleUser,
	})

	// 请求模型接口

	resp, err := client.CreateChatCompletion(req)
	if err != nil {
		return "", err
	}

	if resp.Error.Message != "" {
		return "", errors.New(resp.Error.Message)
	}

	if len(resp.Candidates) == 0 || resp.Candidates[0].Content.Role == "" {
		if resp.PromptFeedback.BlockReason != "" {
			return "", errors.New("BlockReason:" + resp.PromptFeedback.BlockReason)
		}
		return "", errors.New("未得到预期的结果")
	}

	// 更新历史记录

	item1 := &MsgHistory{Content: ask, Role: "user"}
	item2 := &MsgHistory{Content: resp.Candidates[0].Content.Parts[0].Text, Role: "assistant"}

	AddHistory(id, rid, item1, item2)

	return item2.Content, nil

}

func GoogleVison(id, rid, ask, img string) (string, error) {

	img, mime := ReadImage(img)

	imageTypeRegex := regexp.MustCompile(`^image/(png|jpeg|webp|heic|heif)$`)
	if !imageTypeRegex.MatchString(mime) {
		return "", errors.New("不支持此图片格式")
	}

	// 获取参数

	llmc := UserModel(id, rid)

	// 初始化模型

	client := google.NewClient(llmc.Secret)

	if llmc.Endpoint != "" {
		client.ApiBaseUrl = llmc.Endpoint
	}

	req := []google.ChatCompletionMessage{
		{
			Parts: []google.ChatCompletionMessagePart{
				{Text: ask},
				{InlineData: &google.ChatCompletionInlineData{Data: img, MimeType: mime}},
			},
		},
	}

	// 请求模型接口

	resp, err := client.CreateVisionCompletion(req)
	if err != nil {
		return "", err
	}

	if resp.Error.Message != "" {
		return "", errors.New(resp.Error.Message)
	}

	if len(resp.Candidates) == 0 || resp.Candidates[0].Content.Role == "" {
		if resp.PromptFeedback.BlockReason != "" {
			return "", errors.New("BlockReason:" + resp.PromptFeedback.BlockReason)
		}
		return "", errors.New("未得到预期的结果")
	}

	// 返回结果

	return resp.Candidates[0].Content.Parts[0].Text, nil

}
