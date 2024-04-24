package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/opentdp/wrest-chat/wclient/whapp/gitea/events"
	"github.com/opentdp/wrest-chat/wclient/whapp/gitea/templates"
)

func CreateEventHandler(msg string) (string, error) {
	data := &events.GiteaCreateEvent{}
	err := json.Unmarshal([]byte(msg), &data)

	if err != nil {
		return "", errors.New("解析Gitea Create事件失败")
	}

	switch data.RefType {
	case "tag":
		return templates.Render(templates.TemplateCreateTag, data)
	}

	return fmt.Sprintf("暂时不支持的 Create 事件 Ref 类型： %s", data.Ref), nil
}
