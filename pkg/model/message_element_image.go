package model

import "encoding/json"

type MessageElementImage struct {
	File string `json:"file"`
	//图片类型
	ImageType string `json:"type"`
	Url       string `json:"url"`
	Cache     uint32 `json:"cache"`
	Proxy     uint32 `json:"proxy"`
	Timeout   uint32 `json:"timeout"`
}

func (msg MessageElementImage) Type() string {
	return "image"
}

func (msg MessageElementImage) ToGRPC() *MessageElementImageGRPC {
	return &MessageElementImageGRPC{
		File:      msg.File,
		ImageType: msg.ImageType,
		Url:       msg.Url,
		Cache:     msg.Cache,
		Proxy:     msg.Proxy,
		Timeout:   msg.Timeout,
	}
}

func (msg *MessageElementImageGRPC) ToStruct() *MessageElementImage {
	return &MessageElementImage{
		File:      msg.File,
		ImageType: msg.ImageType,
		Url:       msg.Url,
		Cache:     msg.Cache,
		Proxy:     msg.Proxy,
		Timeout:   msg.Timeout,
	}
}

func init() {
	messageElementUnmarshalJSONMap["image"] = func(data []byte) (MessageElement, error) {
		var result MessageElementImage
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
