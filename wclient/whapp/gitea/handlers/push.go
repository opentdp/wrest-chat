package handlers

import (
	"encoding/json"
	"errors"
	"github.com/opentdp/wrest-chat/wclient/whapp"
	"github.com/opentdp/wrest-chat/wclient/whapp/gitea"
)

func PushEventHandler(msg string) (string, error) {
	data := &gitea.GiteaPushEvent{}
	err := json.Unmarshal([]byte(msg), &data)

	if err != nil {
		return "", errors.New("解析Gitea Push事件失败")
	}

	return whapp.Render(gitea.TemplatePush, data)
}
