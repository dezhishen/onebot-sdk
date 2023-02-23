package event

import (
	"encoding/json"

	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type onebotRequestEventCli interface {
	onebotBaseEventCli
	ListenRequestFriend(handler func(data model.EventRequestFriend) error)
	ListenRequestGroup(handler func(data model.EventRequestGroup) error)
}

func NewOnebotRequestEventCli() (onebotRequestEventCli, error) {
	return &websocketOnebotRequestEventCli{}, nil
}

type websocketOnebotRequestEventCli struct {
	requestFriendHandlers []func(data model.EventRequestFriend) error
	requestGroupHandlers  []func(data model.EventRequestGroup) error
}

func (c *websocketOnebotRequestEventCli) EventType() OnebotEventType {
	return OnebotEventTypeRequest
}

func (c *websocketOnebotRequestEventCli) Handler(data []byte) error {
	return c.requestHandler(data)
}

func (c *websocketOnebotRequestEventCli) requestHandler(data []byte) error {
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
		for _, e := range c.requestFriendHandlers {
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
		for _, e := range c.requestGroupHandlers {
			err = e(relReq)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *websocketOnebotRequestEventCli) ListenRequestFriend(handler func(data model.EventRequestFriend) error) {
	c.requestFriendHandlers = append(c.requestFriendHandlers, handler)
}
func (c *websocketOnebotRequestEventCli) ListenRequestGroup(handler func(data model.EventRequestGroup) error) {
	c.requestGroupHandlers = append(c.requestGroupHandlers, handler)
}
