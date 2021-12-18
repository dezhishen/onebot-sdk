package model

type EventMsgBase struct {
	EventBase
	MessageType string           `json:"message_type"`
	SubType     string           `json:"sub_type"`
	MessageID   int32            `json:"message_id"`
	UserID      int64            `json:"user_id"`
	Message     []MessageSegment `json:"message"`
}

type EventPrivateMsg struct {
	EventMsgBase
}

type EventGroupMsg struct {
	EventMsgBase
}
