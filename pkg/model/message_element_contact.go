package model

import "encoding/json"

type MessageElementContact struct {
	ContactType string `json:"type"`
	Id          string `json:"id"`
}

func (msg *MessageElementContact) Type() string {
	return "contact"
}

func (msg *MessageElementContact) ToGRPC() *MessageElementContactGRPC {
	var result MessageElementContactGRPC
	result.ContactType = msg.ContactType
	result.Id = msg.Id
	return &result
}

func (msg *MessageElementContactGRPC) ToStruct() *MessageElementContact {
	var result MessageElementContact
	result.ContactType = msg.ContactType
	result.Id = msg.Id
	return &result
}

func init() {
	messageElementUnmarshalJSONMap["contact"] = func(data []byte) (MessageElement, error) {
		var result MessageElementContact
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
