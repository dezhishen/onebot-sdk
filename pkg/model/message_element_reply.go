package model

import "encoding/json"

type MessageElementReply struct {
	// id	int	回复时所引用的消息id, 必须为本群消息.
	Id int64 `json:"id"`
	// 	text	string	自定义回复的信息
	Text string `json:"text"`
	// qq	int64	自定义回复时的自定义QQ, 如果使用自定义信息必须指定.
	QQ int64 `json:"qq"`
	// time	int64	自定义回复时的时间, 格式为Unix时间
	Time int64 `json:"time"`
	// seq	int64	起始消息序号, 可通过 get_msg 获得
	Seq int64 `json:"seq"`
}

func (msg *MessageElementReply) Enabled() bool {
	return true
}

func (msg *MessageElementReply) Type() string {
	return "reply"
}

// ProcessGRPC
func (msg *MessageElementReply) ProcessGRPC(segment *MessageSegmentGRPC) {
	segment.Data = &MessageSegmentGRPC_MessageElementReply{
		MessageElementReply: msg.ToGRPC(),
	}
}

func (msg *MessageElementReply) ToGRPC() *MessageElementReplyGRPC {
	return &MessageElementReplyGRPC{
		Id:   msg.Id,
		Text: msg.Text,
		Qq:   msg.QQ,
		Time: msg.Time,
		Seq:  msg.Seq,
	}
}

func (msg *MessageElementReplyGRPC) ToStruct() *MessageElementReply {
	return &MessageElementReply{
		Id:   msg.Id,
		Text: msg.Text,
		QQ:   msg.Qq,
		Time: msg.Time,
		Seq:  msg.Seq,
	}
}

func init() {
	messageElementUnmarshalJSONMap["reply"] = func(data []byte) (MessageElement, error) {
		var result MessageElementReply
		err := json.Unmarshal(data, &result)
		return &result, err
	}
	messageSegmentGRPCToStructMap["reply"] = func(msg *MessageSegmentGRPC) (MessageElement, error) {
		if msg.Data == nil {
			return nil, nil
		}
		switch data := msg.Data.(type) {
		case *MessageSegmentGRPC_MessageElementReply:
			return data.MessageElementReply.ToStruct(), nil
		default:
			return nil, nil
		}
	}
}
