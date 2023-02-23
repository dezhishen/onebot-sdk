package event

import (
	"encoding/json"

	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type onebotMetaEventCli interface {
	onebotBaseEventCli
	ListenMetaLifecycle(handler func(data model.EventMetaLifecycle) error)
	ListenMetaHeartBeat(handler func(data model.EventMetaHeartbeat) error)
}

func NewOnebotMetaEventCli() (onebotMetaEventCli, error) {
	return &websocketOnebotMetaEventCli{}, nil
}

type websocketOnebotMetaEventCli struct {
	metaLifecycleHandlers []func(data model.EventMetaLifecycle) error
	metaHeartbeatHandlers []func(data model.EventMetaHeartbeat) error
}

func (c *websocketOnebotMetaEventCli) EventType() OnebotEventType {
	return OnebotEventTypeMetaEvent
}

func (c *websocketOnebotMetaEventCli) Handler(data []byte) error {
	return c.metaHandler(data)
}

func (c *websocketOnebotMetaEventCli) ListenMetaLifecycle(handler func(data model.EventMetaLifecycle) error) {
	c.metaLifecycleHandlers = append(c.metaLifecycleHandlers, handler)
}

func (c *websocketOnebotMetaEventCli) ListenMetaHeartBeat(handler func(data model.EventMetaHeartbeat) error) {
	c.metaHeartbeatHandlers = append(c.metaHeartbeatHandlers, handler)
}

func (c *websocketOnebotMetaEventCli) metaHandler(data []byte) error {
	var meta model.EventMetaBase
	err := json.Unmarshal(data, &meta)
	if err != nil {
		return err
	}
	if meta.MetaEventType == "lifecycle" {
		var relMeta model.EventMetaLifecycle
		err = json.Unmarshal(data, &relMeta)
		if err != nil {
			return err
		}
		for _, e := range c.metaLifecycleHandlers {
			err = e(relMeta)
			if err != nil {
				return err
			}
		}
	} else if meta.MetaEventType == "heartbeat" {
		var reqMeta model.EventMetaHeartbeat
		err = json.Unmarshal(data, &reqMeta)
		if err != nil {
			return err
		}
		for _, e := range c.metaHeartbeatHandlers {
			err = e(reqMeta)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
