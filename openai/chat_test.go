package openai

import (
	"fmt"
	c "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
	"io/ioutil"
	"testing"
)

var answerBody = "\n\t{\n  \"id\": \"chatcmpl-6p5FEv1JHictSSnDZsGU4KvbuBsbu\",\n  \"object\": \"messages\",\n  \"created\": 1677693600,\n  \"model\": \"gpt-3.5-turbo\",\n  \"choices\": [\n    {\n      \"index\": 0,\n      \"finish_reason\": \"stop\",\n      \"messages\": [\n        {\n          \"role\": \"assistant\",\n          \"content\": \"OpenAI's mission is to ensure that artificial general intelligence benefits all of humanity.\"\n        }\n      ]\n    }\n  ],\n  \"usage\": {\n    \"prompt_tokens\": 20,\n    \"completion_tokens\": 18,\n    \"total_tokens\": 38\n  }\n}\n"
var answerBody1 = "{\n    \"id\":\"chatcmpl-6qCZ34NGbxSVt6RQOnm0kW7DiONb9\",\n    \"object\":\"chat.completion\",\n    \"created\":1677899373,\n    \"model\":\"gpt-3.5-turbo-0301\",\n    \"usage\":{\n        \"prompt_tokens\":9,\n        \"completion_tokens\":11,\n        \"total_tokens\":20\n    },\n    \"choices\":[\n        {\n            \"message\":{\n                \"role\":\"assistant\",\n                \"content\":\"\\n\\nHello! How can I assist you today?\"\n            },\n            \"finish_reason\":\"stop\",\n            \"index\":0\n        }\n    ]\n}"
func TestChatUnmarshalDemo(t *testing.T)  {
	c.Convey("test unmarshal answer", t, func() {
		answer, err := ChatClient.formatResp([]byte(answerBody1))
		c.So(err, c.ShouldBeNil)
		fmt.Println(answer)
	})
}

func TestSendRequest(t *testing.T)  {
	c.Convey("test send request", t, func() {
		req, _ := ChatClient.newReq("hello!")
		all, err := ioutil.ReadAll(req.Body)
		c.So(err, c.ShouldBeNil)
		fmt.Println(string(all))
	})
}

func TestGetAnswer(t *testing.T) {
	c.Convey("test get answer", t, func() {
		answer := ChatClient.GetAnswer("hello")
		fmt.Println(answer)
	})
}

func TestViperConfig(t *testing.T)  {
	vip := viper.New()
	vip.AddConfigPath("../config")
	vip.SetConfigName("chat_config") //设置读取的文件名
	vip.SetConfigType("yaml") //设置文件的类型
	if err := vip.ReadInConfig(); err != nil {
		panic(err)
	}
	var body *chatBody
	err := vip.Unmarshal(&body)
	if err != nil {
		panic(err)
	}
	fmt.Println(body)
}