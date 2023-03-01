package model

import "encoding/json"

type MessageElementShare struct {
	Url     string `json:"url"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
}

func (msg *MessageElementShare) Type() string {
	return "share"
}

func (msg *MessageElementShare) Enabled() bool {
	return true
}

// ProcessGRPC
func (msg *MessageElementShare) ProcessGRPC(segment *MessageSegmentGRPC) {
	segment.Data = &MessageSegmentGRPC_MessageElementShare{
		MessageElementShare: msg.ToGRPC(),
	}
}

func (msg *MessageElementShare) ToGRPC() *MessageElementShareGRPC {
	return &MessageElementShareGRPC{
		Url:     msg.Url,
		Title:   msg.Title,
		Content: msg.Content,
		Image:   msg.Image,
	}
}

func (msg *MessageElementShareGRPC) ToStruct() *MessageElementShare {
	return &MessageElementShare{
		Url:     msg.Url,
		Title:   msg.Title,
		Content: msg.Content,
		Image:   msg.Image,
	}
}

func init() {
	messageElementUnmarshalJSONMap["share"] = func(data []byte) (MessageElement, error) {
		var result MessageElementShare
		err := json.Unmarshal(data, &result)
		return &result, err
	}
	messageSegmentGRPCToStructMap["share"] = func(msg *MessageSegmentGRPC) (MessageElement, error) {
		if msg.Data == nil {
			return nil, nil
		}
		switch data := msg.Data.(type) {
		case *MessageSegmentGRPC_MessageElementShare:
			return data.MessageElementShare.ToStruct(), nil
		default:
			return nil, nil
		}
	}
}
