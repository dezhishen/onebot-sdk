package model

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestForward(t *testing.T) {
	data := []byte("[{\"type\":\"forward\",\"data\":{\"id\":\"123456\"}}]")
	var msg []MessageSegment
	err := json.Unmarshal(data, &msg)
	if err != nil {
		panic(err)
	}
	if msg[0].Data.Type() != "forward" {
		panic(errors.New("类型错误"))
	}
	_, ok := msg[0].Data.(*MessageElementForward)
	if !ok {
		panic(errors.New("类型错误"))
	}
}
