package model

import "encoding/json"

type MusicMessage struct {
	MusicType string `json:"type"`
	Url       string `json:"url"`
	Audio     string `json:"audio"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Image     string `json:"image"`
}

func (msg MusicMessage) Type() string {
	return "music"
}

func init() {
	unmarshalJSONMap["music"] = func(data []byte) (MessageElement, error) {
		var result MusicMessage
		err := json.Unmarshal(data, &result)
		return result, err
	}
}
