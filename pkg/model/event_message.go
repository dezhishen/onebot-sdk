package model

type EventMsgBase struct {
	EventBase
	MessageType string           `json:"message_type"`
	SubType     string           `json:"sub_type"`
	MessageID   int32            `json:"message_id"`
	UserID      int64            `json:"user_id"`
	Message     []MessageSegment `json:"message"`
	RawMessage  string           `json:"raw_message"`
	Font        int32            `json:"font"`
	Sender      *Sender          `json:"sender"`
}

type EventPrivateMsg struct {
	EventMsgBase
}

type EventGroupMsg struct {
	GroupID   int64      `json:"group_id"`
	Anonymous *Anonymous `json:"anonymous"`
	EventMsgBase
}
