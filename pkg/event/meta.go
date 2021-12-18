package event

import (
	"encoding/json"

	"github.com/dezhishen/onebot-sdk/pkg/model"
)

var metaLifecycleHandlers []func(data model.EventMetaLifecycle) error
var metaHeartbeatHandlers []func(data model.EventMetaHeartbeat) error

func init() {
	setHandler(
		EventTypeRequest,
		metaHandler,
	)
}

func metaHandler(data []byte) error {
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
		for _, e := range metaLifecycleHandlers {
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
		for _, e := range metaHeartbeatHandlers {
			err = e(reqMeta)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func ListenMetaLifecycle(handler func(data model.EventMetaLifecycle) error) {
	metaLifecycleHandlers = append(metaLifecycleHandlers, handler)
}
func ListenMetaHeartBeat(handler func(data model.EventMetaHeartbeat) error) {
	metaHeartbeatHandlers = append(metaHeartbeatHandlers, handler)
}
