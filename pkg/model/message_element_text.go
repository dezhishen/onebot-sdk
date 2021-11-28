package model

import "encoding/json"

type TextMessage struct {
	Text string `json:"text"`
}

func (msg TextMessage) Type() string {
	return "text"
}

func init() {
	unmarshalJSONMap["text"] = func(data []byte) (MessageElement, error) {
		var result TextMessage
		err := json.Unmarshal(data, &result)
		return result, err
	}
}
