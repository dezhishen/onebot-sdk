package channel

import (
	"context"
	"errors"
	"fmt"

	"github.com/dezhishen/onebot-sdk/pkg/config"
)

type EventChannel interface {
	StartListen(context.Context, func(message []byte) error) error
	// callback func(message []byte) error
}

func NewEventChannel(conf *config.OnebotEventConfig) (EventChannel, error) {
	if conf == nil {
		return nil, errors.New("config is nil")
	}
	if conf.Type == "" {
		return nil, errors.New("type is empty")
	}
	switch conf.Type {
	// websocket
	case "websocket":
		return NewWebsocketEventChannel(conf)
	// 反向http
	case "http-reverse":
		return NewHttpReverseEventChannel(conf)
	// 反向websocket
	case "websocket-reverse":
		return NewWebsocketReverseEventChannel(conf)
	default:
		return nil, fmt.Errorf(
			"type %s not support,just support [%s,%s,%s]",
			conf.Type,
			"websocket",
			"http-reverse",
			"websocket-reverse",
		)
	}

}
