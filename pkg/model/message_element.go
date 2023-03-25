package model

import (
	"encoding/json"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

type Message []*MessageSegment

func (m *Message) UnmarshalJSON(data []byte) error {
	if data[0] == '"' && data[len(data)-1] == '"' {
		code, err := ParseStringMessage(strings.Trim(string(data), "\""))
		if err != nil {
			return err
		}
		*m = code
	} else {
		var msgSeg []*MessageSegment
		err := json.Unmarshal(data, &msgSeg)
		if err != nil {
			return err
		}
	}
	return nil
}

type MessageSegment struct {
	Type string         `json:"type"`
	Data MessageElement `json:"data"`
}

func (msg *MessageSegment) ToGRPC() *MessageSegmentGRPC {
	if msg == nil {
		return nil
	}
	result := &MessageSegmentGRPC{
		Type: msg.Type,
	}
	if msg.Data != nil {
		msg.Data.ProcessGRPC(result)
	}
	return result
}

var messageSegmentGRPCToStructMap = make(map[string]func(msg *MessageSegmentGRPC) (MessageElement, error))

func (msg *MessageSegmentGRPC) ToStruct() *MessageSegment {
	if msg == nil {
		return nil
	}
	result := &MessageSegment{
		Type: msg.Type,
	}
	if _, ok := messageSegmentGRPCToStructMap[msg.Type]; !ok {
		log.Errorf("未定义的消息类型:%v", msg.Type)
		return nil
	}
	data, err := messageSegmentGRPCToStructMap[msg.Type](msg)
	if err != nil {
		log.Errorf("消息类型转换失败:%v", err)
		return nil
	}
	result.Data = data
	return result
}

func MessageSegmentGRPCArray2MessageSegmentArray(source []*MessageSegmentGRPC) []*MessageSegment {
	if len(source) == 0 {
		return nil
	}
	var result []*MessageSegment
	for _, e := range source {
		result = append(result, e.ToStruct())
	}
	return result
}

func MessageSegmentArray2MessageSegmentGRPCArray(source []*MessageSegment) []*MessageSegmentGRPC {
	if len(source) == 0 {
		return nil
	}
	var result []*MessageSegmentGRPC
	for _, e := range source {
		result = append(result, e.ToGRPC())
	}
	return result
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
	if err != nil {
		return fmt.Errorf("json转换失败,%v", err)
	}
	msgSeg.Type = messageType
	msgSeg.Data = element
	return err
}

type MessageElement interface {
	Type() string
	Enabled() bool
	ProcessGRPC(msg *MessageSegmentGRPC)
}
