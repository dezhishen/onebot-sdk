package model

import "encoding/json"

type MessageElementVideo struct {
	File    string `json:"file"`
	Url     string `json:"url"`
	Cache   uint32 `json:"cache"`
	Proxy   uint32 `json:"proxy"`
	Timeout uint32 `json:"timeout"`
}

func (msg MessageElementVideo) Type() string {
	return "video"
}

func (msg MessageElementVideo) ToGRPC() *MessageElementVideoGRPC {
	return &MessageElementVideoGRPC{
		File:    msg.File,
		Url:     msg.Url,
		Cache:   msg.Cache,
		Proxy:   msg.Proxy,
		Timeout: msg.Timeout,
	}
}

func (msg *MessageElementVideoGRPC) ToStruct() *MessageElementVideo {
	return &MessageElementVideo{
		File:    msg.File,
		Url:     msg.Url,
		Cache:   msg.Cache,
		Proxy:   msg.Proxy,
		Timeout: msg.Timeout,
	}
}
func init() {
	messageElementUnmarshalJSONMap["video"] = func(data []byte) (MessageElement, error) {
		var result MessageElementVideo
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
