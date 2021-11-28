package model

import "encoding/json"

type NodeMessage struct {
	User_id  string           `json:"user_id"`
	Nickname string           `json:"nickname"`
	Content  []MessageSegment `json:"content"`
}

func (msg NodeMessage) Type() string {
	return "node"
}

func init() {
	unmarshalJSONMap["node"] = func(data []byte) (MessageElement, error) {
		var result NodeMessage
		err := json.Unmarshal(data, &result)
		return result, err
	}
}
