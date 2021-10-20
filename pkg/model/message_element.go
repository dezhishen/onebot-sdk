package model

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
)

type MessageSegment struct {
	Type string         `json:"type"`
	Data MessageElement `json:"data"`
}

var unmarshalJSONMap = make(map[string]func(data []byte) (MessageElement, error))

func (msgSeg *MessageSegment) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || data[0] == 'n' { // copied from the Q, can be improved
		return nil
	}
	messageType := gjson.GetBytes(data, "type").Str
	decoder, ok := unmarshalJSONMap[messageType]
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

type FaceMessage struct {
	ID string `json:"id"`
}

func (msg FaceMessage) Type() string {
	return "face"
}

type TextMessage struct {
	Text string `json:"text"`
}

func (msg TextMessage) Type() string {
	return "text"
}

func init() {
	unmarshalJSONMap["face"] = func(data []byte) (MessageElement, error) {
		var result FaceMessage
		err := json.Unmarshal(data, &result)
		return result, err
	}
	unmarshalJSONMap["text"] = func(data []byte) (MessageElement, error) {
		var result TextMessage
		err := json.Unmarshal(data, &result)
		return result, err
	}

}
