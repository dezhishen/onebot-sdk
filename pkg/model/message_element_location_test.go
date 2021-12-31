package model

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestLocation(t *testing.T) {
	data := []byte("[{\"type\":\"location\",\"data\":{\"lat\":\"39.8969426\",\"lon\":\"116.3109099\",\"title\":\"发送时可选，标题\",\"content\":\"发送时可选，内容描述\"}}]")
	var msg []MessageSegment
	err := json.Unmarshal(data, &msg)
	if err != nil {
		panic(err)
	}
	if msg[0].Data.Type() != "location" {
		panic(errors.New("类型错误"))
	}
	_, ok := msg[0].Data.(*MessageElementLocation)
	if !ok {
		panic(errors.New("类型错误"))
	}
}
