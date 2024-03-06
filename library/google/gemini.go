package google

import (
	"encoding/json"

	"github.com/opentdp/go-helper/request"
)

const ApiBaseUrl = "https://generativelanguage.googleapis.com"
const ApiVersion = "v1beta"

const ChatMessageRoleAssistant = "model"
const ChatMessageRoleUser = "user"

func NewClient(key string) *Client {

	return &Client{
		ApiBaseUrl: ApiBaseUrl,
		ApiVersion: ApiVersion,
		ApiKey:     key,
		Model:      "gemini-pro",
	}

}

func (c *Client) CreateChatCompletion(rq *RequestBody) (*ResponseBody, error) {

	body, _ := json.Marshal(rq)
	heaner := request.H{
		"Content-Type":   "application/json",
		"x-goog-api-key": c.ApiKey,
	}

	url := c.ApiBaseUrl + "/" + c.ApiVersion + "/models/" + c.Model + ":generateContent"
	response, err := request.Post(url, string(body), heaner)
	if err != nil {
		return nil, err
	}

	var resp ResponseBody
	err = json.Unmarshal(response, &resp)

	return &resp, err

}

func (c *Client) CreateImageCompletion(text, imageBase64, imageType string) (*ResponseBody, error) {

	rq := RequestBody{
		Contents: []*Content{
			{
				Parts: []*Part{
					{Text: text},
					{InlineData: &InlineData{Data: imageBase64, MimeType: imageType}},
				},
			},
		},
	}

	body, _ := json.Marshal(rq)
	heaner := request.H{
		"Content-Type":   "application/json",
		"x-goog-api-key": c.ApiKey,
	}

	url := c.ApiBaseUrl + "/" + c.ApiVersion + "/models/gemini-pro-vision:generateContent"
	response, err := request.Post(url, string(body), heaner)
	if err != nil {
		return nil, err
	}

	var resp ResponseBody
	err = json.Unmarshal(response, &resp)

	return &resp, err

}
