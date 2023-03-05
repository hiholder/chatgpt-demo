package openai

import (
	"bytes"
	"encoding/json"
	"flag"
	gerrors "github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

var (
	client *http.Client
	ImageClient *imageClient
	ChatClient *chatClient
	TextClient *textClient
	EditClient *editClient
)

func CreateReq(method, url string, reqBody interface{}) (*http.Request, error) {
	marshal, err := json.Marshal(reqBody)
	if err != nil {
		logrus.Errorf("req json marshal err: %v, body: %v", err, reqBody)
		return nil, gerrors.WithStack(err)
	}
	req, err := http.NewRequest(method, url, bytes.NewReader(marshal))
	if err != nil {
		logrus.Errorf("new req err: %v, body: %v", err, reqBody)
		return nil, gerrors.WithStack(err)
	}
	vip := viper.New()
	if err := setConfigPath("key", vip); err != nil {
		logrus.Errorf("set config path err: %v", err)
		return nil, gerrors.WithStack(err)
	}
	req.Header.Add("Authorization", vip.Get("Authorization").(string))
	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

func readConfig(configName string, body interface{}) error {
	vip := viper.New()
	if err := setConfigPath(configName, vip); err != nil {
		return err
	}
	err := vip.Unmarshal(body)
	if err != nil {
		logrus.Errorf("config unmarshal struct err: %v", err)
		return gerrors.WithStack(err)
	}
	return nil
}

func setConfigPath(configName string, vip *viper.Viper) error {
	if flag.Lookup("test.v") != nil {
		vip.AddConfigPath("../config")
	} else {
		path, err := os.Getwd()
		if err != nil {
			return gerrors.WithStack(err)
		}
		vip.AddConfigPath(path + "/config")  //设置读取的文件路径
	}
	vip.SetConfigName(configName) //设置读取的文件名
	vip.SetConfigType("yaml") //设置文件的类型
	if err := vip.ReadInConfig(); err != nil {
		return gerrors.WithStack(err)
	}
	return nil
}

func init() {
	client = &http.Client{}
}

