package model

type EventMeesageBase struct {
	EventBase
	MessageType string            `json:"message_type"`
	SubType     string            `json:"sub_type"`
	MessageId   int32             `json:"message_id"`
	UserId      int64             `json:"user_id"`
	Message     []*MessageSegment `json:"message"`
	RawMessage  string            `json:"raw_message"`
	Font        int32             `json:"font"`
	Sender      *Sender           `json:"sender"`
}

func (a *EventMeesageBase) ToGRPC() *EventMeesageBaseGRPC {
	result := &EventMeesageBaseGRPC{
		EventBase:   a.EventBase.ToGRPC(),
		MessageType: a.MessageType,
		SubType:     a.SubType,
		MessageId:   a.MessageId,
		Message:     MessageSegmentArray2MessageSegmentGRPCArray(a.Message),
		RawMessage:  a.RawMessage,
		Font:        a.Font,
	}
	if a.Sender != nil {
		result.Sender = a.Sender.ToGRPC()
	}
	return result
}

func (a *EventMeesageBaseGRPC) ToStruct() *EventMeesageBase {
	result := &EventMeesageBase{
		EventBase:   *a.EventBase.ToStruct(),
		MessageType: a.MessageType,
		SubType:     a.SubType,
		MessageId:   a.MessageId,
		Message:     MessageSegmentGRPCArray2MessageSegmentArray(a.Message),
		RawMessage:  a.RawMessage,
		Font:        a.Font,
	}
	if a.Sender != nil {
		result.Sender = a.Sender.ToStruct()
	}
	return result
}

type EventMessagePrivate struct {
	EventMeesageBase
}

func (a *EventMessagePrivate) ToGRPC() *EventMessagePrivateGRPC {
	return &EventMessagePrivateGRPC{
		EventMeesageBase: a.EventMeesageBase.ToGRPC(),
	}
}

func (a *EventMessagePrivateGRPC) ToStruct() *EventMessagePrivate {
	return &EventMessagePrivate{
		EventMeesageBase: *a.EventMeesageBase.ToStruct(),
	}
}

type EventMessageGroup struct {
	EventMeesageBase
	GroupId   int64      `json:"group_id"`
	Anonymous *Anonymous `json:"anonymous"`
}

func (a *EventMessageGroup) ToGRPC() *EventMessageGroupGRPC {
	res := &EventMessageGroupGRPC{
		EventMeesageBase: a.EventMeesageBase.ToGRPC(),
		GroupId:          a.GroupId,
	}
	if a.Anonymous != nil {
		res.Anonymous = a.Anonymous.ToGRPC()
	}
	return res
}

func (a *EventMessageGroupGRPC) ToStruct() *EventMessageGroup {
	res := &EventMessageGroup{
		EventMeesageBase: *a.EventMeesageBase.ToStruct(),
		GroupId:          a.GroupId,
	}
	if a.Anonymous != nil {
		res.Anonymous = a.Anonymous.ToStruct()
	}
	return res
}
