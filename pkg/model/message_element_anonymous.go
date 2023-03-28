package model

import (
	"encoding/json"
)

type MessageElementAnonymous struct {
	// ignore		✓	0 1	可选, 表示无法匿名时是否继续发送
	Ignore string `json:"ignore,omitempty"`
}

func (msg *MessageElementAnonymous) Type() string {
	return "anonymous"
}

func (msg *MessageElementAnonymous) Enabled() bool {
	return false
}

func (msg *MessageElementAnonymous) ToCQCode() string {
	return CQPrefix + msg.Type() + CQSuffix
}

func (msg *MessageElementAnonymous) ProcessGRPC(segment *MessageSegmentGRPC) {
	segment.Data = &MessageSegmentGRPC_MessageElementAnonymous{
		MessageElementAnonymous: msg.ToGRPC(),
	}
}

func (msg *MessageElementAnonymous) ToGRPC() *MessageElementAnonymousGRPC {
	return &MessageElementAnonymousGRPC{
		Ignore: msg.Ignore,
	}
}

func (msg *MessageElementAnonymousGRPC) ToStruct() *MessageElementAnonymous {
	return &MessageElementAnonymous{
		Ignore: msg.Ignore,
	}
}

func init() {
	messageElementUnmarshalJSONMap["anonymous"] = func(data []byte) (MessageElement, error) {
		var result MessageElementAnonymous
		err := json.Unmarshal(data, &result)
		return &result, err
	}
	messageSegmentGRPCToStructMap["anonymous"] = func(msg *MessageSegmentGRPC) (MessageElement, error) {
		if msg.Data == nil {
			return nil, nil
		}
		var result MessageElementAnonymous
		switch data := msg.Data.(type) {
		case *MessageSegmentGRPC_MessageElementAnonymous:
			result = *data.MessageElementAnonymous.ToStruct()
		}
		return &result, nil
	}
}
