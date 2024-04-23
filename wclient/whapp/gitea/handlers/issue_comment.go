package handlers

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"github.com/opentdp/wrest-chat/wclient/whapp"
	"github.com/opentdp/wrest-chat/wclient/whapp/gitea"
)

func IssueCommentEventHandler(msg string) (string, error) {
	data := &gitea.GiteaIssueCommentEvent{}

	err := jsoniter.UnmarshalFromString(msg, &data)

	if err != nil {
		return "", errors.New("解析Gitea IssueComment事件失败")
	}

	switch data.Action {
	case "created":
		return whapp.Render(gitea.TemplateCreateIssueComment, data)
	}

	return "", errors.New("解析Gitea IssueComment事件失败")
}
