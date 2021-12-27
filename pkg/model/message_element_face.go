package model

import "encoding/json"

type MessageElementFace struct {
	ID string `json:"id"`
}

func (msg MessageElementFace) Type() string {
	return "face"
}

func (msg MessageElementFace) ToGRPC() *MessageElementFaceGRPC {
	return &MessageElementFaceGRPC{
		ID: msg.ID,
	}
}

func (msg *MessageElementFaceGRPC) ToStruct() *MessageElementFace {
	return &MessageElementFace{
		ID: msg.ID,
	}
}

func init() {
	messageElementUnmarshalJSONMap["face"] = func(data []byte) (MessageElement, error) {
		var result MessageElementFace
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
