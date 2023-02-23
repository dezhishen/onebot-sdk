package client

import "github.com/dezhishen/onebot-sdk/pkg/config"

type WebsocketCli struct {
	cli         interface{}
	endpoint    string
	accessToken string
}

func NewWebsocketCli(conf config.OnebotConfig) *WebsocketCli {
	return &WebsocketCli{
		cli:         nil,
		endpoint:    conf.Endpoint,
		accessToken: conf.AccessToken,
	}
}
