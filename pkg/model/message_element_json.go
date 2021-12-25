package model

import "encoding/json"

// type MessageElementJson struct {
// 	Data string `json:"data"`
// }

func (msg MessageElementJson) Type() string {
	return "json"
}

func init() {
	unmarshalJSONMap["json"] = func(data []byte) (MessageElement, error) {
		var result MessageElementJson
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
