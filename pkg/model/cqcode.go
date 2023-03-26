package model

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	CQTypeText      = "text"
	CQTypeFace      = "face"
	CQTypeRecord    = "record"
	CQTypeVideo     = "video"
	CQTypeAt        = "at"
	CQTypeRps       = "rps"
	CQTypeDice      = "dice"
	CQTypeShake     = "shake"
	CQTypeAnonymous = "anonymous"
	CQTypeShare     = "share"
	CQTypeContact   = "contact"
	CQTypeLocation  = "location"
	CQTypeMusic     = "music"
	CQTypeImage     = "image"
	CQTypeReply     = "reply"
	CQTypeRedbag    = "redbag"
	CQTypePoke      = "poke"
	CQTypeGift      = "gift"
	CQTypeForward   = "forward"
	CQTypeNode      = "node"
	CQTypeXml       = "xml"
	CQTypeJson      = "json"
	CQTypeCardImage = "cardimage"
	CQTypeTts       = "tts"
)

const (
	CQPrefix = "[CQ:"
	CQSuffix = "]"

	CQEleSplit = ","
	CQKVSplit  = "="
)

func ParseStringMessage(strMsg string) ([]*MessageSegment, error) {
	var res []*MessageSegment
	for {
		if strMsg == "" {
			return res, nil
		}
		cqstart := strings.Index(strMsg, CQPrefix)
		if cqstart < 0 {
			// end parse
			res = append(res, &MessageSegment{
				Type: CQTypeText,
				Data: &MessageElementText{
					strMsg,
				},
			})
			return res, nil
		}

		if cqstart != 0 {
			res = append(res, &MessageSegment{
				Type: CQTypeText,
				Data: &MessageElementText{
					strMsg[:cqstart],
				},
			})
			strMsg = strMsg[cqstart:]
		}

		cqend := strings.Index(strMsg, CQSuffix)
		if cqend < 0 {
			return nil, fmt.Errorf("invalid cq code, can not find surfix in %s", strMsg)
		}
		cqCode := strMsg[:cqend+1]
		strMsg = strMsg[cqend+1:]

		segment, err := ParseCQCode(cqCode)
		if err != nil {
			return nil, err
		}
		res = append(res, segment)
	}
}

func ParseCQCode(cqCode string) (*MessageSegment, error) {
	cqCodeBody := cqCode[len(CQPrefix) : len(cqCode)-len(CQSuffix)]
	split := strings.Split(cqCodeBody, CQEleSplit)
	if len(split) == 0 {
		return nil, fmt.Errorf("empty CQcode: %s", cqCodeBody)
	}
	cqData := map[string]string{}
	for _, s := range split[1:] {
		kv := strings.SplitN(s, CQKVSplit, 2)
		if len(kv) != 2 {
			return nil, fmt.Errorf("error kv in CQCode: %s", cqCode)
		}
		cqData[kv[0]] = kv[1]
	}
	cqDataStr, err := json.Marshal(cqData)
	if err != nil {
		return nil, err
	}

	res := &MessageSegment{
		Type: split[0],
	}
	messageElement := getMessageElementFromType(res.Type)
	if messageElement != nil {
		if err = json.Unmarshal(cqDataStr, messageElement); err != nil {
			return nil, err
		}
		res.Data = messageElement
	} else {
		res.Data = &MessageElementText{
			Text: cqCode,
		}
	}
	return res, nil
}

func getMessageElementFromType(cqType string) MessageElement {
	switch cqType {
	case CQTypeText:
		return &MessageElementText{}
	case CQTypeFace:
		return &MessageElementFace{}
	case CQTypeRecord:
		return &MessageElementRecord{}
	case CQTypeVideo:
		return &MessageElementVideo{}
	case CQTypeAt:
		return &MessageElementAt{}
	case CQTypeRps:
		return &MessageElementRps{}
	case CQTypeDice:
		return &MessageElementDice{}
	case CQTypeShake:
		return &MessageElementShake{}
	case CQTypeAnonymous:
		return &MessageElementAnonymous{}
	case CQTypeShare:
		return &MessageElementShare{}
	case CQTypeLocation:
		return &MessageElementLocation{}
	case CQTypeMusic:
		return &MessageElementMusic{}
	case CQTypeImage:
		return &MessageElementImage{}
	case CQTypeReply:
		return &MessageElementReply{}
	case CQTypeRedbag:
		return &MessageElementRedbag{}
	case CQTypePoke:
		return &MessageElementPoke{}
	case CQTypeGift:
		return &MessageElementGift{}
	case CQTypeForward:
		return &MessageElementForward{}
	case CQTypeNode:
		return &MessageElementNode{}
	case CQTypeXml:
		return &MessageElementXml{}
	case CQTypeJson:
		return &MessageElementJson{}
	case CQTypeCardImage:
		return nil
	case CQTypeTts:
		return &MessageElementTts{}
	default:
		return nil
	}
}
