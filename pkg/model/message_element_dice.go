package model

import "encoding/json"

//掷骰子魔法表情
type MessageElementDice struct {
}

//掷骰子魔法表情
func (msg MessageElementDice) Type() string {
	return "dice"
}

func init() {
	//掷骰子魔法表情
	unmarshalJSONMap["dice"] = func(data []byte) (MessageElement, error) {
		var result MessageElementDice
		err := json.Unmarshal(data, &result)
		return result, err
	}
}
