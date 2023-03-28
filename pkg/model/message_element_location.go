package model

import "encoding/json"

type MessageElementLocation struct {
	Lat     string `json:"lat"`
	Lon     string `json:"lon"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (msg *MessageElementLocation) Type() string {
	return "location"
}

func (msg *MessageElementLocation) Enabled() bool {
	return false
}

func (msg *MessageElementLocation) ToCQCode() string {
	return CQPrefix + msg.Type() +
		CQEleSplit + "lat=" + msg.Lat +
		CQEleSplit + "lon=" + msg.Lon +
		CQEleSplit + "title=" + msg.Title +
		CQEleSplit + "content=" + msg.Content +
		CQSuffix
}

// ProcessGRPC
func (msg *MessageElementLocation) ProcessGRPC(segment *MessageSegmentGRPC) {
	segment.Data = &MessageSegmentGRPC_MessageElementLocation{
		MessageElementLocation: msg.ToGRPC(),
	}
}

func (msg *MessageElementLocation) ToGRPC() *MessageElementLocationGRPC {
	return &MessageElementLocationGRPC{
		Lat:     msg.Lat,
		Lon:     msg.Lon,
		Title:   msg.Title,
		Content: msg.Content,
	}
}

func (msg *MessageElementLocationGRPC) ToStruct() *MessageElementLocation {
	return &MessageElementLocation{
		Lat:     msg.Lat,
		Lon:     msg.Lon,
		Title:   msg.Title,
		Content: msg.Content,
	}
}

func init() {
	messageElementUnmarshalJSONMap["location"] = func(data []byte) (MessageElement, error) {
		var result MessageElementLocation
		err := json.Unmarshal(data, &result)
		return &result, err
	}
	messageSegmentGRPCToStructMap["jsolocationn"] = func(msg *MessageSegmentGRPC) (MessageElement, error) {
		if msg.Data == nil {
			return nil, nil
		}
		switch data := msg.Data.(type) {
		case *MessageSegmentGRPC_MessageElementLocation:
			return data.MessageElementLocation.ToStruct(), nil
		default:
			return nil, nil
		}
	}
}
