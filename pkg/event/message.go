package event

import (
	"encoding/json"

	"github.com/dezhishen/onebot-sdk/pkg/model"
)

func init() {
	allhandler[EventTypeMessage] = func(data []byte) error {
		var message model.EventMsgBase
		err := json.Unmarshal(data, &message)
		if err != nil {
			return err
		}
		if message.SubType == "group" {

		} else {

		}
		return nil
	}
}

func ListenReciverMsg(handler func(data model.EventPrivateMsg) error) {
}
