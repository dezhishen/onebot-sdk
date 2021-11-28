package model

import "encoding/json"

type ContactMessage struct {
	ContactType string `json:"type"`
	Id          string `json:"id"`
}

func (msg ContactMessage) Type() string {
	return "contact"
}

func init() {
	unmarshalJSONMap["contact"] = func(data []byte) (MessageElement, error) {
		var result ContactMessage
		err := json.Unmarshal(data, &result)
		return result, err
	}
}
