package model

import "encoding/json"

type MessageElementNode struct {
	Id      int32             `json:"id"`
	Name    string            `json:"name"`
	Uin     int64             `json:"uin"`
	Content []*MessageSegment `json:"content"`
}

func (msg *MessageElementNode) Type() string {
	return "node"
}

func (msg *MessageElementNodeGRPC) ToStruct() *MessageElementNode {
	return &MessageElementNode{
		Id:      msg.Id,
		Name:    msg.Name,
		Uin:     msg.Uin,
		Content: MessageSegmentGRPCArray2MessageSegmentArray(msg.Content),
	}
}

func (msg *MessageElementNode) ToGRPC() *MessageElementNodeGRPC {
	return &MessageElementNodeGRPC{
		Id:      msg.Id,
		Name:    msg.Name,
		Uin:     msg.Uin,
		Content: MessageSegmentArray2MessageSegmentGRPCArray(msg.Content),
	}
}

func init() {
	messageElementUnmarshalJSONMap["node"] = func(data []byte) (MessageElement, error) {
		var result MessageElementNode
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
