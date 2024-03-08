package google

type Client struct {
	ApiBaseUrl string
	ApiVersion string
	ApiKey     string
	Model      string
}

// request

type RequestBody struct {
	Contents       []*Content       `json:"contents"`
	SafetySettings []*SafetySetting `json:"safetySettings"`
}

type Content struct {
	Parts []*Part `json:"parts"`
	Role  string  `json:"role,omitempty"`
}

type Part struct {
	Text       string      `json:"text,omitempty"`
	InlineData *InlineData `json:"inline_data,omitempty"`
}

type InlineData struct {
	MimeType string `json:"mime_type"`
	Data     string `json:"data"`
}

type SafetySetting struct {
	Category  string `json:"category"`
	Threshold string `json:"threshold"`
}

// response

type ResponseBody struct {
	Candidates     []*Candidate `json:"candidates"`
	Error          *Error       `json:"error"`
	PromptFeedback struct {
		BlockReason   string          `json:"blockReason"`
		SafetyRatings []*SafetyRating `json:"safetyRatings"`
	} `json:"promptFeedback"`
}

type Candidate struct {
	Content       *Content        `json:"content"`
	FinishReason  string          `json:"finishReason"`
	Index         int             `json:"index"`
	SafetyRatings []*SafetyRating `json:"safetyRatings"`
}

type SafetyRating struct {
	Category    string `json:"category"`
	Probability string `json:"probability"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}
