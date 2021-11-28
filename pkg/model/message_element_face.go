package model

import "encoding/json"

type FaceMessage struct {
	ID string `json:"id"`
}

func (msg FaceMessage) Type() string {
	return "face"
}

func init() {
	unmarshalJSONMap["face"] = func(data []byte) (MessageElement, error) {
		var result FaceMessage
		err := json.Unmarshal(data, &result)
		return result, err
	}
}
