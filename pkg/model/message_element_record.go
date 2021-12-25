package model

import "encoding/json"

// type MessageElementRecord struct {
// 	File string `json:"file"`
// }

func (msg MessageElementRecord) Type() string {
	return "record"
}

func init() {
	unmarshalJSONMap["record"] = func(data []byte) (MessageElement, error) {
		var result MessageElementRecord
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
