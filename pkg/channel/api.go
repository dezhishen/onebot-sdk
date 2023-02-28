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

func NewApiChannel(conf *config.OnebotConfig) (ApiChannel, error) {
	if conf.Api == nil {
		return nil, errors.New("api config not found")
	}
	if conf.Api.Type == "http" {
		return NewHttpApiChannel(conf.Api), nil
	}
	if conf.Api.Type == "websocket" {
		return NewWebsocketApiChannel(conf.Api), nil
	}
	return nil, errors.New("channel not support")
}
