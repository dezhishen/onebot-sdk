package model

import "encoding/json"

type RecordMessage struct {
	File string `json:"file"`
}

func (msg RecordMessage) Type() string {
	return "record"
}

func init() {
	unmarshalJSONMap["record"] = func(data []byte) (MessageElement, error) {
		var result RecordMessage
		err := json.Unmarshal(data, &result)
		return result, err
	}
}
