package model

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestReply(t *testing.T) {
	data := []byte("[{\"type\":\"reply\",\"data\":{\"id\":\"123456\"}}]")
	var msg []MessageSegment
	err := json.Unmarshal(data, &msg)
	if err != nil {
		panic(err)
	}
	if msg[0].Data.Type() != "reply" {
		panic(errors.New("类型错误"))
	}
	_, ok := msg[0].Data.(MessageElementReply)
	if !ok {
		panic(errors.New("类型错误"))
	}
}
