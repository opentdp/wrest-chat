package whapp

import (
	"bytes"
	"github.com/opentdp/wrest-chat/wclient/whapp/gitea"
	"github.com/opentdp/wrest-chat/wclient/whapp/github"
	"github.com/opentdp/wrest-chat/wclient/whapp/text"
	"net/http"
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"inc": func(i int) int {
		return i + 1
	},
	"getShortMsg": func(msg string) string {
		msgs := strings.Split(msg, "\n")
		if len(msgs) <= 1 {
			return strings.ReplaceAll(msg, "\n", "")
		}
		return strings.ReplaceAll(msgs[0], "\n", "")
	},
}

func Handler(header http.Header, app string, msg string) string {

	var res string
	var err error

	switch app {
	case "github":
		res, err = github.HandleWebhook(header, msg)
	case "gitea":
		res, err = gitea.HandleWebhook(header, msg)
	case "text":
		res, err = text.HandleWebhook(msg)
	default:
		res = "暂不支持该应用的 webhook"
	}

	if err != nil {
		return err.Error()
	}

	return res
}

func NewTemplate(name string, content string) *template.Template {
	return template.Must(template.New(name).Funcs(funcMap).Parse(content))
}

func Render(t *template.Template, data interface{}) (string, error) {
	var buf bytes.Buffer
	err := t.Execute(&buf, data)
	if err != nil {
		return "渲染通知模板失败", err
	}
	return buf.String(), nil
}
