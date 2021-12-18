package model

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestImage(t *testing.T) {
	data := []byte("[{\"type\":\"image\",\"data\":{\"file\":\"http://baidu.com/1.jpg\"}}]")
	var msg []MessageSegment
	err := json.Unmarshal(data, &msg)
	if err != nil {
		panic(err)
	}
	if msg[0].Data.Type() != "image" {
		panic(errors.New("类型错误"))
	}
	field, ok := msg[0].Data.(MessageElementImage)
	if !ok {
		panic(errors.New("类型错误"))
	}
	if field.File != "http://baidu.com/1.jpg" {
		panic(errors.New("值错误"))
	}
}
