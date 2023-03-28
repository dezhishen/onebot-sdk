package model

type EventMessageBase struct {
	EventBase
	MessageType string          `json:"message_type"`
	SubType     string          `json:"sub_type"`
	MessageId   int32           `json:"message_id"`
	UserId      int64           `json:"user_id"`
	Message     MessageSegments `json:"message"`
	RawMessage  string          `json:"raw_message"`
	Font        int32           `json:"font"`
	Sender      *Sender         `json:"sender"`
}

func (a *EventMessageBase) ToGRPC() *EventMeesageBaseGRPC {
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

func (a *EventMeesageBaseGRPC) ToStruct() *EventMessageBase {
	result := &EventMessageBase{
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
	EventMessageBase
}

func (a *EventMessagePrivate) ToGRPC() *EventMessagePrivateGRPC {
	return &EventMessagePrivateGRPC{
		EventMeesageBase: a.EventMessageBase.ToGRPC(),
	}
}

func (a *EventMessagePrivateGRPC) ToStruct() *EventMessagePrivate {
	return &EventMessagePrivate{
		EventMessageBase: *a.EventMeesageBase.ToStruct(),
	}
}

type EventMessageGroup struct {
	EventMessageBase
	GroupId   int64      `json:"group_id"`
	Anonymous *Anonymous `json:"anonymous"`
}

func (a *EventMessageGroup) ToGRPC() *EventMessageGroupGRPC {
	res := &EventMessageGroupGRPC{
		EventMeesageBase: a.EventMessageBase.ToGRPC(),
		GroupId:          a.GroupId,
	}
	if a.Anonymous != nil {
		res.Anonymous = a.Anonymous.ToGRPC()
	}
	return res
}

func (a *EventMessageGroupGRPC) ToStruct() *EventMessageGroup {
	res := &EventMessageGroup{
		EventMessageBase: *a.EventMeesageBase.ToStruct(),
		GroupId:          a.GroupId,
	}
	if a.Anonymous != nil {
		res.Anonymous = a.Anonymous.ToStruct()
	}
	return res
}
