package model

import "encoding/json"

type MessageElementShare struct {
	Url     string `json:"url"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
}

func (msg MessageElementShare) Type() string {
	return "share"
}

func init() {
	messageElementUnmarshalJSONMap["share"] = func(data []byte) (MessageElement, error) {
		var result MessageElementShare
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
