package templates

import (
	"bytes"
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

var (
	TemplateUnsupport = `
🔔 来自%s的消息
⚠️ 暂不支持该类型
🙈 我们正在努力支持更多类型，敬请期待！
`
	TemplatePush = NewTemplate("GITEA_PUSH", `🔔 有人推送代码啦
📦 仓库：{{ .Repository.FullName }}
📊 提交记录：{{ range $index, $val := .Commits }}
    {{inc $index}}. {{ getShortMsg $val.Message }}(by @{{ $val.Author.Name }}){{ end }}
`)
	TemplateCreateTag = NewTemplate("GITEA_CREATE_TAG", `🔖 有人创建Tag啦
📦 {{ .Repository.FullName }}
🏷️ {{ .Ref }}
`)
	TemplateOpenIssue = NewTemplate("OPEN_ISSUE", `✨ 有人提Issue啦
📦 {{ .Repository.FullName }}#{{ .Issue.Number }}
💡 {{ .Issue.Title }}
👤 {{ .Sender.FullName }}（{{ .Sender.Email }}）
🏷️ {{ range $index, $val := .Issue.Labels }}{{ $val.Name }} {{ end }} 
`)

	TemplateCreateIssueComment = NewTemplate("CREATE_ISSUE_COMMENT", `🗨️ {{ .Repository.Name }}#{{ .Issue.Number }} 有新评论
📦 {{ .Repository.FullName }}#{{ .Issue.Number }}
👤 {{ .Sender.FullName }}（{{ .Sender.Email }}）
`)
)

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
