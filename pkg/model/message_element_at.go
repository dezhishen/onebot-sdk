package model

import "encoding/json"

type MessageElementAt struct {
	Qq string `json:"qq"`
}

func (msg MessageElementAt) Type() string {
	return "at"
}

func init() {
	messageElementUnmarshalJSONMap["at"] = func(data []byte) (MessageElement, error) {
		var result MessageElementAt
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
