package handlers

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"github.com/opentdp/wrest-chat/wclient/whapp/gitea/events"
	"github.com/opentdp/wrest-chat/wclient/whapp/gitea/templates"
)

func IssuesEventHandler(msg string) (string, error) {
	data := &events.GiteaIssuesEvent{}

	err := jsoniter.UnmarshalFromString(msg, &data)
	if err != nil {
		return "", errors.New("解析 Gitea Issues 事件失败")
	}

	switch data.Action {
	case "opened":
		return templates.Render(templates.TemplateOpenIssue, data)
	}

	return "", errors.New("解析 Gitea Issues 事件失败")
}
