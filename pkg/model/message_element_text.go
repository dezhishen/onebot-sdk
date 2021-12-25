package model

import "encoding/json"

// type MessageElementText struct {
// 	Text string `json:"text"`
// }

func (msg MessageElementText) Type() string {
	return "text"
}

func init() {
	unmarshalJSONMap["text"] = func(data []byte) (MessageElement, error) {
		var result MessageElementText
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
