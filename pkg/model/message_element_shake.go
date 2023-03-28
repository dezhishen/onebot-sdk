package model

import "encoding/json"

type MessageElementShake struct {
}

func (msg *MessageElementShake) Type() string {
	return "shake"
}

func (msg *MessageElementShake) Enabled() bool {
	return false
}

func (msg *MessageElementShake) ToCQCode() string {
	return CQPrefix + msg.Type() + CQSuffix
}

// ProcessGRPC
func (msg *MessageElementShake) ProcessGRPC(segment *MessageSegmentGRPC) {
	segment.Data = &MessageSegmentGRPC_MessageElementShake{
		MessageElementShake: msg.ToGRPC(),
	}
}

func (msg *MessageElementShake) ToGRPC() *MessageElementShakeGRPC {
	var result MessageElementShakeGRPC
	return &result
}

func (msg *MessageElementShakeGRPC) ToStruct() *MessageElementShake {
	var result MessageElementShake
	return &result
}

func init() {
	messageElementUnmarshalJSONMap["shake"] = func(data []byte) (MessageElement, error) {
		var result MessageElementShake
		err := json.Unmarshal(data, &result)
		return &result, err
	}
	messageSegmentGRPCToStructMap["shake"] = func(msg *MessageSegmentGRPC) (MessageElement, error) {
		if msg.Data == nil {
			return nil, nil
		}
		switch data := msg.Data.(type) {
		case *MessageSegmentGRPC_MessageElementShake:
			return data.MessageElementShake.ToStruct(), nil
		default:
			return nil, nil
		}
	}
}
