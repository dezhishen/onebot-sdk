package model

type MessageSegment struct {
	Type string         `json:"type"`
	Data MessageElement `json:"data"`
}

type MessageElement interface {
	Type() string
}

type FaceMessage struct {
}

func (msg *FaceMessage) Type() string {
	return "face"
}

type TextMessage struct {
	Text string `json:"text"`
}

func (msg *TextMessage) Type() string {
	return "text"
}
