package model

import "encoding/json"

type MessageElementVideo struct {
	File    string `json:"file"`
	Url     string `json:"url"`
	Cache   uint   `json:"cache"`
	Proxy   uint   `json:"proxy"`
	Timeout uint   `json:"timeout"`
}

func (msg MessageElementVideo) Type() string {
	return "video"
}

func init() {
	unmarshalJSONMap["video"] = func(data []byte) (MessageElement, error) {
		var result MessageElementVideo
		err := json.Unmarshal(data, &result)
		return result, err
	}
}
