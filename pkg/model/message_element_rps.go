package model

import "encoding/json"

type RpsMessage struct {
}

func (msg RpsMessage) Type() string {
	return "rps"
}

func init() {
	unmarshalJSONMap["rps"] = func(data []byte) (MessageElement, error) {
		var result RpsMessage
		err := json.Unmarshal(data, &result)
		return result, err
	}
}
