package model

import "encoding/json"

type ShareMessage struct {
	Url     string `json:"url"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
}

func (msg ShareMessage) Type() string {
	return "share"
}

func init() {
	unmarshalJSONMap["share"] = func(data []byte) (MessageElement, error) {
		var result ShareMessage
		err := json.Unmarshal(data, &result)
		return result, err
	}
}
