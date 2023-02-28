package main

import (
	"github.com/dezhishen/onebot-sdk/pkg/api"
	"github.com/dezhishen/onebot-sdk/pkg/config"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

func main() {
	conf, err := config.LoadConfig("websocket-reverse")
	if err != nil {
		panic(err)
	}
	cli, err := api.NewOnebotApiClient(conf.Api)
	if err != nil {
		panic(err)
	}
	result, err := cli.SendGroupMsg(
		&model.GroupMsg{
			GroupId: 727670105,
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
