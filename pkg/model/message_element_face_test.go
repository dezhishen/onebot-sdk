package model

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestFace(t *testing.T) {
	data := []byte("[{\"type\":\"face\",\"data\":{\"id\":\"1231\"}}]")
	var msg []MessageSegment
	err := json.Unmarshal(data, &msg)
	if err != nil {
		panic(err)
	}
	if msg[0].Data.Type() != "face" {
		panic(errors.New("类型错误"))
	}
	field, ok := msg[0].Data.(*MessageElementFace)
	if !ok {
		panic(errors.New("类型错误"))
	}
	if field.Id != "1231" {
		panic(errors.New("类型错误"))
	}
}
