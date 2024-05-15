package gitea

import (
	"fmt"
	"github.com/opentdp/wrest-chat/wclient/whapp/gitea/handlers"
	"github.com/opentdp/wrest-chat/wclient/whapp/gitea/templates"
	"net/http"
)

var (
	name = "Gitea"
)

func HandleWebhook(header http.Header, msg string) (string, error) {
	hookType := header.Get("X-Gitea-Event")
	switch hookType {
	case "push":
		return handlers.PushEventHandler(msg)
	case "create":
		return handlers.CreateEventHandler(msg)
	case "issue_comment":
		return handlers.IssueCommentEventHandler(msg)
	case "issues":
		return handlers.IssuesEventHandler(msg)
	}

	return fmt.Sprintf(templates.TemplateUnsupport, name), nil

}
