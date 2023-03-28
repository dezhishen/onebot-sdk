package model

import "encoding/json"

type MessageElementImage struct {
	// 图片文件名
	File string `json:"file"`
	// 图片类型, flash 表示闪照, show 表示秀图, 默认普通图片
	ImageType string `json:"type"`
	// 图片子类型, 只出现在群聊
	// 0	正常图片
	// 1	表情包, 在客户端会被分类到表情包图片并缩放显示
	// 2	热图
	// 3	斗图
	// 4	智图?
	// 7	贴图
	// 8	自拍
	// 9	贴图广告?
	// 10	有待测试
	// 13	热搜图
	SubType string `json:"sub_type"`
	// 图片 URL
	Url string `json:"url"`
	// 是否使用缓存 0/1 只在通过网络 URL 发送时有效, 表示是否使用已缓存的文件, 默认 1
	Cache string `json:"cache"`
	// 发送秀图时的特效id, 默认为40000
	// 40000	普通
	// 40001	幻影
	// 40002	抖动
	// 40003	生日
	// 40004	爱你
	// 40005	征友
	Id string `json:"id"`
	// 通过网络下载图片时的线程数, 默认单线程. (在资源不支持并发时会自动处理)
	C string `json:"c"`
}

func (msg *MessageElementImage) Type() string {
	return "image"
}
func (msg *MessageElementImage) Enabled() bool {
	return true
}

func (msg *MessageElementImage) ToCQCode() string {
	return CQPrefix + msg.Type() +
		CQEleSplit + "file=" + msg.File +
		CQEleSplit + "type=" + msg.ImageType +
		CQEleSplit + "sub_type=" + msg.SubType +
		CQEleSplit + "url=" + msg.Url +
		CQEleSplit + "cache=" + msg.Cache +
		CQEleSplit + "id=" + msg.Id +
		CQEleSplit + "c=" + msg.C +
		CQSuffix
}

// ProcessGRPC
func (msg *MessageElementImage) ProcessGRPC(segment *MessageSegmentGRPC) {
	segment.Data = &MessageSegmentGRPC_MessageElementImage{
		MessageElementImage: msg.ToGRPC(),
	}
}

// ToGRPC
func (msg *MessageElementImage) ToGRPC() *MessageElementImageGRPC {
	return &MessageElementImageGRPC{
		File:      msg.File,
		Url:       msg.Url,
		ImageType: msg.ImageType,
		SubType:   msg.SubType,
		Cache:     msg.Cache,
		Id:        msg.Id,
		C:         msg.C,
	}
}

func (msg *MessageElementImageGRPC) ToStruct() *MessageElementImage {
	return &MessageElementImage{
		File:      msg.File,
		Url:       msg.Url,
		ImageType: msg.ImageType,
		SubType:   msg.SubType,
		Cache:     msg.Cache,
		Id:        msg.Id,
		C:         msg.C,
	}
}

func init() {
	messageElementUnmarshalJSONMap["image"] = func(data []byte) (MessageElement, error) {
		var result MessageElementImage
		err := json.Unmarshal(data, &result)
		return &result, err
	}
	messageSegmentGRPCToStructMap["image"] = func(msg *MessageSegmentGRPC) (MessageElement, error) {
		if msg.Data == nil {
			return nil, nil
		}
		switch data := msg.Data.(type) {
		case *MessageSegmentGRPC_MessageElementImage:
			return data.MessageElementImage.ToStruct(), nil
		default:
			return nil, nil
		}
	}
}
