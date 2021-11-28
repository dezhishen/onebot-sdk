package model

import "encoding/json"

type LocationMessage struct {
	Lat     string `json:"lat"`
	Lon     string `json:"lon"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (msg LocationMessage) Type() string {
	return "location"
}

func init() {
	unmarshalJSONMap["location"] = func(data []byte) (MessageElement, error) {
		var result LocationMessage
		err := json.Unmarshal(data, &result)
		return result, err
	}
}
