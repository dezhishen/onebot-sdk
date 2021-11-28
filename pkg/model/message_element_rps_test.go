package model

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestRps(t *testing.T) {
	data := []byte("[{\"type\":\"rps\",\"data\":{}}]")
	var msg []MessageSegment
	err := json.Unmarshal(data, &msg)
	if err != nil {
		panic(err)
	}
	if msg[0].Data.Type() != "rps" {
		panic(errors.New("类型错误"))
	}
	_, ok := msg[0].Data.(RpsMessage)
	if !ok {
		panic(errors.New("类型错误"))
	}
}
