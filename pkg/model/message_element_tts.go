package model

import "encoding/json"

type MessageElementTts struct {
	// text	string	内容
	Text string `json:"text"`
}

func (msg *MessageElementTts) Type() string {
	return "tts"
}

func (msg *MessageElementTts) Enabled() bool {
	return true
}

// ProcessGRPC
func (msg *MessageElementTts) ProcessGRPC(segment *MessageSegmentGRPC) {
	segment.Data = &MessageSegmentGRPC_MessageElementTts{
		MessageElementTts: msg.ToGRPC(),
	}
}

func (msg *MessageElementTts) ToGRPC() *MessageElementTtsGRPC {
	return &MessageElementTtsGRPC{
		Text: msg.Text,
	}
}

func (msg *MessageElementTtsGRPC) ToStruct() *MessageElementTts {
	return &MessageElementTts{
		Text: msg.Text,
	}
}

func init() {
	messageElementUnmarshalJSONMap["tts"] = func(data []byte) (MessageElement, error) {
		var result MessageElementTts
		err := json.Unmarshal(data, &result)
		return &result, err
	}
	messageSegmentGRPCToStructMap["tts"] = func(msg *MessageSegmentGRPC) (MessageElement, error) {
		if msg.Data == nil {
			return nil, nil
		}
		switch data := msg.Data.(type) {
		case *MessageSegmentGRPC_MessageElementTts:
			return data.MessageElementTts.ToStruct(), nil
		default:
			return nil, nil
		}
	}
}
