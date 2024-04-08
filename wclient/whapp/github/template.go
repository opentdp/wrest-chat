package github

var (
	TemplateUnsupport = `
🔔 来自%s的消息
⚠️ 暂不支持该类型
🙈 我们正在努力支持更多类型，敬请期待！
`
	TemplatePush = `
🔔 来自%s的消息
👤 %s（%s）
📌 向仓库 %s 推送了%d次提交
🔗 详情查看：%s
📊 提交记录一目了然，快来一探究竟吧！
`
	TemplatePing = `
🔔 来自%s的消息
📌 仓库 %s
🎉 添加webhook成功
✅ 操作顺利，一切就绪！
🔗 详情查看：%s
`
)
