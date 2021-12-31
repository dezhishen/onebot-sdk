package model

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestAt(t *testing.T) {
	data := []byte("[{\"type\":\"at\",\"data\":{\"qq\":\"10001000\"}}]")
	var msg []MessageSegment
	err := json.Unmarshal(data, &msg)
	if err != nil {
		panic(err)
	}
	if msg[0].Data.Type() != "at" {
		panic(errors.New("类型错误"))
	}
	_, ok := msg[0].Data.(*MessageElementAt)
	if !ok {
		panic(errors.New("类型错误"))
	}
}
