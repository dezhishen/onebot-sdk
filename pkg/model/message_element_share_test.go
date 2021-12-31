package model

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestShare(t *testing.T) {
	data := []byte("[{\"type\":\"share\",\"data\":{\"url\":\"http://baidu.com\",\"title\":\"百度\",\"content\":\"发送时可选，内容描述\",\"image\":\"发送时可选，图片 URL\"}}]")
	var msg []MessageSegment
	err := json.Unmarshal(data, &msg)
	if err != nil {
		panic(err)
	}
	if msg[0].Data.Type() != "share" {
		panic(errors.New("类型错误"))
	}
	_, ok := msg[0].Data.(*MessageElementShare)
	if !ok {
		panic(errors.New("类型错误"))
	}
}
