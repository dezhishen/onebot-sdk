package model

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

type MessageSegment struct {
	Type string         `json:"type"`
	Data MessageElement `json:"data"`
}

func (msg *MessageSegment) ToGRPC() *MessageSegmentGRPC {
	if msg == nil {
		return nil
	}
	result := &MessageSegmentGRPC{
		Type: msg.Type,
	}
	switch msg.Type {
	case "at":
		result.Data = &MessageSegmentGRPC_MessageElementAt{
			MessageElementAt: msg.Data.(*MessageElementAt).ToGRPC(),
		}
	case "contact":
		result.Data = &MessageSegmentGRPC_MessageElementContact{
			MessageElementContact: msg.Data.(*MessageElementContact).ToGRPC(),
		}
	case "dice":
		result.Data = &MessageSegmentGRPC_MessageElementDice{
			MessageElementDice: msg.Data.(*MessageElementDice).ToGRPC(),
		}
	case "face":
		result.Data = &MessageSegmentGRPC_MessageElementFace{
			MessageElementFace: msg.Data.(*MessageElementFace).ToGRPC(),
		}
	case "forward":
		result.Data = &MessageSegmentGRPC_MessageElementForward{
			MessageElementForward: msg.Data.(*MessageElementForward).ToGRPC(),
		}
	case "image":
		result.Data = &MessageSegmentGRPC_MessageElementImage{
			MessageElementImage: msg.Data.(*MessageElementImage).ToGRPC(),
		}
	case "json":
		result.Data = &MessageSegmentGRPC_MessageElementJson{
			MessageElementJson: msg.Data.(*MessageElementJson).ToGRPC(),
		}
	case "location":
		result.Data = &MessageSegmentGRPC_MessageElementLocation{
			MessageElementLocation: msg.Data.(*MessageElementLocation).ToGRPC(),
		}
	case "music":
		result.Data = &MessageSegmentGRPC_MessageElementMusic{
			MessageElementMusic: msg.Data.(*MessageElementMusic).ToGRPC(),
		}
	case "poke":
		result.Data = &MessageSegmentGRPC_MessageElementPoke{
			MessageElementPoke: msg.Data.(*MessageElementPoke).ToGRPC(),
		}
	case "record":
		result.Data = &MessageSegmentGRPC_MessageElementRecord{
			MessageElementRecord: msg.Data.(*MessageElementRecord).ToGRPC(),
		}
	case "reply":
		result.Data = &MessageSegmentGRPC_MessageElementReply{
			MessageElementReply: msg.Data.(*MessageElementReply).ToGRPC(),
		}
	case "rps":
		result.Data = &MessageSegmentGRPC_MessageElementRps{
			MessageElementRps: msg.Data.(*MessageElementRps).ToGRPC(),
		}
	case "shake":
		result.Data = &MessageSegmentGRPC_MessageElementShake{
			MessageElementShake: msg.Data.(*MessageElementShake).ToGRPC(),
		}
	case "share":
		result.Data = &MessageSegmentGRPC_MessageElementShare{
			MessageElementShare: msg.Data.(*MessageElementShare).ToGRPC(),
		}
	case "text":
		result.Data = &MessageSegmentGRPC_MessageElementText{
			MessageElementText: msg.Data.(*MessageElementText).ToGRPC(),
		}
	case "video":
		result.Data = &MessageSegmentGRPC_MessageElementVideo{
			MessageElementVideo: msg.Data.(*MessageElementVideo).ToGRPC(),
		}
	case "xml":
		result.Data = &MessageSegmentGRPC_MessageElementXml{
			MessageElementXml: msg.Data.(*MessageElementXml).ToGRPC(),
		}
	default:
		log.Errorf("未定义的消息类型:%v", result.Type)
	}
	return result
}

func (msg *MessageSegmentGRPC) ToStruct() *MessageSegment {
	if msg == nil {
		return nil
	}
	result := &MessageSegment{
		Type: msg.Type,
	}
	switch msg.Type {
	case "at":
		result.Data = msg.GetMessageElementAt().ToStruct()
	case "contact":
		result.Data = msg.GetMessageElementContact().ToStruct()
	case "dice":
		result.Data = msg.GetMessageElementDice().ToStruct()
	case "face":
		result.Data = msg.GetMessageElementFace().ToStruct()
	case "forward":
		result.Data = msg.GetMessageElementForward().ToStruct()
	case "image":
		result.Data = msg.GetMessageElementImage().ToStruct()
	case "json":
		result.Data = msg.GetMessageElementJson().ToStruct()
	case "location":
		result.Data = msg.GetMessageElementLocation().ToStruct()
	case "music":
		result.Data = msg.GetMessageElementMusic().ToStruct()
	case "poke":
		result.Data = msg.GetMessageElementPoke().ToStruct()
	case "record":
		result.Data = msg.GetMessageElementRecord().ToStruct()
	case "reply":
		result.Data = msg.GetMessageElementReply().ToStruct()
	case "rps":
		result.Data = msg.GetMessageElementRps().ToStruct()
	case "shake":
		result.Data = msg.GetMessageElementShake().ToStruct()
	case "share":
		result.Data = msg.GetMessageElementShare().ToStruct()
	case "text":
		result.Data = msg.GetMessageElementText().ToStruct()
	case "video":
		result.Data = msg.GetMessageElementVideo().ToStruct()
	case "xml":
		result.Data = msg.GetMessageElementXml().ToStruct()
	default:
		log.Errorf("未定义的消息类型:%v", result.Type)
	}
	return result
}

func MessageSegmentGRPCArray2MessageSegmentArray(source []*MessageSegmentGRPC) []*MessageSegment {
	if len(source) == 0 {
		return nil
	}
	var result []*MessageSegment
	for _, e := range source {
		result = append(result, e.ToStruct())
	}
	return result
}

func MessageSegmentArray2MessageSegmentGRPCArray(source []*MessageSegment) []*MessageSegmentGRPC {
	if source == nil || len(source) == 0 {
		return nil
	}
	var result []*MessageSegmentGRPC
	for _, e := range source {
		result = append(result, e.ToGRPC())
	}
	return result
}

var messageElementUnmarshalJSONMap = make(map[string]func(data []byte) (MessageElement, error))

func (msgSeg *MessageSegment) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || data[0] == 'n' { // copied from the Q, can be improved
		return nil
	}
	messageType := gjson.GetBytes(data, "type").Str
	decoder, ok := messageElementUnmarshalJSONMap[messageType]
	if !ok {
		return fmt.Errorf("未找到指定的消息类型,%v", messageType)
	}
	element, err := decoder([]byte(gjson.GetBytes(data, "data").Raw))
	if !ok {
		return fmt.Errorf("未找到指定的消息类型,%v", messageType)
	}
	msgSeg.Type = messageType
	msgSeg.Data = element
	return err
}

type MessageElement interface {
	Type() string
}
