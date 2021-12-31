package model

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestShake(t *testing.T) {
	data := []byte("[{\"type\":\"shake\",\"data\":{}}]")
	var msg []MessageSegment
	err := json.Unmarshal(data, &msg)
	if err != nil {
		panic(err)
	}
	if msg[0].Data.Type() != "shake" {
		panic(errors.New("类型错误"))
	}
	_, ok := msg[0].Data.(*MessageElementShake)
	if !ok {
		panic(errors.New("类型错误"))
	}
}
