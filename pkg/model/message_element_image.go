package model

import "encoding/json"

type ImageMessage struct {
	File string `json:"file"`
	//图片类型
	ImageType string `json:"type"`
	Url       string `json:"url"`
	Cache     uint   `json:"cache"`
	Proxy     uint   `json:"proxy"`
	Timeout   uint   `json:"timeout"`
}

func (msg ImageMessage) Type() string {
	return "image"
}

func init() {
	unmarshalJSONMap["image"] = func(data []byte) (MessageElement, error) {
		var result ImageMessage
		err := json.Unmarshal(data, &result)
		return result, err
	}
}
