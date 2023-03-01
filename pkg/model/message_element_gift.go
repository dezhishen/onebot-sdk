package model

import "encoding/json"

type MessageElementGift struct {
	// qq	int64	接收礼物的成员
	Qq int64 `json:"qq"`
	// id	int	礼物的类型
	// 0	甜 Wink
	// 1	快乐肥宅水
	// 2	幸运手链
	// 3	卡布奇诺
	// 4	猫咪手表
	// 5	绒绒手套
	// 6	彩虹糖果
	// 7	坚强
	// 8	告白话筒
	// 9	牵你的手
	// 10	可爱猫咪
	// 11	神秘面具
	// 12	我超忙的
	// 13	爱心口罩
	Id int64 `json:"id"`
}

func (msg *MessageElementGift) Type() string {
	return "gift"
}

func (msg *MessageElementGift) Enabled() bool {
	return true
}

// ProcessGRPC
func (msg *MessageElementGift) ProcessGRPC(segment *MessageSegmentGRPC) {
	segment.Data = &MessageSegmentGRPC_MessageElementGift{
		MessageElementGift: msg.ToGRPC(),
	}
}

func (msg *MessageElementGift) ToGRPC() *MessageElementGiftGRPC {
	return &MessageElementGiftGRPC{
		Qq: msg.Qq,
		Id: msg.Id,
	}
}

func (msg *MessageElementGiftGRPC) ToStruct() *MessageElementGift {
	return &MessageElementGift{
		Qq: msg.Qq,
		Id: msg.Id,
	}
}

func init() {
	messageElementUnmarshalJSONMap["gift"] = func(data []byte) (MessageElement, error) {
		var result MessageElementGift
		err := json.Unmarshal(data, &result)
		return &result, err
	}
	messageSegmentGRPCToStructMap["gift"] = func(msg *MessageSegmentGRPC) (MessageElement, error) {
		if msg.Data == nil {
			return nil, nil
		}
		switch data := msg.Data.(type) {
		case *MessageSegmentGRPC_MessageElementGift:
			return data.MessageElementGift.ToStruct(), nil
		default:
			return nil, nil
		}
	}
}
