package webhookApp

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/opentdp/wrest-chat/wclient/webhookApp/event"
	"net/http"
)

func GithubWebhook(header http.Header, msg string) (string, error) {
	hookType := header.Get("X-GitHub-Event")
	targetType := header.Get("X-GitHub-Hook-Installation-Target-Type")

	switch hookType {
	case "push":
		return githubPushEventHandler(targetType, msg)
	case "ping":
		return githubPingEventHandler(targetType, msg)
	}

	return fmt.Sprintf("收到来自 Github 的 Webhook，暂不支持的类型： %s", hookType), nil
}

func githubPingEventHandler(targetType string, msg string) (string, error) {
	data := &event.GithubPingEvent{}
	err := json.Unmarshal([]byte(msg), &data)

	if err != nil {
		return "", errors.New("解析Github Ping事件数据失败")
	}

	target := getTarget(targetType)

	return fmt.Sprintf("[Github] 收到来自 %s %s 的测试请求", target, data.Repository.Name), nil
}

func githubPushEventHandler(targetType string, msg string) (string, error) {
	data := &event.GithubPushEvent{}
	err := json.Unmarshal([]byte(msg), &data)

	if err != nil {
		return "", errors.New("解析Github Push事件数据失败")
	}

	target := getTarget(targetType)

	return fmt.Sprintf("[Github] %s(%s) 向 %s %s 推送了%d次提交\n\n详情查看: %s", data.Pusher.Name, data.Pusher.Email, target, data.Repository.Name, len(data.Commits), data.Compare), nil
}

func getTarget(targetType string) string {
	switch targetType {
	case "repository":
		return "仓库"
	}

	return "未知"
}
