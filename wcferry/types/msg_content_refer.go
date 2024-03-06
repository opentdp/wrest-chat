package types

// 引用消息
// @type 49
// @field Msg.Content

type ReferMsg struct {
	Msg struct {
		AppMsg struct {
			Action   string `json:"action"`
			ReferMsg struct {
				Svrid string `json:"svrid"`
			} `json:"refermsg"`
			Title string `json:"title"`
			Type  string `json:"type"`
			URL   string `json:"url"`
		} `json:"appmsg"`
		CommentURL   string `json:"commenturl"`
		FromUsername string `json:"fromusername"`
		Scene        string `json:"scene"`
	} `json:"msg"`
}
