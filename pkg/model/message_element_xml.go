package model

import "encoding/json"

type MessageElementXml struct {
	Data string `json:"data"`
}

func (msg MessageElementXml) Type() string {
	return "xml"
}

func init() {
	unmarshalJSONMap["xml"] = func(data []byte) (MessageElement, error) {
		var result MessageElementXml
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
