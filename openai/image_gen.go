package openai

import (
	"encoding/json"
	gerrors "github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

const (
	urlImage = "https://api.openai.com/v1/images/generations"
)

type imageClient struct {}

type imageRequest struct {
	Prompt string `json:"prompt"`
	N    int `json:"n"`
	Size string `json:"size"`
}
type imageBody struct {
	Created int64 `json:"created"`
	Data []uri
}
type uri struct {
	Url string `json:"url"`
}
func (c *imageClient) newReq(describe string) (*http.Request, error) {
	reqBody := &imageRequest{}
	if err := readConfig("image_config", reqBody); err != nil {
		return nil, err
	}
	reqBody.Prompt = describe
	return CreateReq(http.MethodPost, urlImage, reqBody)
}

func (c *imageClient) formatImageUrl(body []byte) []string {
	var a imageBody
	if err := json.Unmarshal(body, &a); err != nil {
		logrus.Errorf("gen iamge unmarshal err: %v", err)
		return []string{""}
	}
	urls := make([]string, 0, len(a.Data))
	for _, url := range a.Data {
		urls = append(urls, url.Url)
	}
	return urls
}

func (c *imageClient)GetImage(describe string) ([]string, error) {
	req, err := c.newReq(describe)
	if err != nil {
		return []string{""}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorf("send gen image request err: %v", err)
		return []string{""}, gerrors.WithStack(err)
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("read gen image resp err: %v", err)
		return []string{""}, gerrors.WithStack(err)
	}
	urls := c.formatImageUrl(all)
	return urls, nil
}
