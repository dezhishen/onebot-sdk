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

func (msg *MessageElementMusic) Type() string {
	return "music"
}

func (msg *MessageElementMusic) ToGRPC() *MessageElementMusicGRPC {
	return &MessageElementMusicGRPC{
		MusicType: msg.MusicType,
		Url:       msg.Url,
		Audio:     msg.Audio,
		Title:     msg.Title,
		Content:   msg.Content,
		Image:     msg.Image,
	}
}

func (msg *MessageElementMusicGRPC) ToStruct() *MessageElementMusic {
	return &MessageElementMusic{
		MusicType: msg.MusicType,
		Url:       msg.Url,
		Audio:     msg.Audio,
		Title:     msg.Title,
		Content:   msg.Content,
		Image:     msg.Image,
	}
}

func init() {
	messageElementUnmarshalJSONMap["music"] = func(data []byte) (MessageElement, error) {
		var result MessageElementMusic
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
