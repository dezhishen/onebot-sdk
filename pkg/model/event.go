package model

type EventBase struct {
	Time     int64  `json:"time"`
	SelfId   int64  `json:"self_id"`
	PostType string `json:"post_type"`
}

func (a *EventBase) ToGRPC() *EventBaseGRPC {
	return &EventBaseGRPC{
		Time:     a.Time,
		SelfId:   a.SelfId,
		PostType: a.PostType,
	}
}

func (a *EventBaseGRPC) ToStruct() *EventBase {
	return &EventBase{
		Time:     a.Time,
		SelfId:   a.SelfId,
		PostType: a.PostType,
	}
}
