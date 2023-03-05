package openai

import (
	"encoding/json"
	gerrors "github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

const (
	urlText = "https://api.openai.com/v1/completions"
)

type textClient struct {}

type textAnswer struct {
	ID string `json:"id"`
	Object string `json:"object"`
	Created int64 `json:"created"`
	Model string	`json:"model"`
	Choices []choice2 `json:"choices"`
}
type choice2 struct {
	Text string    `json:"text"`
	Index int	   `json:"index"`
	Logprobs struct{} `json:"logprobs"`
	FinishReason string `json:"finish_reason"`
}

type textRequest struct {
	Model string `json:"model"`
	Prompt string `json:"prompt"`
	MaxTokens int `json:"max_tokens"`
	Temperature int	`json:"temperature"`
	TopP int `json:"top_p"`
	N int `json:"n" yaml:"n"`
	Stream bool `json:"stream"`
}

func (c *textClient) newReq(question string) (*http.Request, error) {
	reqBody := &textRequest{}
	if err := readConfig("text_config", reqBody); err != nil {
		return nil, err
	}
	reqBody.Prompt = question
	return CreateReq(http.MethodPost, urlText, reqBody)
}

func (c *textClient) formatResp(body []byte) (string, error) {
	var a textAnswer
	if err := json.Unmarshal(body, &a); err != nil {
		logrus.Errorf("gen text resp unmarshal err: %v", err)
		return "", gerrors.WithStack(err)
	}
	if err := validateText(a); err != nil {
		logrus.Errorf("validate text err: %v", err)
		return "", gerrors.WithStack(err)
	}
	return a.Choices[0].Text, nil
}

func validateText(answer textAnswer) error {
	if len(answer.Choices) < 1 {
		return gerrors.New("body choices nil")
	}
	return nil
}

func (c *textClient)GetText(question string) (string, error) {
	req, err := c.newReq(question)
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorf("send gen text request err: %v", err)
		return "", err
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("read gen text resp err: %v", err)
		return "", gerrors.WithStack(err)
	}
	text, err := c.formatResp(all)
	if err != nil {
		return "", gerrors.WithStack(err)
	}
	return text, err
}

