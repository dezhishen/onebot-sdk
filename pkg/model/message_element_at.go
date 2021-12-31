package model

import "encoding/json"

type MessageElementAt struct {
	Qq string `json:"qq"`
}

func (msg *MessageElementAt) Type() string {
	return "at"
}

func (msg *MessageElementAt) ToGRPC() *MessageElementAtGRPC {
	var result MessageElementAtGRPC
	result.Qq = msg.Qq
	return &result
}

func (msg *MessageElementAtGRPC) ToStruct() *MessageElementAt {
	var result MessageElementAt
	result.Qq = msg.Qq
	return &result
}

func init() {
	messageElementUnmarshalJSONMap["at"] = func(data []byte) (MessageElement, error) {
		var result MessageElementAt
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
