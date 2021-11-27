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
	field, ok := msg[0].Data.(FaceMessage)
	if !ok {
		panic(errors.New("类型错误"))
	}
	if field.ID != "1231" {
		panic(errors.New("类型错误"))
	}
}

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
