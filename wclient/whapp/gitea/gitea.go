package gitea

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	name = "Gitea"
)

func HandleWebhook(header http.Header, msg string) (string, error) {
	hookType := header.Get("X-Gitea-Event")
	switch hookType {
	case "push":
		return giteaPushEventHandler(msg)
	}

	return fmt.Sprintf(TemplateUnsupport, name), nil

}

func giteaPushEventHandler(msg string) (string, error) {
	data := &GiteaPushEvent{}
	err := json.Unmarshal([]byte(msg), &data)

	if err != nil {
		return "", errors.New("解析Gitea Push事件失败")
	}

	return fmt.Sprintf(TemplatePush,
		name,
		data.Pusher.FullName, data.Pusher.Email,
		data.Repository.FullName, data.TotalCommits,
		data.CompareUrl,
	), nil
}
