package google

type Client struct {
	ApiBaseUrl string
	ApiVersion string
	ApiKey     string
	Model      string
}

type RequestBody struct {
	Contents []*Content `json:"contents"`
}

type ResponseBody struct {
	Candidates     []*Candidate `json:"candidates"`
	Error          *Error       `json:"error"`
	PromptFeedback struct {
		BlockReason   string `json:"blockReason"`
		SafetyRatings []struct {
			Category    string `json:"category"`
			Probability string `json:"probability"`
		} `json:"safetyRatings"`
	} `json:"promptFeedback"`
}

type Content struct {
	Parts []*Part `json:"parts"`
	Role  string  `json:"role"`
}

type Part struct {
	Text       string      `json:"text"`
	InlineData *InlineData `json:"inline_data"`
}

type InlineData struct {
	MimeType string `json:"mime_type"`
	Data     string `json:"data"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type Candidate struct {
	Content       *Content `json:"content"`
	FinishReason  string   `json:"finishReason"`
	Index         int      `json:"index"`
	SafetyRatings []*struct {
		Category    string `json:"category"`
		Probability string `json:"probability"`
	} `json:"safetyRatings"`
}
