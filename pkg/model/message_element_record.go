package model

import "encoding/json"

type MessageElementRecord struct {
	// file	✓	✓[1]	-	语音文件名
	File string `json:"file"`
	// magic	✓	✓	0 1	发送时可选, 默认 0, 设置为 1 表示变声
	Magic string `json:"magic"`
	// url	✓		-	语音 URL
	Url string `json:"url"`
	// cache		✓	0 1	只在通过网络 URL 发送时有效, 表示是否使用已缓存的文件, 默认 1
	Cache string `json:"cache"`
	// proxy		✓	0 1	只在通过网络 URL 发送时有效, 表示是否通过代理下载文件 ( 需通过环境变量或配置文件配置代理 ) , 默认 1
	Proxy string `json:"proxy"`
	// timeout		✓	-	只在通过网络 URL 发送时有效, 单位秒, 表示下载网络文件的超时时间 , 默认不超时
	Timeout string `json:"timeout"`
}

func (msg *MessageElementRecord) Type() string {
	return "record"
}

func (msg *MessageElementRecord) Enabled() bool {
	return true
}

func (msg *MessageElementRecord) ToCQCode() string {
	return CQPrefix + msg.Type() +
		CQEleSplit + "file=" + msg.File +
		CQEleSplit + "magic=" + msg.Magic +
		CQEleSplit + "url=" + msg.Url +
		CQEleSplit + "cache=" + msg.Cache +
		CQEleSplit + "proxy=" + msg.Proxy +
		CQEleSplit + "timeout=" + msg.Timeout +
		CQSuffix
}

// ProcessGRPC
func (msg *MessageElementRecord) ProcessGRPC(segment *MessageSegmentGRPC) {
	segment.Data = &MessageSegmentGRPC_MessageElementRecord{
		MessageElementRecord: msg.ToGRPC(),
	}
}

func (msg *MessageElementRecord) ToGRPC() *MessageElementRecordGRPC {
	return &MessageElementRecordGRPC{
		File:    msg.File,
		Magic:   msg.Magic,
		Url:     msg.Url,
		Cache:   msg.Cache,
		Proxy:   msg.Proxy,
		Timeout: msg.Timeout,
	}
}

func (msg *MessageElementRecordGRPC) ToStruct() *MessageElementRecord {
	return &MessageElementRecord{
		File:    msg.File,
		Magic:   msg.Magic,
		Url:     msg.Url,
		Cache:   msg.Cache,
		Proxy:   msg.Proxy,
		Timeout: msg.Timeout,
	}
}

func init() {
	messageElementUnmarshalJSONMap["record"] = func(data []byte) (MessageElement, error) {
		var result MessageElementRecord
		err := json.Unmarshal(data, &result)
		return &result, err
	}
	messageSegmentGRPCToStructMap["record"] = func(msg *MessageSegmentGRPC) (MessageElement, error) {
		if msg.Data == nil {
			return nil, nil
		}
		switch data := msg.Data.(type) {
		case *MessageSegmentGRPC_MessageElementRecord:
			return data.MessageElementRecord.ToStruct(), nil
		default:
			return nil, nil
		}
	}
}
