package model

import "encoding/json"

type ReplyMessage struct {
	Id string `json:"id"`
}

func (msg ReplyMessage) Type() string {
	return "reply"
}

func init() {
	unmarshalJSONMap["reply"] = func(data []byte) (MessageElement, error) {
		var result ReplyMessage
		err := json.Unmarshal(data, &result)
		return result, err
	}
}
