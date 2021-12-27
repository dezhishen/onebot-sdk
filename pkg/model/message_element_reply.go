package model

import "encoding/json"

type MessageElementReply struct {
	Id string `json:"id"`
}

func (msg MessageElementReply) Type() string {
	return "reply"
}

func init() {
	messageElementUnmarshalJSONMap["reply"] = func(data []byte) (MessageElement, error) {
		var result MessageElementReply
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
