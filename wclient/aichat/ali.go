package aichat

import (
	"errors"
	client "github.com/aliyun/alibabacloud-bailian-go-sdk/client"
	"log"
	"strings"
)

func AliText(id, rid, ask string) (string, error) {

	llmc := UserModel(id, rid)

	keys := strings.Split(llmc.Secret, ",")
	if len(keys) != 4 {
		return "", errors.New("密钥格式错误")
	}

	appId := keys[0]
	agentKey := keys[1]
	accessKeyId := keys[2]
	accessKeySecret := keys[3]

	// 初始化模型
	aliClient := client.AccessTokenClient{AccessKeyId: accessKeyId, AccessKeySecret: accessKeySecret, AgentKey: agentKey}
	token, err := aliClient.GetToken()
	if err != nil {
		log.Printf("%v\n", err)
		return "", errors.New("[通义千问]初始化Token失败")
	}

	cc := client.CompletionClient{Token: token}

	if len(llmc.Endpoint) > 1 {
		cc.Endpoint = llmc.Endpoint
	}

	req := &client.CompletionRequest{
		AppId:      appId,
		Messages:   []client.ChatCompletionMessage{},
		Parameters: &client.CompletionRequestModelParameter{ResultFormat: "message"},
	}

	// 设置上下文

	if llmc.RoleContext != "" {
		req.Messages = []client.ChatCompletionMessage{
			{Content: llmc.RoleContext, Role: "user"},
			{Content: "OK", Role: "assistant"},
		}
	}

	for _, msg := range msgHistories[id] {
		role := msg.Role
		if role == "user" {
			role = "user"
		}
		if role == "model" {
			role = "assistant"
		}
		req.Messages = append(req.Messages, client.ChatCompletionMessage{
			Content: msg.Content, Role: role,
		})
	}

	req.Messages = append(req.Messages, client.ChatCompletionMessage{
		Content: ask, Role: "user",
	})

	// 请求模型接口

	res, err := cc.CreateCompletion(req)
	if err != nil {
		return "", err
	}

	if !res.Success {
		return "", errors.New("未得到预期的结果")
	}

	if len(res.Data.Choices) <= 0 {
		return "", errors.New("未得到预期的结果")
	}

	reply := ""

	reply += res.Data.Choices[0].Message.Content

	// 更新历史记录
	item1 := &MsgHistory{Content: ask, Role: "user"}
	item2 := &MsgHistory{Content: reply, Role: "model"}

	AppendHistory(id, item1, item2)

	return item2.Content, nil

}
