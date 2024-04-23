package handlers

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"github.com/opentdp/wrest-chat/wclient/whapp/gitea/events"
	"github.com/opentdp/wrest-chat/wclient/whapp/gitea/templates"
)

func IssueCommentEventHandler(msg string) (string, error) {
	data := &events.GiteaIssueCommentEvent{}

	err := jsoniter.UnmarshalFromString(msg, &data)

	if err != nil {
		return "", errors.New("解析Gitea IssueComment事件失败")
	}

	switch data.Action {
	case "created":
		return templates.Render(templates.TemplateCreateIssueComment, data)
	}

	return "", errors.New("解析Gitea IssueComment事件失败")
}
