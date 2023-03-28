package model

import "encoding/json"

//掷骰子魔法表情
type MessageElementDice struct {
}

//掷骰子魔法表情
func (msg *MessageElementDice) Type() string {
	return "dice"
}

func (msg *MessageElementDice) Enabled() bool {
	return false
}
func (msg *MessageElementDice) ToCQCode() string {
	return CQPrefix + msg.Type() + CQSuffix
}

// ProcessGRPC
func (msg *MessageElementDice) ProcessGRPC(segment *MessageSegmentGRPC) {
	segment.Data = &MessageSegmentGRPC_MessageElementDice{
		MessageElementDice: msg.ToGRPC(),
	}
}

func (msg *MessageElementDice) ToGRPC() *MessageElementDiceGRPC {
	var result MessageElementDiceGRPC
	return &result
}

func (msg *MessageElementDiceGRPC) ToStruct() *MessageElementDice {
	var result MessageElementDice
	return &result
}

func init() {
	//掷骰子魔法表情
	messageElementUnmarshalJSONMap["dice"] = func(data []byte) (MessageElement, error) {
		var result MessageElementDice
		err := json.Unmarshal(data, &result)
		return &result, err
	}
	messageSegmentGRPCToStructMap["dice"] = func(msg *MessageSegmentGRPC) (MessageElement, error) {
		if msg.Data == nil {
			return nil, nil
		}
		switch data := msg.Data.(type) {
		case *MessageSegmentGRPC_MessageElementDice:
			return data.MessageElementDice.ToStruct(), nil
		default:
			return nil, nil
		}
	}

}
