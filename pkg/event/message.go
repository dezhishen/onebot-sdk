package event

import (
	"encoding/json"

	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type onebotMessageEventCli interface {
	onebotBaseEventCli
	ListenMessagePrivate(handler func(data model.EventMessagePrivate) error)
	ListenMessageGroup(handler func(data model.EventMessageGroup) error)
}

func NewOnebotMessageEventCli() (onebotMessageEventCli, error) {
	return &websocketOnebotMessageEventCli{}, nil
}

type websocketOnebotMessageEventCli struct {
	messagePrivateHandlers []func(data model.EventMessagePrivate) error
	messageGroupHandlers   []func(data model.EventMessageGroup) error
}

func (c *websocketOnebotMessageEventCli) EventType() OnebotEventType {
	return OnebotEventTypeMessage
}

func (c *websocketOnebotMessageEventCli) Handler(data []byte) error {
	var message model.EventMeesageBase
	err := json.Unmarshal(data, &message)
	if err != nil {
		return err
	}
	if message.MessageType == "group" {
		var groupMessage model.EventMessageGroup
		err = json.Unmarshal(data, &groupMessage)
		if err != nil {
			return err
		}
		for _, e := range c.messageGroupHandlers {
			err = e(groupMessage)
			if err != nil {
				return err
			}
		}
	} else {
		var privateMessage model.EventMessagePrivate
		err = json.Unmarshal(data, &privateMessage)
		if err != nil {
			return err
		}
		for _, e := range c.messagePrivateHandlers {
			err = e(privateMessage)
			if err != nil {
				return err
			}
		}
	}
	return nil
	// setHandler(c.EventType(), handler)
}

func (c *websocketOnebotMessageEventCli) ListenMessagePrivate(handler func(data model.EventMessagePrivate) error) {
	c.messagePrivateHandlers = append(c.messagePrivateHandlers, handler)
}

func (c *websocketOnebotMessageEventCli) ListenMessageGroup(handler func(data model.EventMessageGroup) error) {
	c.messageGroupHandlers = append(c.messageGroupHandlers, handler)
}
