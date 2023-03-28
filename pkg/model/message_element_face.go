package model

import "encoding/json"

type MessageElementFace struct {
	Id string `json:"id"`
}

func (msg *MessageElementFace) Type() string {
	return "face"
}

func (msg *MessageElementFace) Enabled() bool {
	return true
}

func (msg *MessageElementFace) ToCQCode() string {
	return CQPrefix + msg.Type() + CQEleSplit + "id=" + msg.Id + CQSuffix
}

// ProcessGRPC
func (msg *MessageElementFace) ProcessGRPC(segment *MessageSegmentGRPC) {
	segment.Data = &MessageSegmentGRPC_MessageElementFace{
		MessageElementFace: msg.ToGRPC(),
	}
}

func (msg *MessageElementFace) ToGRPC() *MessageElementFaceGRPC {
	return &MessageElementFaceGRPC{
		Id: msg.Id,
	}
}

func (msg *MessageElementFaceGRPC) ToStruct() *MessageElementFace {
	return &MessageElementFace{
		Id: msg.Id,
	}
}

func init() {
	messageElementUnmarshalJSONMap["face"] = func(data []byte) (MessageElement, error) {
		var result MessageElementFace
		err := json.Unmarshal(data, &result)
		return &result, err
	}
	messageSegmentGRPCToStructMap["face"] = func(msg *MessageSegmentGRPC) (MessageElement, error) {
		if msg.Data == nil {
			return nil, nil
		}
		switch data := msg.Data.(type) {
		case *MessageSegmentGRPC_MessageElementFace:
			return data.MessageElementFace.ToStruct(), nil
		default:
			return nil, nil
		}
	}
}
