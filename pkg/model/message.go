package model

import (
	"fmt"

	"github.com/tidwall/gjson"
)

type PrivateMsg struct {
	UserID     int64             `json:"user_id"`
	GroupID    int64             `json:"group_id"`
	Message    []*MessageSegment `json:"message"`
	AutoEscape bool              `json:"auto_escape"`
}

type GroupMsg struct {
	GroupID    int64             `json:"group_id"`
	Message    []*MessageSegment `json:"message"`
	AutoEscape bool              `json:"auto_escape"`
}

type MsgNode struct {
	ID      int32             `json:"id"`
	Name    string            `json:"name"`
	Uin     int64             `json:"uin"`
	Content []*MessageSegment `json:"content"`
	Seq     []*MessageSegment `json:"seq"`
}

type GroupForwardMsg struct {
	GroupID  int64      `json:"group_id"`
	Messages []*MsgNode `json:"messages"`
}

type MessageType string

const (
	PrivateMessageType MessageType = "private"
	GroupMessageType   MessageType = "group"
)

type MsgForSend struct {
	UserID      int64             `json:"user_id"`
	GroupID     int64             `json:"group_id"`
	Message     []*MessageSegment `json:"message"`
	AutoEscape  bool              `json:"auto_escape"`
	MessageType MessageType       `json:"message_type"`
}

type SendMessageResultData struct {
	MessageId int `json:"message_id"`
}

type SendMessageResult struct {
	Data    *SendMessageResultData `json:"data"`
	RetCode int                    `json:"retcode"`
	Status  string                 `json:"status"`
}

type MessageResultData struct {
	//发送时间
	Time int `json:"time"`
	//消息类型，同 消息事件
	MessageType string `json:"message_type"`
	//消息 ID
	MessageID int `json:"message_id"`
	//消息真实 ID
	RealID int `json:"real_id"`
	//发送人信息，同 消息事件
	Sender *Sender `json:"sender"`
	//消息内容
	Message *MessageElement `json:"message"`
}

func (msgSeg *MessageResultData) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || data[0] == 'n' { // copied from the Q, can be improved
		return nil
	}
	messageType := gjson.GetBytes(data, "type").Str
	decoder, ok := messageElementUnmarshalJSONMap[messageType]
	if !ok {
		return fmt.Errorf("未找到指定的消息类型,%v", messageType)
	}
	element, err := decoder([]byte(gjson.GetBytes(data, "message").Raw))
	if !ok {
		return fmt.Errorf("未找到指定的消息类型,%v", messageType)
	}
	msgSeg.Message = &element
	return err
}

type MessageResult struct {
	Data    []*MessageResultData `json:"data"`
	RetCode int                  `json:"retcode"`
	Status  string               `json:"status"`
}
type ForwardMessageResult struct {
	Data    []*MessageSegment `json:"data"`
	RetCode int               `json:"retcode"`
	Status  string            `json:"status"`
}
