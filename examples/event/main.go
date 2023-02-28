package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dezhishen/onebot-sdk/pkg/config"
	"github.com/dezhishen/onebot-sdk/pkg/event"
	"github.com/dezhishen/onebot-sdk/pkg/model"
	"github.com/sirupsen/logrus"
)

func main() {
	conf, err := config.LoadConfig("http-reverse")
	// conf, err := config.LoadConfig("websocket-reverse")
	// conf, err := config.LoadConfig("websocket")

	if err != nil {
		panic(err)
	}
	eventCli, err := event.NewOnebotEventCli(conf.Event)
	if err != nil {
		panic(err)
	}
	eventCli.ListenMessageGroup(func(data model.EventMessageGroup) error {
		if data.Message[0].Type == "text" {
			fmt.Println(data.Message[0].Data)
		}
		return nil
	})
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(10 * time.Minute)
		logrus.Info("cancel...")
		cancel()
	}()
	err = eventCli.StartListenWithCtx(ctx)
	if err != nil {
		panic(err)
	}
}
