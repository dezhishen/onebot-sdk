package model

import "encoding/json"

//掷骰子魔法表情
type MessageElementDice struct {
}

//掷骰子魔法表情
func (msg MessageElementDice) Type() string {
	return "dice"
}

func (msg MessageElementDice) ToGRPC() *MessageElementDiceGRPC {
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
}
