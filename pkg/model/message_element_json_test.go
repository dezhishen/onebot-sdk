package model

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestJson(t *testing.T) {
	data := []byte("[{\"type\":\"text\",\"data\":{\"text\":\"1231\"}},{\"type\":\"face\",\"data\":{\"id\":\"1\"}}]")
	var messageSegment []MessageSegment
	err := json.Unmarshal(data, &messageSegment)
	if err != nil {
		panic(err)
	}
	if len(messageSegment) != 2 {
		panic(errors.New("转换错误，长度错误"))
	}
	if messageSegment[0].Type != "text" {
		panic(errors.New("转换错误，第一个消息体应该是text"))
	}
	if messageSegment[1].Type != "face" {
		panic(errors.New("转换错误，第二个消息体应该是face"))
	}
}
