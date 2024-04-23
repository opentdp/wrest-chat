package handlers

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"github.com/opentdp/wrest-chat/wclient/whapp"
	"github.com/opentdp/wrest-chat/wclient/whapp/gitea"
)

func IssuesEventHandler(msg string) (string, error) {
	data := &gitea.GiteaIssuesEvent{}

	err := jsoniter.UnmarshalFromString(msg, &data)
	if err != nil {
		return "", errors.New("解析 Gitea Issues 事件失败")
	}

	switch data.Action {
	case "opened":
		return whapp.Render(gitea.TemplateOpenIssue, data)
	}

	return "", errors.New("解析 Gitea Issues 事件失败")
}
