package model

import "encoding/json"

// type MessageElementRps struct {
// }

func (msg MessageElementRps) Type() string {
	return "rps"
}

func init() {
	unmarshalJSONMap["rps"] = func(data []byte) (MessageElement, error) {
		var result MessageElementRps
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
