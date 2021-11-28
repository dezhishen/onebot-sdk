package model

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestDice(t *testing.T) {
	data := []byte("[{\"type\":\"dice\",\"data\":{}}]")
	var msg []MessageSegment
	err := json.Unmarshal(data, &msg)
	if err != nil {
		panic(err)
	}
	if msg[0].Data.Type() != "dice" {
		panic(errors.New("类型错误"))
	}
	_, ok := msg[0].Data.(DiceMessage)
	if !ok {
		panic(errors.New("类型错误"))
	}
}
