package model

import "encoding/json"

type MessageElementReply struct {
	Id string `json:"id"`
}

func (msg *MessageElementReply) Type() string {
	return "reply"
}

func (msg *MessageElementReply) ToGRPC() *MessageElementReplyGRPC {
	return &MessageElementReplyGRPC{
		Id: msg.Id,
	}
}

func (msg *MessageElementReplyGRPC) ToStruct() *MessageElementReply {
	return &MessageElementReply{
		Id: msg.Id,
	}
}

func init() {
	messageElementUnmarshalJSONMap["reply"] = func(data []byte) (MessageElement, error) {
		var result MessageElementReply
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
