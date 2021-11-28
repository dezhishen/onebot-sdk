package model

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestText(t *testing.T) {
	data := []byte("[{\"type\":\"text\",\"data\":{\"text\":\"1231\"}}]")
	var msg []MessageSegment
	err := json.Unmarshal(data, &msg)
	if err != nil {
		panic(err)
	}
	if msg[0].Data.Type() != "text" {
		panic(errors.New("类型错误"))
	}
	field, ok := msg[0].Data.(TextMessage)
	if !ok {
		panic(errors.New("类型错误"))
	}
	if field.Text != "1231" {
		panic(errors.New("类型错误"))
	}
}
