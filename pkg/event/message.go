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
	if message.SubType == "group" {
		var data model.EventGroupMsg
		for _, e := range groupMsgHandlers {
			err = e(data)
			if err != nil {
				return err
			}
		}
	} else {
		var data model.EventPrivateMsg
		for _, e := range privateMsgHandlers {
			err = e(data)
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
