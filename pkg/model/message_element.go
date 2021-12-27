package model

import (
	"fmt"

	"github.com/tidwall/gjson"
)

type MessageSegment struct {
	Type string         `json:"type"`
	Data MessageElement `json:"data"`
}

var messageElementUnmarshalJSONMap = make(map[string]func(data []byte) (MessageElement, error))

func (msgSeg *MessageSegment) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || data[0] == 'n' { // copied from the Q, can be improved
		return nil
	}
	messageType := gjson.GetBytes(data, "type").Str
	decoder, ok := messageElementUnmarshalJSONMap[messageType]
	if !ok {
		return fmt.Errorf("未找到指定的消息类型,%v", messageType)
	}
	element, err := decoder([]byte(gjson.GetBytes(data, "data").Raw))
	if !ok {
		return fmt.Errorf("未找到指定的消息类型,%v", messageType)
	}
	msgSeg.Type = messageType
	msgSeg.Data = element
	return err
}

type MessageElement interface {
	Type() string
}
