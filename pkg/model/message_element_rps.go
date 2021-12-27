package model

import "encoding/json"

type MessageElementRps struct {
}

func (msg MessageElementRps) Type() string {
	return "rps"
}

func (msg MessageElementRps) ToGRPC() *MessageElementRpsGRPC {
	var result MessageElementRpsGRPC
	return &result
}

func (msg *MessageElementRpsGRPC) ToStruct() *MessageElementRps {
	var result MessageElementRps
	return &result
}

func init() {
	messageElementUnmarshalJSONMap["rps"] = func(data []byte) (MessageElement, error) {
		var result MessageElementRps
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
