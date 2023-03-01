package model

import "encoding/json"

type MessageElementJson struct {
	// data	string	json内容, json的所有字符串记得实体化处理(转义)
	Data string `json:"data"`
	// resid	int32	默认不填为0, 走小程序通道, 填了走富文本通道发送
	Resid int32 `json:"resid"`
}

func (msg *MessageElementJson) Type() string {
	return "json"
}

func (msg *MessageElementJson) Enabled() bool {
	return true
}

// ProcessGRPC
func (msg *MessageElementJson) ProcessGRPC(segment *MessageSegmentGRPC) {
	segment.Data = &MessageSegmentGRPC_MessageElementJson{
		MessageElementJson: msg.ToGRPC(),
	}
}

func (msg *MessageElementJson) ToGRPC() *MessageElementJsonGRPC {
	return &MessageElementJsonGRPC{
		Data:  msg.Data,
		Resid: msg.Resid,
	}
}

func (msg *MessageElementJsonGRPC) ToStruct() *MessageElementJson {
	return &MessageElementJson{
		Data:  msg.Data,
		Resid: msg.Resid,
	}
}

func init() {
	messageElementUnmarshalJSONMap["json"] = func(data []byte) (MessageElement, error) {
		var result MessageElementJson
		err := json.Unmarshal(data, &result)
		return &result, err
	}
	messageSegmentGRPCToStructMap["json"] = func(msg *MessageSegmentGRPC) (MessageElement, error) {
		if msg.Data == nil {
			return nil, nil
		}
		switch data := msg.Data.(type) {
		case *MessageSegmentGRPC_MessageElementJson:
			return data.MessageElementJson.ToStruct(), nil
		default:
			return nil, nil
		}
	}
}
