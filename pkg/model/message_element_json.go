package model

import "encoding/json"

type JsonMessage struct {
	Data string `json:"data"`
}

func (msg JsonMessage) Type() string {
	return "json"
}

func init() {
	unmarshalJSONMap["json"] = func(data []byte) (MessageElement, error) {
		var result JsonMessage
		err := json.Unmarshal(data, &result)
		return result, err
	}
}
