package handlers

import (
	"encoding/json"
	"errors"
	"github.com/opentdp/wrest-chat/wclient/whapp/gitea/events"
	"github.com/opentdp/wrest-chat/wclient/whapp/gitea/templates"
)

func PushEventHandler(msg string) (string, error) {
	data := &events.GiteaPushEvent{}
	err := json.Unmarshal([]byte(msg), &data)

	if err != nil {
		return "", errors.New("解析Gitea Push事件失败")
	}

	return templates.Render(templates.TemplatePush, data)
}
