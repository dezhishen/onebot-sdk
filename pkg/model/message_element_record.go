package model

import "encoding/json"

type MessageElementRecord struct {
	File string `json:"file"`
}

func (msg MessageElementRecord) Type() string {
	return "record"
}

func (msg *MessageElementRecord) ToGRPC() *MessageElementRecordGRPC {
	return &MessageElementRecordGRPC{
		File: msg.File,
	}
}

func (msg *MessageElementRecordGRPC) ToStruct() *MessageElementRecord {
	return &MessageElementRecord{
		File: msg.File,
	}
}

func init() {
	messageElementUnmarshalJSONMap["record"] = func(data []byte) (MessageElement, error) {
		var result MessageElementRecord
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
