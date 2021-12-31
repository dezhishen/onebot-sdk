package model

import "encoding/json"

type MessageElementLocation struct {
	Lat     string `json:"lat"`
	Lon     string `json:"lon"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (msg *MessageElementLocation) Type() string {
	return "location"
}

func (msg *MessageElementLocation) ToGRPC() *MessageElementLocationGRPC {
	return &MessageElementLocationGRPC{
		Lat:     msg.Lat,
		Lon:     msg.Lon,
		Title:   msg.Title,
		Content: msg.Content,
	}
}

func (msg *MessageElementLocationGRPC) ToStruct() *MessageElementLocation {
	return &MessageElementLocation{
		Lat:     msg.Lat,
		Lon:     msg.Lon,
		Title:   msg.Title,
		Content: msg.Content,
	}
}

func init() {
	messageElementUnmarshalJSONMap["location"] = func(data []byte) (MessageElement, error) {
		var result MessageElementLocation
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
