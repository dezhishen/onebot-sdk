package model

import "encoding/json"

type MessageElementShake struct {
}

func (msg MessageElementShake) Type() string {
	return "shake"
}

func init() {
	messageElementUnmarshalJSONMap["shake"] = func(data []byte) (MessageElement, error) {
		var result MessageElementShake
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
