package model

import "encoding/json"

type MessageElementRedbag struct {
	// title	string	祝福语/口令
	Title string `json:"title"`
}

func (msg *MessageElementRedbag) Type() string {
	return "redbag"
}

func (msg *MessageElementRedbag) Enabled() bool {
	return true
}

// ProcessGRPC
func (msg *MessageElementRedbag) ProcessGRPC(segment *MessageSegmentGRPC) {
	segment.Data = &MessageSegmentGRPC_MessageElementRedbag{
		MessageElementRedbag: msg.ToGRPC(),
	}
}

func (msg *MessageElementRedbag) ToGRPC() *MessageElementRedbagGRPC {
	return &MessageElementRedbagGRPC{
		Title: msg.Title,
	}
}

func (msg *MessageElementRedbagGRPC) ToStruct() *MessageElementRedbag {
	return &MessageElementRedbag{
		Title: msg.Title,
	}
}

func init() {
	messageElementUnmarshalJSONMap["redbag"] = func(data []byte) (MessageElement, error) {
		var result MessageElementRedbag
		err := json.Unmarshal(data, &result)
		return &result, err
	}
	messageSegmentGRPCToStructMap["redbag"] = func(msg *MessageSegmentGRPC) (MessageElement, error) {
		if msg.Data == nil {
			return nil, nil
		}
		switch data := msg.Data.(type) {
		case *MessageSegmentGRPC_MessageElementRedbag:
			return data.MessageElementRedbag.ToStruct(), nil
		default:
			return nil, nil
		}
	}
}
