package event

import (
	"encoding/json"

	"github.com/dezhishen/onebot-sdk/pkg/model"
)

var privateMsgHandlers []func(data model.EventPrivateMsg) error
var groupMsgHandlers []func(data model.EventGroupMsg) error

func init() {
	setHandler(
		EventTypeMessage,
		messageHandler,
	)
}

func messageHandler(data []byte) error {
	var message model.EventMsgBase
	err := json.Unmarshal(data, &message)
	if err != nil {
		return err
	}
	if message.MessageType == "group" {
		var groupMessage model.EventGroupMsg
		err = json.Unmarshal(data, &groupMessage)
		if err != nil {
			return err
		}
		for _, e := range groupMsgHandlers {
			err = e(groupMessage)
			if err != nil {
				return err
			}
		}
	} else {
		var privateMessage model.EventPrivateMsg
		err = json.Unmarshal(data, &privateMessage)
		if err != nil {
			return err
		}
		for _, e := range privateMsgHandlers {
			err = e(privateMessage)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func ListenReciverPrivateMsg(handler func(data model.EventPrivateMsg) error) {
	privateMsgHandlers = append(privateMsgHandlers, handler)
}

func ListenReciverGroupMsg(handler func(data model.EventGroupMsg) error) {
	groupMsgHandlers = append(groupMsgHandlers, handler)
}
