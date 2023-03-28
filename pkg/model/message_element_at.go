package model

import (
	"encoding/json"
)

type MessageElementAt struct {
	Qq   string `json:"qq"`
	Name string `json:"name"`
}

func (msg *MessageElementAt) Type() string {
	return "at"
}

func (msg *MessageElementAt) Enabled() bool {
	return true
}

// ToCQCode
func (msg *MessageElementAt) ToCQCode() string {
	cqCode := CQPrefix + msg.Type() + CQEleSplit + "qq=" + msg.Qq
	if msg.Name != "" {
		cqCode += (CQEleSplit + "name=" + msg.Name)
	}
	return cqCode + CQSuffix
}

// ProcessGRPC
func (msg *MessageElementAt) ProcessGRPC(segment *MessageSegmentGRPC) {
	segment.Data = &MessageSegmentGRPC_MessageElementAt{
		MessageElementAt: msg.ToGRPC(),
	}
}

func (msg *MessageElementAt) ToGRPC() *MessageElementAtGRPC {
	var result MessageElementAtGRPC
	result.Qq = msg.Qq
	result.Name = msg.Name
	return &result
}

func (msg *MessageElementAtGRPC) ToStruct() *MessageElementAt {
	var result MessageElementAt
	result.Qq = msg.Qq
	result.Name = msg.Name
	return &result
}

func init() {
	messageElementUnmarshalJSONMap["at"] = func(data []byte) (MessageElement, error) {
		var result MessageElementAt
		err := json.Unmarshal(data, &result)
		return &result, err
	}
	messageSegmentGRPCToStructMap["at"] = func(msg *MessageSegmentGRPC) (MessageElement, error) {
		if msg.Data == nil {
			return nil, nil
		}
		switch data := msg.Data.(type) {
		case *MessageSegmentGRPC_MessageElementAt:
			return data.MessageElementAt.ToStruct(), nil
		default:
			return nil, nil
		}
	}
}
