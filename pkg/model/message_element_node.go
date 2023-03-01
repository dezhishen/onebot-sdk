package model

import "encoding/json"

type MessageElementNode struct {
	// id	int32	转发消息id	直接引用他人的消息合并转发, 实际查看顺序为原消息发送顺序 与下面的自定义消息二选一
	Id string `json:"id,omitempty"`
	// name	string	发送者显示名字	用于自定义消息 (自定义消息并合并转发, 实际查看顺序为自定义消息段顺序)
	Name string `json:"name"`
	// uin	int64	发送者QQ号	用于自定义消息
	Uin int64 `json:"uin"`
	// content	message	具体消息	用于自定义消息 不支持转发套娃
	Content *MessageSegment `json:"content"`
	// seq	message	具体消息	用于自定义消息
	Seq *MessageSegment `json:"seq"`
}

func (msg *MessageElementNode) Type() string {
	return "node"
}

func (msg *MessageElementNode) Enabled() bool {
	return true
}

// ProcessGRPC
func (msg *MessageElementNode) ProcessGRPC(segment *MessageSegmentGRPC) {
	segment.Data = &MessageSegmentGRPC_MessageElementNode{
		MessageElementNode: msg.ToGRPC(),
	}
}

func (msg *MessageElementNode) ToGRPC() *MessageElementNodeGRPC {
	return &MessageElementNodeGRPC{
		Id:      msg.Id,
		Name:    msg.Name,
		Uin:     msg.Uin,
		Content: msg.Content.ToGRPC(),
		Seq:     msg.Seq.ToGRPC(),
	}
}

func (msg *MessageElementNodeGRPC) ToStruct() *MessageElementNode {
	return &MessageElementNode{
		Id:      msg.Id,
		Name:    msg.Name,
		Uin:     msg.Uin,
		Content: msg.Content.ToStruct(),
		Seq:     msg.Seq.ToStruct(),
	}
}

func init() {
	messageElementUnmarshalJSONMap["node"] = func(data []byte) (MessageElement, error) {
		var result MessageElementNode
		err := json.Unmarshal(data, &result)
		return &result, err
	}
	messageSegmentGRPCToStructMap["node"] = func(msg *MessageSegmentGRPC) (MessageElement, error) {
		if msg.Data == nil {
			return nil, nil
		}
		switch data := msg.Data.(type) {
		case *MessageSegmentGRPC_MessageElementNode:
			return data.MessageElementNode.ToStruct(), nil
		default:
			return nil, nil
		}
	}
}
