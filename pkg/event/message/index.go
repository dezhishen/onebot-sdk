package message

import (
	"encoding/json"

	"github.com/dezhishen/onebot-sdk/pkg/event/base"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type OnebotMessageEventCli interface {
	base.OnebotBaseEventCli
	ListenMessagePrivate(handler func(data model.EventMessagePrivate) error)
	ListenMessageGroup(handler func(data model.EventMessageGroup) error)
}

func NewOnebotMessageEventCli() (OnebotMessageEventCli, error) {
	return &websocketOnebotMessageEventCli{}, nil
}

type websocketOnebotMessageEventCli struct {
	messagePrivateHandlers []func(data model.EventMessagePrivate) error
	messageGroupHandlers   []func(data model.EventMessageGroup) error
}

func (c *websocketOnebotMessageEventCli) EventType() base.OnebotEventType {
	return base.OnebotEventTypeMessage
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
}

func (c *websocketOnebotMessageEventCli) ListenMessagePrivate(handler func(data model.EventMessagePrivate) error) {
	c.messagePrivateHandlers = append(c.messagePrivateHandlers, handler)
}

func (c *websocketOnebotMessageEventCli) ListenMessageGroup(handler func(data model.EventMessageGroup) error) {
	c.messageGroupHandlers = append(c.messageGroupHandlers, handler)
}
