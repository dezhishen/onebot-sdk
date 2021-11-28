package model

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestContact(t *testing.T) {
	data := []byte("[{\"type\":\"contact\",\"data\":{\"type\":\"group\",\"id\":\"100100\"}}]")
	var msg []MessageSegment
	err := json.Unmarshal(data, &msg)
	if err != nil {
		panic(err)
	}
	if msg[0].Data.Type() != "contact" {
		panic(errors.New("类型错误"))
	}
	_, ok := msg[0].Data.(ContactMessage)
	if !ok {
		panic(errors.New("类型错误"))
	}
}
