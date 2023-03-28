package model

import "encoding/json"

type MessageElementXml struct {
	// data	string	xml内容, xml中的value部分, 记得实体化处理
	Data string `json:"data"`
	// resid	int32	可能为空, 或空字符串
	Resid int32 `json:"resid"`
}

func (msg *MessageElementXml) Type() string {
	return "xml"
}

func (msg *MessageElementXml) Enabled() bool {
	return true
}

func (msg *MessageElementXml) ToCQCode() string {
	return CQPrefix + msg.Type() + CQEleSplit + "data=" + msg.Data + CQSuffix
}

// ProcessGRPC
func (msg *MessageElementXml) ProcessGRPC(segment *MessageSegmentGRPC) {
	segment.Data = &MessageSegmentGRPC_MessageElementXml{
		MessageElementXml: msg.ToGRPC(),
	}
}

func (msg *MessageElementXml) ToGRPC() *MessageElementXmlGRPC {
	return &MessageElementXmlGRPC{
		Data:  msg.Data,
		Resid: msg.Resid,
	}
}

func (msg *MessageElementXmlGRPC) ToStruct() *MessageElementXml {
	return &MessageElementXml{
		Data:  msg.Data,
		Resid: msg.Resid,
	}
}
func init() {
	messageElementUnmarshalJSONMap["xml"] = func(data []byte) (MessageElement, error) {
		var result MessageElementXml
		err := json.Unmarshal(data, &result)
		return &result, err
	}
	messageSegmentGRPCToStructMap["xml"] = func(msg *MessageSegmentGRPC) (MessageElement, error) {
		if msg.Data == nil {
			return nil, nil
		}
		switch data := msg.Data.(type) {
		case *MessageSegmentGRPC_MessageElementXml:
			return data.MessageElementXml.ToStruct(), nil
		default:
			return nil, nil
		}
	}
}
