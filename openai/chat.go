package openai

import (
	"encoding/json"
	gerrors "github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

const (
	urlChat = "https://api.openai.com/v1/chat/completions"
)

type chatClient struct {}

type chatBody struct {
	Model   string  `json:"model"`
	Messages []message `json:"messages"`
}
type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type chatAnswer struct {
	Choices []choice `json:"choices"`
}

type choice struct {
	Message message `json:"message"`
}

func (c *chatClient) newReq(question string) (*http.Request, error) {
	reqBody := &chatBody{}
	if err := readConfig("chat_config", reqBody); err != nil {
		return nil, err
	}
	reqBody.Messages[0].Content = question
	return CreateReq(http.MethodPost, urlChat, reqBody)
}

func (c *chatClient)GetAnswer(question string) string {
	req, err := c.newReq(question)
	if err != nil {
		logrus.Errorf("format params err: %v", err)
		return ""
	}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorf("send request err: %v", err)
		return ""
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("read resp.Body err: %v", err)
		return ""
	}
	answer, err := c.formatResp(all)
	if err != nil {
		return ""
	}
	return answer
}

func (c *chatClient)formatResp(body []byte) (string, error) {
	var a chatAnswer
	if err := json.Unmarshal(body, &a); err != nil {
		logrus.Errorf("chat resp json unmarshal err: %v", err)
		return "", gerrors.WithStack(err)
	}
	if err := validateAnswer(a); err != nil {
		logrus.Errorf("validate answer err: %v", err)
		return "", gerrors.WithStack(err)
	}
	return a.Choices[0].Message.Content, nil
}

func validateAnswer(answer chatAnswer) error {
	if len(answer.Choices) < 1 {
		return gerrors.New("body choices nil")
	}
	return nil
}