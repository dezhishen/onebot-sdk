package model

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

type MsgForSend struct {
	UserID      int64             `json:"user_id"`
	GroupID     int64             `json:"group_id"`
	Message     []*MessageSegment `json:"message"`
	AutoEscape  bool              `json:"auto_escape"`
	MessageType string            `json:"message_type"`
}
