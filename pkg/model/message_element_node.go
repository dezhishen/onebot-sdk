package model

import "encoding/json"

type MessageElementNode struct {
	User_id  string           `json:"user_id"`
	Nickname string           `json:"nickname"`
	Content  []MessageElement `json:"content"`
}

func (msg MessageElementNode) Type() string {
	return "node"
}

func init() {
	unmarshalJSONMap["node"] = func(data []byte) (MessageElement, error) {
		var result MessageElementNode
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
