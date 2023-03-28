package model

import (
	"encoding/json"
	"fmt"
)

type MessageElementVideo struct {
	// file	string	-	视频地址, 支持http和file发送
	File string `json:"file"`
	// cover	string	-	视频封面, 支持http, file和base64发送, 格式必须为jpg
	Cover string `json:"cover"`
	// c	int	2 3	通过网络下载视频时的线程数, 默认单线程. (在资源不支持并发时会自动处理)
	C uint32 `json:"c"`
}

func (msg *MessageElementVideo) Type() string {
	return "video"
}

func (msg *MessageElementVideo) Enabled() bool {
	return true
}

func (msg *MessageElementVideo) ToCQCode() string {
	return CQPrefix + msg.Type() +
		CQEleSplit + "file=" + msg.File +
		CQEleSplit + "cover=" + msg.Cover +
		CQEleSplit + "c=" + fmt.Sprintf("%d", msg.C) +
		CQSuffix
}

// ProcessGRPC
func (msg *MessageElementVideo) ProcessGRPC(segment *MessageSegmentGRPC) {
	segment.Data = &MessageSegmentGRPC_MessageElementVideo{
		MessageElementVideo: msg.ToGRPC(),
	}
}

func (msg *MessageElementVideo) ToGRPC() *MessageElementVideoGRPC {
	return &MessageElementVideoGRPC{
		File:  msg.File,
		Cover: msg.Cover,
		C:     msg.C,
	}
}

func (msg *MessageElementVideoGRPC) ToStruct() *MessageElementVideo {
	return &MessageElementVideo{
		File:  msg.File,
		Cover: msg.Cover,
		C:     msg.C,
	}
}

func init() {
	messageElementUnmarshalJSONMap["video"] = func(data []byte) (MessageElement, error) {
		var result MessageElementVideo
		err := json.Unmarshal(data, &result)
		return &result, err
	}
	messageSegmentGRPCToStructMap["video"] = func(msg *MessageSegmentGRPC) (MessageElement, error) {
		if msg.Data == nil {
			return nil, nil
		}
		switch data := msg.Data.(type) {
		case *MessageSegmentGRPC_MessageElementVideo:
			return data.MessageElementVideo.ToStruct(), nil
		default:
			return nil, nil
		}
	}
}
