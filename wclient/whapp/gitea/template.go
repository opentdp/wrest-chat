package gitea

import (
	"github.com/opentdp/wrest-chat/wclient/whapp"
)

var (
	TemplateUnsupport = `
🔔 来自%s的消息
⚠️ 暂不支持该类型
🙈 我们正在努力支持更多类型，敬请期待！
`
	TemplatePush = whapp.NewTemplate("GITEA_PUSH", `🔔 来自Gitea的消息
👤 {{ .Pusher.FullName }}（{{ .Pusher.Email }}）
📌 向仓库 {{ .Repository.FullName }} 推送了{{ .TotalCommits }}次提交
📊 提交记录：{{ range $index, $val := .Commits }}
{{inc $index}}: {{ getShortMsg $val.Message }}(by {{ $val.Author.Name }}){{ end }}
`)
	TemplateCreateTag = whapp.NewTemplate("GITEA_CREATE_TAG", `🔖 新Tag
📦 {{ .Repository.FullName }}
🏷️ {{ .Ref }}
👤 {{ .Sender.FullName }}（{{ .Sender.Email }}）
`)
	TemplateOpenIssue = whapp.NewTemplate("GITEA_CREATE_TAG", `✨ 有人提Issue了
📦 {{ .Repository.FullName }}#{{ .Issue.Number }}
💡 {{ .Issue.Title }}
👤 {{ .Sender.FullName }}（{{ .Sender.Email }}）
🏷️ {{ range _, $val := .Issue.Labels }}{{ $val.Name }} {{ end }} 
`)

	TemplateCreateIssueComment = whapp.NewTemplate("GITEA_CREATE_TAG", `🗨️ {{ .Repository.Name }}#{{ .Issue.Number }} 有新评论
📦 {{ .Repository.FullName }}
🏷️ {{ .Ref }}
👤 {{ .Sender.FullName }}（{{ .Sender.Email }}）
`)
)
