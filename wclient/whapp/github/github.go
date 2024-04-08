package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	name = "Github"
)

func HandleWebhook(header http.Header, msg string) (string, error) {

	hookType := header.Get("X-GitHub-Event")

	switch hookType {
	case "push":
		return githubPushEventHandler(msg)
	case "ping":
		return githubPingEventHandler(msg)
	}

	return fmt.Sprintf("收到来自 Github 的 Webhook，暂不支持的类型： %s", hookType), nil

}

func githubPingEventHandler(msg string) (string, error) {

	data := &GithubPingEvent{}
	err := json.Unmarshal([]byte(msg), &data)

	if err != nil {
		return "", errors.New("解析Github Ping事件数据失败")
	}

	return fmt.Sprintf(TemplatePing,
		name,
		data.Repository.FullName,
		data.Repository.HtmlUrl,
	), nil

}

func githubPushEventHandler(msg string) (string, error) {

	data := &GithubPushEvent{}
	err := json.Unmarshal([]byte(msg), &data)

	if err != nil {
		return "", errors.New("解析Github Push事件数据失败")
	}

	return fmt.Sprintf(TemplatePush,
		name,
		data.Pusher.Name, data.Pusher.Email,
		data.Repository.Name, len(data.Commits),
		data.Compare,
	), nil

}
