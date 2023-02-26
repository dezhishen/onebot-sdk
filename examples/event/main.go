package main

// import (
// 	"context"
// 	"fmt"

// 	"github.com/dezhishen/onebot-sdk/pkg/api"
// 	"github.com/dezhishen/onebot-sdk/pkg/config"
// 	"github.com/dezhishen/onebot-sdk/pkg/event"
// 	"github.com/dezhishen/onebot-sdk/pkg/model"
// )

// func main() {
// 	websocketConf, err := config.LoadConfig("websocket")
// 	if err != nil {
// 		panic(err)
// 	}
// 	eventCli, err := event.NewOnebotEventCli(websocketConf)
// 	if err != nil {
// 		panic(err)
// 	}
// 	httpConf, err := config.LoadConfig("http")
// 	if err != nil {
// 		panic(err)
// 	}
// 	onebotApi, err := api.NewOnebotApiClientByConfig(httpConf)
// 	if err != nil {
// 		panic(err)
// 	}
// 	eventCli.ListenMessageGroup(func(data model.EventMessageGroup) error {
// 		if data.Message[0].Type == "text" {
// 			fmt.Println(data.Message[0].Data)
// 			result, err := onebotApi.SendGroupMsg(
// 				&model.GroupMsg{
// 					GroupId: data.GroupId,
// 					Message: []*model.MessageSegment{
// 						{
// 							Type: "text",
// 							Data: &model.MessageElementText{
// 								Text: "我复读机",
// 							},
// 						},
// 					},
// 				},
// 			)
// 			if err != nil {
// 				fmt.Printf("%v", err)
// 			}
// 			fmt.Println(result)
// 		}
// 		return nil
// 	})
// 	err = eventCli.StartWithContext(context.Background())
// 	if err != nil {
// 		panic(err)
// 	}

// }
