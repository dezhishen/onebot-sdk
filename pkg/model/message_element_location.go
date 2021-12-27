package model

import "encoding/json"

type MessageElementLocation struct {
	Lat     string `json:"lat"`
	Lon     string `json:"lon"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (msg MessageElementLocation) Type() string {
	return "location"
}

func init() {
	messageElementUnmarshalJSONMap["location"] = func(data []byte) (MessageElement, error) {
		var result MessageElementLocation
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
