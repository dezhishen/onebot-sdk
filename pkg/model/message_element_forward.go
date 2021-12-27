package model

import "encoding/json"

type MessageElementForward struct {
	Id string `json:"id"`
}

func (msg MessageElementForward) Type() string {
	return "forward"
}

func init() {
	messageElementUnmarshalJSONMap["forward"] = func(data []byte) (MessageElement, error) {
		var result MessageElementForward
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
