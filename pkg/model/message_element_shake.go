package model

import "encoding/json"

// type MessageELementShake struct {
// }

func (msg MessageELementShake) Type() string {
	return "shake"
}

func init() {
	unmarshalJSONMap["shake"] = func(data []byte) (MessageElement, error) {
		var result MessageELementShake
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
