package model

type EventMetaBase struct {
	EventBase
	MetaEventType string `json:"meta_event_type"`
}

func (e *EventMetaBase) ToGRPC() *EventMetaBaseGRPC {
	return &EventMetaBaseGRPC{
		EventBase:     e.EventBase.ToGRPC(),
		MetaEventType: e.MetaEventType,
	}
}

func (e *EventMetaBaseGRPC) ToStruct() *EventMetaBase {
	return &EventMetaBase{
		EventBase:     *e.EventBase.ToStruct(),
		MetaEventType: e.MetaEventType,
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

func (e *EventMetaLifecycle) ToGRPC() *EventMetaLifecycleGRPC {
	return &EventMetaLifecycleGRPC{
		EventMetaBase: e.EventMetaBase.ToGRPC(),
		SubType:       e.SubType,
	}
}

func (e *EventMetaLifecycleGRPC) ToStruct() *EventMetaLifecycle {
	return &EventMetaLifecycle{
		EventMetaBase: *e.EventMetaBase.ToStruct(),
		SubType:       e.SubType,
	}
}

type EventMetaHeartbeatStatusStat struct {
	PacketReceived  uint64 `json:"packet_received"`
	PacketSent      uint64 `json:"packet_sent"`
	PacketLost      uint64 `json:"packet_lost"`
	MessageReceived uint64 `json:"message_received"`
	MessageSent     uint64 `json:"message_sent"`
	DisconnectTimes uint64 `json:"disconnect_times"`
	LostTimes       uint64 `json:"lost_times"`
	LastMessageTime uint64 `json:"last_message_time"`
}

func (e *EventMetaHeartbeatStatusStat) ToGRPC() *EventMetaHeartbeatStatusStatGRPC {
	return &EventMetaHeartbeatStatusStatGRPC{
		PacketReceived:  e.PacketReceived,
		PacketSent:      e.PacketSent,
		PacketLost:      e.PacketLost,
		MessageReceived: e.MessageReceived,
		MessageSent:     e.MessageSent,
		DisconnectTimes: e.DisconnectTimes,
		LostTimes:       e.LostTimes,
		LastMessageTime: e.LastMessageTime,
	}
}

func (e *EventMetaHeartbeatStatusStatGRPC) ToStruct() *EventMetaHeartbeatStatusStat {
	return &EventMetaHeartbeatStatusStat{
		PacketReceived:  e.PacketReceived,
		PacketSent:      e.PacketSent,
		PacketLost:      e.PacketLost,
		MessageReceived: e.MessageReceived,
		MessageSent:     e.MessageSent,
		DisconnectTimes: e.DisconnectTimes,
		LostTimes:       e.LostTimes,
		LastMessageTime: e.LastMessageTime,
	}
}

type EventMetaHeartbeatStatus struct {
	AppEnabled     bool                          `json:"app_enabled"`
	AppGood        bool                          `json:"app_good"`
	AppInitialized bool                          `json:"app_initialized"`
	Good           bool                          `json:"good"`
	Online         bool                          `json:"online"`
	PluginsGood    bool                          `json:"plugin_good"`
	Stat           *EventMetaHeartbeatStatusStat `json:"stat"`
}

func (e *EventMetaHeartbeatStatus) ToGRPC() *EventMetaHeartbeatStatusGRPC {
	result := &EventMetaHeartbeatStatusGRPC{
		AppEnabled:     e.AppEnabled,
		AppGood:        e.AppGood,
		AppInitialized: e.AppInitialized,
		Good:           e.Good,
		Online:         e.Online,
		PluginsGood:    e.PluginsGood,
	}
	if e.Stat != nil {
		result.Stat = e.Stat.ToGRPC()
	}
	return result

}
func (e *EventMetaHeartbeatStatusGRPC) ToStruct() *EventMetaHeartbeatStatus {
	result := &EventMetaHeartbeatStatus{
		AppEnabled:     e.AppEnabled,
		AppGood:        e.AppGood,
		AppInitialized: e.AppInitialized,
		Good:           e.Good,
		Online:         e.Online,
		PluginsGood:    e.PluginsGood,
	}
	if e.Stat != nil {
		result.Stat = e.Stat.ToStruct()
	}
	return result
}

type EventMetaHeartbeat struct {
	EventMetaBase
	Status *EventMetaHeartbeatStatus `json:"status"`
	//到下次心跳的间隔，单位毫秒
	Interval uint64 `json:"interval"`
}

func (e *EventMetaHeartbeat) ToGRPC() *EventMetaHeartbeatGRPC {
	result := &EventMetaHeartbeatGRPC{
		EventMetaBase: e.EventMetaBase.ToGRPC(),
		// Status:        e.Status,
		Interval: e.Interval,
	}
	if e.Status != nil {
		result.Status = e.Status.ToGRPC()
	}
	return result
}

func (e *EventMetaHeartbeatGRPC) ToStruct() *EventMetaHeartbeat {
	result := &EventMetaHeartbeat{
		EventMetaBase: *e.EventMetaBase.ToStruct(),
		// Status:        e.Status,
		Interval: e.Interval,
	}
	if e.Status != nil {
		result.Status = e.Status.ToStruct()

	}
	return result
}
