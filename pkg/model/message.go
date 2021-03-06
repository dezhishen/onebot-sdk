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
	Msg     string                 `json:"msg"`
	Wording string                 `json:"wording"`
}

func (a *SendMessageResult) ToGRPC() *SendMessageResultGRPC {
	result := &SendMessageResultGRPC{
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
	}
	if a.Data != nil {
		result.Data = a.Data.ToGRPC()
	}
	return result
}

func (a *SendMessageResultGRPC) ToStruct() *SendMessageResult {
	result := &SendMessageResult{
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
	}
	if a.Data != nil {
		result.Data = a.Data.ToStruct()
	}
	return result
}

type MessageData struct {
	//????????????
	Time int64 `json:"time"`
	//?????????????????? ????????????
	MessageType string `json:"message_type"`
	//?????? Id
	MessageId int64 `json:"message_id"`
	//???????????? Id
	RealId int64 `json:"real_id"`
	//????????????????????? ????????????
	Sender *Sender `json:"sender"`
	//????????????
	Message []*MessageSegment `json:"message"`
}

func (a *MessageData) ToGRPC() *MessageDataGRPC {
	result := &MessageDataGRPC{
		Time:        a.Time,
		MessageType: a.MessageType,
		MessageId:   a.MessageId,
		RealId:      a.RealId,
		Message:     MessageSegmentArray2MessageSegmentGRPCArray(a.Message),
	}
	if a.Sender != nil {
		result.Sender = a.Sender.ToGRPC()
	}
	return result
}

func (a *MessageDataGRPC) ToStruct() *MessageData {
	result := &MessageData{
		Time:        a.Time,
		MessageType: a.MessageType,
		MessageId:   a.MessageId,
		RealId:      a.RealId,
		Message:     MessageSegmentGRPCArray2MessageSegmentArray(a.Message),
	}
	if a.Sender != nil {
		result.Sender = a.Sender.ToStruct()
	}
	return result
}

type MessageDataResult struct {
	Data    *MessageData `json:"data"`
	Retcode int64        `json:"retcode"`
	Status  string       `json:"status"`
	Msg     string       `json:"msg"`
	Wording string       `json:"wording"`
}

func (a *MessageDataResult) ToGRPC() *MessageDataResultGRPC {
	result := &MessageDataResultGRPC{
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
	}
	if a.Data != nil {
		result.Data = a.Data.ToGRPC()
	}
	return result
}

func (a *MessageDataResultGRPC) ToStruct() *MessageDataResult {
	result := &MessageDataResult{
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
	}
	if a.Data != nil {
		result.Data = a.Data.ToStruct()
	}
	return result
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
	Msg     string              `json:"msg"`
	Wording string              `json:"wording"`
}

func (a *ForwardMessageDataResult) ToGRPC() *ForwardMessageDataResultGRPC {
	result := &ForwardMessageDataResultGRPC{
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
	}
	if a.Data != nil {
		result.Data = a.Data.ToGRPC()
	}
	return result
}

func (a *ForwardMessageDataResultGRPC) ToStruct() *ForwardMessageDataResult {
	result := &ForwardMessageDataResult{
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
	}
	if a.Data != nil {
		result.Data = a.Data.ToStruct()
	}
	return result
}

type SendGroupForwardMessageData struct {
	MessageId int64 `json:"message_id"`
}

func (a *SendGroupForwardMessageData) ToGRPC() *SendGroupForwardMessageDataGRPC {
	return &SendGroupForwardMessageDataGRPC{
		MessageId: a.MessageId,
	}
}

func (a *SendGroupForwardMessageDataGRPC) ToStruct() *SendGroupForwardMessageData {
	return &SendGroupForwardMessageData{
		MessageId: a.MessageId,
	}
}

type SendGroupForwardMessageDataResult struct {
	Retcode int64                        `json:"retcode"`
	Data    *SendGroupForwardMessageData `json:"data"`
	Status  string                       `json:"status"`
	Msg     string                       `json:"msg"`
	Wording string                       `json:"wording"`
}

func (a *SendGroupForwardMessageDataResult) ToGRPC() *SendGroupForwardMessageDataResultGRPC {
	result := &SendGroupForwardMessageDataResultGRPC{
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
	}
	if a.Data != nil {
		result.Data = a.Data.ToGRPC()
	}
	return result
}

func (a *SendGroupForwardMessageDataResultGRPC) ToStruct() *SendGroupForwardMessageDataResult {
	result := &SendGroupForwardMessageDataResult{
		Retcode: a.Retcode,
		Status:  a.Status,
		Msg:     a.Msg,
		Wording: a.Wording,
	}
	if a.Data != nil {
		result.Data = a.Data.ToStruct()
	}
	return result
}
