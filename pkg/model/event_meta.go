package model

type EventMetaBase struct {
	EventBase
	MetaEventType string `json:"meta_event_type"`
}

func (a *EventMetaBase) ToGRPC() *EventMetaBaseGRPC {
	return &EventMetaBaseGRPC{
		EventBase:     a.EventBase.ToGRPC(),
		MetaEventType: a.MetaEventType,
	}
}

func (a *EventMetaBaseGRPC) ToStruct() *EventMetaBase {
	return &EventMetaBase{
		EventBase:     *a.EventBase.ToStruct(),
		MetaEventType: a.MetaEventType,
	}
}

type EventMetaLifecycle struct {
	EventMetaBase
	/*
		enable、disable、connect	事件子类型，分别表示 OneBot 启用、停用、WebSocket 连接成功

		注意，目前生命周期元事件中，只有 HTTP POST 的情况下可以收到 enable 和 disable

		只有正向 WebSocket 和反向 WebSocket 可以收到 connect。
	*/
	SubType string `json:"sub_type"`
}

func (a *EventMetaLifecycle) ToGRPC() *EventMetaLifecycleGRPC {
	return &EventMetaLifecycleGRPC{
		EventMetaBase: a.EventMetaBase.ToGRPC(),
		SubType:       a.SubType,
	}
}

func (a *EventMetaLifecycleGRPC) ToStruct() *EventMetaLifecycle {
	return &EventMetaLifecycle{
		EventMetaBase: *a.EventMetaBase.ToStruct(),
		SubType:       a.SubType,
	}
}

type EventMetaHeartbeat struct {
	EventMetaBase
	Status map[string]string `json:"status"`
	//到下次心跳的间隔，单位毫秒
	Interval uint64 `json:"interval"`
}

func (a *EventMetaHeartbeat) ToGRPC() *EventMetaHeartbeatGRPC {
	return &EventMetaHeartbeatGRPC{
		EventMetaBase: a.EventMetaBase.ToGRPC(),
		Status:        a.Status,
		Interval:      a.Interval,
	}
}

func (a *EventMetaHeartbeatGRPC) ToStruct() *EventMetaHeartbeat {
	return &EventMetaHeartbeat{
		EventMetaBase: *a.EventMetaBase.ToStruct(),
		Status:        a.Status,
		Interval:      a.Interval,
	}
}
