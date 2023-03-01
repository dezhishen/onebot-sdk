package model

import "encoding/json"

type MessageElementMusic struct {
	// qq 163 xm / custom 分别表示使用 QQ 音乐、网易云音乐、虾米音乐
	MusicType string `json:"type"`
	// 歌曲 ID qq 163 xm 时支持
	Id string `json:"id"`
	// 歌曲 类型为 custom 时支持
	// 点击后跳转目标 URL
	Url string `json:"url"`
	// 歌曲 类型为 custom 时支持
	// 音乐 URL
	Audio string `json:"audio"`
	// 歌曲 类型为 custom 时支持
	// 标题
	Title string `json:"title"`
	// 歌曲 类型为 custom 时支持
	// 发送时可选, 内容描述
	Content string `json:"content"`
	// 歌曲 类型为 custom 时支持
	// 发送时可选, 图片 URL
	Image string `json:"image"`
}

func (msg *MessageElementMusic) Type() string {
	return "music"
}

func (msg *MessageElementMusic) Enabled() bool {
	return true
}

// ProcessGRPC
func (msg *MessageElementMusic) ProcessGRPC(segment *MessageSegmentGRPC) {
	segment.Data = &MessageSegmentGRPC_MessageElementMusic{
		MessageElementMusic: msg.ToGRPC(),
	}
}

func (msg *MessageElementMusic) ToGRPC() *MessageElementMusicGRPC {
	return &MessageElementMusicGRPC{
		MusicType: msg.MusicType,
		Id:        msg.Id,
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
		Id:        msg.Id,
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
	messageSegmentGRPCToStructMap["music"] = func(msg *MessageSegmentGRPC) (MessageElement, error) {
		if msg.Data == nil {
			return nil, nil
		}
		switch data := msg.Data.(type) {
		case *MessageSegmentGRPC_MessageElementMusic:
			return data.MessageElementMusic.ToStruct(), nil
		default:
			return nil, nil
		}
	}
}
