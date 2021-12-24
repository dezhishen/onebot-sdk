package model

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestNode(t *testing.T) {
	data := []byte("[{\"type\":\"node\",\"data\":{\"user_id\":\"10001000\",\"nickname\":\"某人\",\"content\":[{\"type\":\"face\",\"data\":{\"id\":\"123\"}},{\"type\":\"text\",\"data\":{\"text\":\"哈喽～\"}}]}}]")
	var msg []MessageSegment
	err := json.Unmarshal(data, &msg)
	if err != nil {
		panic(err)
	}
	if msg[0].Data.Type() != "node" {
		panic(errors.New("类型错误"))
	}
	_, ok := msg[0].Data.(MessageElementNode)
	if !ok {
		panic(errors.New("类型错误"))
	}
}类型错误"))
	}
}
