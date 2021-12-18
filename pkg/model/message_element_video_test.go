package model

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestVideo(t *testing.T) {
	data := []byte("[{\"type\":\"video\",\"data\":{\"file\":\"http://baidu.com/1.mp4\"}}]")
	var msg []MessageSegment
	err := json.Unmarshal(data, &msg)
	if err != nil {
		panic(err)
	}
	if msg[0].Data.Type() != "video" {
		panic(errors.New("类型错误"))
	}
	_, ok := msg[0].Data.(MessageElementVideo)
	if !ok {
		panic(errors.New("类型错误"))
	}
}
