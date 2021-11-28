package model

import "encoding/json"

type ForwardMessage struct {
	Id string `json:"id"`
}

func (msg ForwardMessage) Type() string {
	return "forward"
}

func init() {
	unmarshalJSONMap["forward"] = func(data []byte) (MessageElement, error) {
		var result ForwardMessage
		err := json.Unmarshal(data, &result)
		return result, err
	}
}
