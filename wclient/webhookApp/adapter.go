package webhookApp

func Handler(app string, msg string) string {
	var res string
	var err error

	switch app {
	case "github":
		res, err = GithubWebhook(msg)
	case "gitea":
		res, err = GiteaWebhook(msg)

	default:
		res = "暂不支持该应用的 webhook"
	}

	if err != nil {
		return err.Error()
	}

	return res
}
