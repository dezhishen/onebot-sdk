package model

import (
	"encoding/json"
	"fmt"
)

type MessageElementPoke struct {
	//qq	int64	需要戳的成员
	Qq int64 `json:"qq"`
}

func (msg *MessageElementPoke) Type() string {
	return "poke"
}

func (msg *MessageElementPoke) Enabled() bool {
	return true
}

func (msg *MessageElementPoke) ToCQCode() string {
	return CQPrefix + msg.Type() + CQEleSplit + "qq=" + fmt.Sprintf("%d", msg.Qq) + CQSuffix
}

// ProcessGRPC
func (msg *MessageElementPoke) ProcessGRPC(segment *MessageSegmentGRPC) {
	segment.Data = &MessageSegmentGRPC_MessageElementPoke{
		MessageElementPoke: msg.ToGRPC(),
	}
}

func (msg *MessageElementPoke) ToGRPC() *MessageElementPokeGRPC {
	return &MessageElementPokeGRPC{
		Qq: msg.Qq,
	}
}

func (msg *MessageElementPokeGRPC) ToStruct() *MessageElementPoke {
	return &MessageElementPoke{
		Qq: msg.Qq,
	}
}

func init() {
	messageElementUnmarshalJSONMap["poke"] = func(data []byte) (MessageElement, error) {
		var result MessageElementPoke
		err := json.Unmarshal(data, &result)
		return &result, err
	}
	messageSegmentGRPCToStructMap["poke"] = func(msg *MessageSegmentGRPC) (MessageElement, error) {
		if msg.Data == nil {
			return nil, nil
		}
		switch data := msg.Data.(type) {
		case *MessageSegmentGRPC_MessageElementPoke:
			return data.MessageElementPoke.ToStruct(), nil
		default:
			return nil, nil
		}
	}
}
