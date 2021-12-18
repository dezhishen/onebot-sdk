package event

import (
	"encoding/json"

	"github.com/dezhishen/onebot-sdk/pkg/model"
)

var messagePrivateHandlers []func(data model.EventMessagePrivate) error
var messageGroupHandlers []func(data model.EventMessageGroup) error

func init() {
	setHandler(
		EventTypeMessage,
		messageHandler,
	)
}

func messageHandler(data []byte) error {
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
		for _, e := range messageGroupHandlers {
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
		for _, e := range messagePrivateHandlers {
			err = e(privateMessage)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func ListenMessagePrivate(handler func(data model.EventMessagePrivate) error) {
	messagePrivateHandlers = append(messagePrivateHandlers, handler)
}

func ListenMessageGroup(handler func(data model.EventMessageGroup) error) {
	messageGroupHandlers = append(messageGroupHandlers, handler)
}
