package model

import "encoding/json"

type MessageElementJson struct {
	Data string `json:"data"`
}

func (msg *MessageElementJson) Type() string {
	return "json"
}

func (msg *MessageElementJson) ToGRPC() *MessageElementJsonGRPC {
	return &MessageElementJsonGRPC{
		Data: msg.Data,
	}
}

func (msg *MessageElementJsonGRPC) ToStruct() *MessageElementJson {
	return &MessageElementJson{
		Data: msg.Data,
	}
}

func init() {
	messageElementUnmarshalJSONMap["json"] = func(data []byte) (MessageElement, error) {
		var result MessageElementJson
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
