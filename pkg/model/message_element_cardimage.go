package model

import (
	"encoding/json"
	"fmt"
)

type MessageElementCardimage struct {
	// file	string	和image的file字段对齐, 支持也是一样的
	File string `json:"file"`
	// minwidth	int64	默认不填为400, 最小width
	Minwidth int64 `json:"minwidth"`
	// minheight	int64	默认不填为400, 最小height
	Minheight int64 `json:"minheight"`
	// maxwidth	int64	默认不填为500, 最大width
	Maxwidth int64 `json:"maxwidth"`
	// maxheight	int64	默认不填为1000, 最大height
	Maxheight int64 `json:"maxheight"`
	// source	string	分享来源的名称, 可以留空
	Source string `json:"source"`
	// icon	string	分享来源的icon图标url, 可以留空
	Icon string `json:"icon"`
}

func (msg *MessageElementCardimage) Type() string {
	return "cardimage"
}

func (msg *MessageElementCardimage) Enabled() bool {
	return true
}

// ToCQCode

func (msg *MessageElementCardimage) ToCQCode() string {
	cqCode := CQPrefix + CQEleSplit + "file=" + msg.File
	if msg.Minwidth != 0 {
		cqCode += (CQEleSplit + fmt.Sprintf("minwidth=%d", msg.Minwidth))
	}
	if msg.Minheight != 0 {
		cqCode += (CQEleSplit + fmt.Sprintf("minheight=%d", msg.Minheight))
	}
	if msg.Icon != "" {
		cqCode += (CQEleSplit + "icon=" + msg.Icon)
	}
	if msg.Source != "" {
		cqCode += (CQEleSplit + "source=" + msg.Source)
	}
	return cqCode + CQSuffix
}

// ProcessGRPC
func (msg *MessageElementCardimage) ProcessGRPC(segment *MessageSegmentGRPC) {
	segment.Data = &MessageSegmentGRPC_MessageElementCardimage{
		MessageElementCardimage: msg.ToGRPC(),
	}
}

func (msg *MessageElementCardimage) ToGRPC() *MessageElementCardimageGRPC {
	return &MessageElementCardimageGRPC{
		File:      msg.File,
		Minwidth:  msg.Minwidth,
		Minheight: msg.Minheight,
		Maxwidth:  msg.Maxwidth,
		Maxheight: msg.Maxheight,
		Source:    msg.Source,
		Icon:      msg.Icon,
	}
}

func (msg *MessageElementCardimageGRPC) ToStruct() *MessageElementCardimage {
	return &MessageElementCardimage{
		File:      msg.File,
		Minwidth:  msg.Minwidth,
		Minheight: msg.Minheight,
		Maxwidth:  msg.Maxwidth,
		Maxheight: msg.Maxheight,
		Source:    msg.Source,
		Icon:      msg.Icon,
	}
}

func init() {
	messageElementUnmarshalJSONMap["cardimage"] = func(data []byte) (MessageElement, error) {
		var result MessageElementCardimage
		err := json.Unmarshal(data, &result)
		return &result, err
	}
	messageSegmentGRPCToStructMap["cardimage"] = func(msg *MessageSegmentGRPC) (MessageElement, error) {
		if msg == nil {
			return nil, nil
		}
		if msg.Data == nil {
			return nil, nil
		}
		switch data := msg.Data.(type) {
		case *MessageSegmentGRPC_MessageElementCardimage:
			return data.MessageElementCardimage.ToStruct(), nil
		default:
			return nil, nil
		}
	}
}
