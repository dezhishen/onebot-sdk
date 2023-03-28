package model

import "encoding/json"

type MessageElementText struct {
	Text string `json:"text"`
}

func (msg *MessageElementText) Type() string {
	return "text"
}

func (msg *MessageElementText) Enabled() bool {
	return true
}

func (msg *MessageElementText) ToCQCode() string {
	return msg.Text
}

// ProcessGRPC
func (msg *MessageElementText) ProcessGRPC(segment *MessageSegmentGRPC) {
	segment.Data = &MessageSegmentGRPC_MessageElementText{
		MessageElementText: msg.ToGRPC(),
	}
}
func (msg *MessageElementText) ToGRPC() *MessageElementTextGRPC {
	return &MessageElementTextGRPC{
		Text: msg.Text,
	}
}

func (msg *MessageElementTextGRPC) ToStruct() *MessageElementText {
	return &MessageElementText{
		Text: msg.Text,
	}
}

func init() {
	messageElementUnmarshalJSONMap["text"] = func(data []byte) (MessageElement, error) {
		var result MessageElementText
		err := json.Unmarshal(data, &result)
		return &result, err
	}
	messageSegmentGRPCToStructMap["text"] = func(msg *MessageSegmentGRPC) (MessageElement, error) {
		if msg.Data == nil {
			return nil, nil
		}
		switch data := msg.Data.(type) {
		case *MessageSegmentGRPC_MessageElementText:
			return data.MessageElementText.ToStruct(), nil
		default:
			return nil, nil
		}
	}
}
