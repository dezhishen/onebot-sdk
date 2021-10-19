package model

type MessageElement interface {
	Type() string
}
type FaceMessage struct {
	Type string `json:"type"`
}

type msg struct {
	Message    *MessageElement `json:"message"`
	AutoEscape bool            `json:"auto_escape"`
}

type PrivateMsg struct {
	UserID int64 `json:"user_id"`
	GroupMsg
}

type GroupMsg struct {
	GroupID int64 `json:"group_id"`
	msg
}

type MsgNode struct {
	ID      int32           `json:"id"`
	Name    string          `json:"name"`
	Uin     int64           `json:"uin"`
	Content *MessageElement `json:"content"`
	Seq     *MessageElement `json:"seq"`
}

type GroupForwardMsg struct {
	GroupID  int64     `json:"group_id"`
	Messages []MsgNode `json:"messages"`
}

type MsgForSend struct {
	PrivateMsg
	MessageType string `json:"message_type"`
}
