package model

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestRecord(t *testing.T) {
	data := []byte("[{\"type\":\"record\",\"data\":{\"file\":\"http://baidu.com/1.mp3\"}}]")
	var msg []MessageSegment
	err := json.Unmarshal(data, &msg)
	if err != nil {
		panic(err)
	}
	if msg[0].Data.Type() != "record" {
		panic(errors.New("类型错误"))
	}
	_, ok := msg[0].Data.(MessageElementRecord)
	if !ok {
		panic(errors.New("类型错误"))
	}
}
