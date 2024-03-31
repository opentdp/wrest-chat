package webhookApp

import "net/http"

func Handler(header http.Header, app string, msg string) string {
	var res string
	var err error

	switch app {
	case "github":
		res, err = GithubWebhook(header, msg)
	case "gitea":
		res, err = GiteaWebhook(msg)
	case "text":
		res, err = TextWebhook(msg)
	default:
		res = "暂不支持该应用的 webhook"
	}

	if err != nil {
		return err.Error()
	}

	return res
}
