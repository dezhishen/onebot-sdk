package model

type PrivateMsg struct {
	UserId     int64             `json:"user_id"`
	GroupId    int64             `json:"group_id"`
	Message    []*MessageSegment `json:"message"`
	AutoEscape bool              `json:"auto_escape"`
}

func (msg *PrivateMsgGRPC) ToStruct() *PrivateMsg {
	return &PrivateMsg{
		UserId:     msg.UserId,
		GroupId:    msg.GroupId,
		Message:    MessageSegmentGRPCArray2MessageSegmentArray(msg.Message),
		AutoEscape: msg.AutoEscape,
	}
}

func (msg *PrivateMsg) ToGRPC() *PrivateMsgGRPC {
	return &PrivateMsgGRPC{
		UserId:     msg.UserId,
		GroupId:    msg.GroupId,
		AutoEscape: msg.AutoEscape,
		Message:    MessageSegmentArray2MessageSegmentGRPCArray(msg.Message),
	}
}

type GroupMsg struct {
	GroupId    int64             `json:"group_id"`
	Message    []*MessageSegment `json:"message"`
	AutoEscape bool              `json:"auto_escape"`
}

func (msg *GroupMsgGRPC) ToStruct() *GroupMsg {
	return &GroupMsg{
		GroupId:    msg.GroupId,
		Message:    MessageSegmentGRPCArray2MessageSegmentArray(msg.Message),
		AutoEscape: msg.AutoEscape,
	}
}

func (msg *GroupMsg) ToGRPC() *GroupMsgGRPC {
	return &GroupMsgGRPC{
		GroupId:    msg.GroupId,
		AutoEscape: msg.AutoEscape,
		Message:    MessageSegmentArray2MessageSegmentGRPCArray(msg.Message),
	}
}

type MessageType string

const (
	PrivateMessageType MessageType = "private"
	GroupMessageType   MessageType = "group"
)

type MsgForSend struct {
	UserId      int64             `json:"user_id"`
	GroupId     int64             `json:"group_id"`
	Message     []*MessageSegment `json:"message"`
	AutoEscape  bool              `json:"auto_escape"`
	MessageType MessageType       `json:"message_type"`
}

func (msg *MsgForSendGRPC) ToStruct() *MsgForSend {
	return &MsgForSend{
		GroupId:     msg.GroupId,
		Message:     MessageSegmentGRPCArray2MessageSegmentArray(msg.Message),
		AutoEscape:  msg.AutoEscape,
		MessageType: MessageType(msg.MessageType),
	}
}

func (msg *MsgForSend) ToGRPC() *MsgForSendGRPC {
	return &MsgForSendGRPC{
		GroupId:     msg.GroupId,
		AutoEscape:  msg.AutoEscape,
		Message:     MessageSegmentArray2MessageSegmentGRPCArray(msg.Message),
		MessageType: string(msg.MessageType),
	}
}

type SendMessageResultData struct {
	MessageId int64 `json:"message_id"`
}

func (msg *SendMessageResultDataGRPC) ToStruct() *SendMessageResultData {
	return &SendMessageResultData{
		MessageId: msg.MessageId,
	}
}

func (msg *SendMessageResultData) ToGRPC() *SendMessageResultDataGRPC {
	return &SendMessageResultDataGRPC{
		MessageId: msg.MessageId,
	}
}

type SendMessageResult struct {
	Data    *SendMessageResultData `json:"data"`
	Retcode int64                  `json:"retcode"`
	Status  string                 `json:"status"`
}

func (a *SendMessageResult) ToGRPC() *SendMessageResultGRPC {
	return &SendMessageResultGRPC{
		Data:    a.Data.ToGRPC(),
		Retcode: a.Retcode,
		Status:  a.Status,
	}
}

func (a *SendMessageResultGRPC) ToStruct() *SendMessageResult {
	return &SendMessageResult{
		Data:    a.Data.ToStruct(),
		Retcode: a.Retcode,
		Status:  a.Status,
	}
}

type MessageData struct {
	//发送时间
	Time int64 `json:"time"`
	//消息类型，同 消息事件
	MessageType string `json:"message_type"`
	//消息 Id
	MessageId int64 `json:"message_id"`
	//消息真实 Id
	RealId int64 `json:"real_id"`
	//发送人信息，同 消息事件
	Sender *Sender `json:"sender"`
	//消息内容
	Message []*MessageSegment `json:"message"`
}

func (a *MessageData) ToGRPC() *MessageDataGRPC {
	return &MessageDataGRPC{
		Time:        a.Time,
		MessageType: a.MessageType,
		MessageId:   a.MessageId,
		RealId:      a.RealId,
		Sender:      a.Sender.ToGRPC(),
		Message:     MessageSegmentArray2MessageSegmentGRPCArray(a.Message),
	}
}

func (a *MessageDataGRPC) ToStruct() *MessageData {
	return &MessageData{
		Time:        a.Time,
		MessageType: a.MessageType,
		MessageId:   a.MessageId,
		RealId:      a.RealId,
		Sender:      a.Sender.ToStruct(),
		Message:     MessageSegmentGRPCArray2MessageSegmentArray(a.Message),
	}
}

type MessageDataResult struct {
	Data    *MessageData `json:"data"`
	Retcode int64        `json:"retcode"`
	Status  string       `json:"status"`
}

func (a *MessageDataResult) ToGRPC() *MessageDataResultGRPC {
	return &MessageDataResultGRPC{
		Data:    a.Data.ToGRPC(),
		Retcode: a.Retcode,
		Status:  a.Status,
	}
}

func (a *MessageDataResultGRPC) ToStruct() *MessageDataResult {
	return &MessageDataResult{
		Data:    a.Data.ToStruct(),
		Retcode: a.Retcode,
		Status:  a.Status,
	}
}

type ForwardMessageData struct {
	Message []*MessageSegment `json:"message"`
}

func (a *ForwardMessageData) ToGRPC() *ForwardMessageDataGRPC {
	return &ForwardMessageDataGRPC{
		Message: MessageSegmentArray2MessageSegmentGRPCArray(a.Message),
	}
}

func (a *ForwardMessageDataGRPC) ToStruct() *ForwardMessageData {
	return &ForwardMessageData{
		Message: MessageSegmentGRPCArray2MessageSegmentArray(a.Message),
	}
}

type ForwardMessageDataResult struct {
	Data    *ForwardMessageData `json:"data"`
	Retcode int64               `json:"retcode"`
	Status  string              `json:"status"`
}

func (a *ForwardMessageDataResult) ToGRPC() *ForwardMessageDataResultGRPC {
	return &ForwardMessageDataResultGRPC{
		Data:    a.Data.ToGRPC(),
		Retcode: a.Retcode,
		Status:  a.Status,
	}
}

func (a *ForwardMessageDataResultGRPC) ToStruct() *ForwardMessageDataResult {
	return &ForwardMessageDataResult{
		Data:    a.Data.ToStruct(),
		Retcode: a.Retcode,
		Status:  a.Status,
	}
}
