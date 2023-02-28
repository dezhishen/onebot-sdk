package channel

import (
	"errors"

	"github.com/dezhishen/onebot-sdk/pkg/config"
)

type ApiChannel interface {
	Post(url string) error
	PostForResult(url string, result interface{}) error
	PostByRequest(url string, req interface{}) error
	PostByRequestForResult(url string, req, result interface{}) error
}

func NewApiChannel(conf *config.OnebotApiConfig) (ApiChannel, error) {
	if conf == nil {
		return nil, errors.New("api config not found")
	}
	if conf.Type == "http" {
		return NewHttpApiChannel(conf)
	}
	if conf.Type == "websocket" {
		return NewWebsocketApiChannel(conf)
	}
	if conf.Type == "websocket-reverse" {
		return NewWebsocketReverseApiChannel(conf)
	}
	return nil, errors.New("channel not support")
}
