package model

import "encoding/json"

type MessageElementForward struct {
	ID string `json:"id"`
}

func (msg MessageElementForward) Type() string {
	return "forward"
}

func (msg MessageElementForward) ToGRPC() *MessageElementForwardGRPC {
	return &MessageElementForwardGRPC{
		ID: msg.ID,
	}
}

func (msg *MessageElementForwardGRPC) ToStruct() *MessageElementForward {
	return &MessageElementForward{
		ID: msg.ID,
	}
}

func init() {
	messageElementUnmarshalJSONMap["forward"] = func(data []byte) (MessageElement, error) {
		var result MessageElementForward
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
