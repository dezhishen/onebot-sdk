package model

import "encoding/json"

type MessageElementForward struct {
	// id	string	合并转发ID, 需要通过 /get_forward_msg API获取转发的具体内容
	Id string `json:"id"`
}

func (msg *MessageElementForward) Type() string {
	return "forward"
}

func (msg *MessageElementForward) Enabled() bool {
	return true
}

// ProcessGRPC
func (msg *MessageElementForward) ProcessGRPC(segment *MessageSegmentGRPC) {
	segment.Data = &MessageSegmentGRPC_MessageElementForward{
		MessageElementForward: msg.ToGRPC(),
	}
}

func (msg *MessageElementForward) ToGRPC() *MessageElementForwardGRPC {
	return &MessageElementForwardGRPC{
		Id: msg.Id,
	}
}

func (msg *MessageElementForwardGRPC) ToStruct() *MessageElementForward {
	return &MessageElementForward{
		Id: msg.Id,
	}
}

func init() {
	messageElementUnmarshalJSONMap["forward"] = func(data []byte) (MessageElement, error) {
		var result MessageElementForward
		err := json.Unmarshal(data, &result)
		return &result, err
	}
	messageSegmentGRPCToStructMap["forward"] = func(msg *MessageSegmentGRPC) (MessageElement, error) {
		if msg.Data == nil {
			return nil, nil
		}
		switch data := msg.Data.(type) {
		case *MessageSegmentGRPC_MessageElementForward:
			return data.MessageElementForward.ToStruct(), nil
		default:
			return nil, nil
		}
	}
}
