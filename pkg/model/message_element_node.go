package model

import "encoding/json"

type MessageElementNode struct {
	UserId   string            `json:"user_id"`
	Nickname string            `json:"nickname"`
	Content  []*MessageSegment `json:"content"`
}

func (msg *MessageElementNode) Type() string {
	return "node"
}

func (msg *MessageElementNodeGRPC) ToStruct() *MessageElementNode {
	return &MessageElementNode{
		UserId:   msg.UserId,
		Nickname: msg.Nickname,
		Content:  MessageSegmentGRPCArray2MessageSegmentArray(msg.Content),
	}
}

func (msg *MessageElementNode) ToGRPC() *MessageElementNodeGRPC {
	return &MessageElementNodeGRPC{
		UserId:   msg.UserId,
		Nickname: msg.Nickname,
		Content:  MessageSegmentArray2MessageSegmentGRPCArray(msg.Content),
	}
}

func init() {
	messageElementUnmarshalJSONMap["node"] = func(data []byte) (MessageElement, error) {
		var result MessageElementNode
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
