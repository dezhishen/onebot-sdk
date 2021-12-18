package event

import (
	"encoding/json"

	"github.com/dezhishen/onebot-sdk/pkg/model"
)

var requestFriendHandlers []func(data model.EventRequestFriend) error
var requestGroupHandlers []func(data model.EventRequestGroup) error

func init() {
	setHandler(
		EventTypeRequest,
		requestHandler,
	)
}

func requestHandler(data []byte) error {
	var req model.EventRequestBase
	err := json.Unmarshal(data, &req)
	if err != nil {
		return err
	}
	if req.RequestType == "friend" {
		var relReq model.EventRequestFriend
		err = json.Unmarshal(data, &relReq)
		if err != nil {
			return err
		}
		for _, e := range requestFriendHandlers {
			err = e(relReq)
			if err != nil {
				return err
			}
		}
	} else if req.RequestType == "group" {
		var relReq model.EventRequestGroup
		err = json.Unmarshal(data, &relReq)
		if err != nil {
			return err
		}
		for _, e := range requestGroupHandlers {
			err = e(relReq)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func ListenRequestFriend(handler func(data model.EventRequestFriend) error) {
	requestFriendHandlers = append(requestFriendHandlers, handler)
}
func ListenRequestGroup(handler func(data model.EventRequestGroup) error) {
	requestGroupHandlers = append(requestGroupHandlers, handler)
}
