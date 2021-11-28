package model

import "encoding/json"

type XmlMessage struct {
	Data string `json:"data"`
}

func (msg XmlMessage) Type() string {
	return "xml"
}

func init() {
	unmarshalJSONMap["xml"] = func(data []byte) (MessageElement, error) {
		var result XmlMessage
		err := json.Unmarshal(data, &result)
		return result, err
	}
}
