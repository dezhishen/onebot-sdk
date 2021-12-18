package model

import "encoding/json"

type MessageElementMusic struct {
	MusicType string `json:"type"`
	Url       string `json:"url"`
	Audio     string `json:"audio"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Image     string `json:"image"`
}

func (msg MessageElementMusic) Type() string {
	return "music"
}

func init() {
	unmarshalJSONMap["music"] = func(data []byte) (MessageElement, error) {
		var result MessageElementMusic
		err := json.Unmarshal(data, &result)
		return result, err
	}
}
