package model

import "encoding/json"

type ShakeMessage struct {
}

func (msg ShakeMessage) Type() string {
	return "shake"
}

func init() {
	unmarshalJSONMap["shake"] = func(data []byte) (MessageElement, error) {
		var result ShakeMessage
		err := json.Unmarshal(data, &result)
		return result, err
	}
}
