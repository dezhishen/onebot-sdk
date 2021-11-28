package model

import "encoding/json"

type AtMessage struct {
	Qq string `json:"qq"`
}

func (msg AtMessage) Type() string {
	return "at"
}

func init() {
	unmarshalJSONMap["at"] = func(data []byte) (MessageElement, error) {
		var result AtMessage
		err := json.Unmarshal(data, &result)
		return result, err
	}
}
