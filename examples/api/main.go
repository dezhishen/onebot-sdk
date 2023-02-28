package main

import (
	"github.com/dezhishen/onebot-sdk/pkg/api"
	"github.com/dezhishen/onebot-sdk/pkg/config"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

func main() {
	conf, err := config.LoadConfig("websocket-reverse")
	// conf, err := config.LoadConfig("websocket")
	// conf, err := config.LoadConfig("http")
	if err != nil {
		panic(err)
	}
	cli, err := api.NewOnebotApiClient(conf.Api)
	if err != nil {
		panic(err)
	}
	result, err := cli.SendGroupMsg(
		&model.GroupMsg{
			// change this to your group id
			GroupId: 123456,
			Message: []*model.MessageSegment{
				{
					Type: "text",
					Data: &model.MessageElementText{
						Text: "hello",
					},
				},
			},
		},
	)
	if err != nil {
		panic(err)
	}
	println(result)
}
