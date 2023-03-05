package openai

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	editUrl = "https://api.openai.com/v1/edits"
)

type editClient struct{}

type EditReqBody struct {
	Model string `json:"model"`
	Input string `json:"input"`
	Instruction string `json:"instruction"`
}
type EditRespBody struct {
	Object string `json:"object"`
	Created int64 `json:"created"`
	Choices []EditChoice `json:"choices"`
}

type EditChoice struct {
	Text string `json:"text"`
	Index int `json:"index"`
}


func (c *editClient) GetEdit(input, instruction string) (string, error) {
	req, err := c.newReq(input, instruction)
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	edit, err := c.formatResp(all)
	if err != nil {
		return "", err
	}
	return edit, nil
}

func (c *editClient) newReq(input, instruction string) (*http.Request, error) {
	reqBody := &EditReqBody{}
	if err := readConfig("edit_config", reqBody); err != nil {
		return nil, err
	}
	reqBody.Input = input
	reqBody.Instruction = instruction
	fmt.Println(reqBody)
	return CreateReq(http.MethodPost, editUrl, reqBody)
}

func (c *editClient) formatResp(body []byte) (string, error) {
	var res *EditRespBody
	err := json.Unmarshal(body, &res)
	if err != nil {
		return "", err
	}
	return res.Choices[0].Text, nil
}

