package model

import "encoding/json"

type MessageElementFace struct {
	Id string `json:"id"`
}

func (msg MessageElementFace) Type() string {
	return "face"
}

func (msg MessageElementFace) ToGRPC() *MessageElementFaceGRPC {
	return &MessageElementFaceGRPC{
		Id: msg.Id,
	}
}

func (msg *MessageElementFaceGRPC) ToStruct() *MessageElementFace {
	return &MessageElementFace{
		Id: msg.Id,
	}
}

func init() {
	messageElementUnmarshalJSONMap["face"] = func(data []byte) (MessageElement, error) {
		var result MessageElementFace
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
