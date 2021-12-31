package model

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestJson(t *testing.T) {
	data := []byte("[{\"type\":\"json\",\"data\":{\"data\":\"{\\\"app\\\": ...\"}}]")
	var msg []MessageSegment
	err := json.Unmarshal(data, &msg)
	if err != nil {
		panic(err)
	}
	if msg[0].Data.Type() != "json" {
		panic(errors.New("类型错误"))
	}
	_, ok := msg[0].Data.(*MessageElementJson)
	if !ok {
		panic(errors.New("类型错误"))
	}
}
