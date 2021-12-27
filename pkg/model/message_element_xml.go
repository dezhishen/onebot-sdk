package model

import "encoding/json"

type MessageElementXml struct {
	Data string `json:"data"`
}

func (msg MessageElementXml) Type() string {
	return "xml"
}

func (msg MessageElementXml) ToGRPC() *MessageElementXmlGRPC {
	return &MessageElementXmlGRPC{
		Data: msg.Data,
	}
}

func (msg *MessageElementXmlGRPC) ToStruct() *MessageElementXml {
	return &MessageElementXml{
		Data: msg.Data,
	}
}
func init() {
	messageElementUnmarshalJSONMap["xml"] = func(data []byte) (MessageElement, error) {
		var result MessageElementXml
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
