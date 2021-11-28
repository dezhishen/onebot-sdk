package model

import "encoding/json"

//掷骰子魔法表情
type DiceMessage struct {
}

//掷骰子魔法表情
func (msg DiceMessage) Type() string {
	return "dice"
}

func init() {
	//掷骰子魔法表情
	unmarshalJSONMap["dice"] = func(data []byte) (MessageElement, error) {
		var result DiceMessage
		err := json.Unmarshal(data, &result)
		return result, err
	}
}
