package model

import "encoding/json"

type MessageElementRps struct {
}

func (msg *MessageElementRps) Type() string {
	return "rps"
}

func (msg *MessageElementRps) Enabled() bool {
	return false
}

func (msg *MessageElementRps) ToCQCode() string {
	return CQPrefix + msg.Type() + CQSuffix
}

// ProcessGRPC
func (msg *MessageElementRps) ProcessGRPC(segment *MessageSegmentGRPC) {
	segment.Data = &MessageSegmentGRPC_MessageElementRps{
		MessageElementRps: msg.ToGRPC(),
	}
}

func (msg *MessageElementRps) ToGRPC() *MessageElementRpsGRPC {
	var result MessageElementRpsGRPC
	return &result
}

func (msg *MessageElementRpsGRPC) ToStruct() *MessageElementRps {
	var result MessageElementRps
	return &result
}

func init() {
	messageElementUnmarshalJSONMap["rps"] = func(data []byte) (MessageElement, error) {
		var result MessageElementRps
		err := json.Unmarshal(data, &result)
		return &result, err
	}
	messageSegmentGRPCToStructMap["rps"] = func(msg *MessageSegmentGRPC) (MessageElement, error) {
		if msg.Data == nil {
			return nil, nil
		}
		switch data := msg.Data.(type) {
		case *MessageSegmentGRPC_MessageElementRps:
			return data.MessageElementRps.ToStruct(), nil
		default:
			return nil, nil
		}
	}
}
