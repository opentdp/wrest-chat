package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/opentdp/wrest-chat/wclient/whapp"
	"github.com/opentdp/wrest-chat/wclient/whapp/gitea"
)

func CreateEventHandler(msg string) (string, error) {
	data := &gitea.GiteaCreateEvent{}
	err := json.Unmarshal([]byte(msg), &data)

	if err != nil {
		return "", errors.New("解析Gitea Create事件失败")
	}

	switch data.Ref {
	case "tag":
		return whapp.Render(gitea.TemplateCreateTag, data)
	}

	return fmt.Sprintf("暂时不支持的 Create 事件 Ref 类型： %s", data.Ref), nil
}
