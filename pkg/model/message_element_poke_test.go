package model

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestPoke(t *testing.T) {
	data := []byte("[{\"type\":\"poke\",\"data\":{\"type\":\"126\",\"id\":\"2003\",\"name\":\"name\"}}]")
	var msg []MessageSegment
	err := json.Unmarshal(data, &msg)
	if err != nil {
		panic(err)
	}
	if msg[0].Data.Type() != "poke" {
		panic(errors.New("类型错误"))
	}
	_, ok := msg[0].Data.(MessageElementPoke)
	if !ok {
		panic(errors.New("类型错误"))
	}
}
