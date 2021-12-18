package model

import "encoding/json"

type MessageElementImage struct {
	File string `json:"file"`
	//图片类型
	ImageType string `json:"type"`
	Url       string `json:"url"`
	Cache     uint   `json:"cache"`
	Proxy     uint   `json:"proxy"`
	Timeout   uint   `json:"timeout"`
}

func (msg MessageElementImage) Type() string {
	return "image"
}

func init() {
	unmarshalJSONMap["image"] = func(data []byte) (MessageElement, error) {
		var result MessageElementImage
		err := json.Unmarshal(data, &result)
		return result, err
	}
}
