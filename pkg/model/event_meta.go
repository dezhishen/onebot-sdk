package model

type EventMetaBase struct {
	EventBase
	MetaEventType string `json:"meta_event_type"`
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

type EventMetaHeartbeat struct {
	EventMetaBase
	Status map[string]interface{} `json:"status"`
	//到下次心跳的间隔，单位毫秒
	Interval int64 `json:"interval"`
}
