package model

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestXml(t *testing.T) {
	data := []byte("[{\"type\":\"xml\",\"data\":{\"data\":\"<?xml ...\"}}]")
	var msg []MessageSegment
	err := json.Unmarshal(data, &msg)
	if err != nil {
		panic(err)
	}
	if msg[0].Data.Type() != "xml" {
		panic(errors.New("类型错误"))
	}
	_, ok := msg[0].Data.(MessageElementXml)
	if !ok {
		panic(errors.New("类型错误"))
	}
}
