package model

import "encoding/json"

type VideoMessage struct {
	File    string `json:"file"`
	Url     string `json:"url"`
	Cache   uint   `json:"cache"`
	Proxy   uint   `json:"proxy"`
	Timeout uint   `json:"timeout"`
}

func (msg VideoMessage) Type() string {
	return "video"
}

func init() {
	unmarshalJSONMap["video"] = func(data []byte) (MessageElement, error) {
		var result VideoMessage
		err := json.Unmarshal(data, &result)
		return result, err
	}
}
