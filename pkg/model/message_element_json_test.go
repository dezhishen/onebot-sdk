package model

import (
	"encoding/json"
	"testing"
)

func TestJson(t *testing.T) {
	data := []byte("[{\"type\":\"text\",\"data\":{\"text\":\"1231\"}},{\"type\":\"face\",\"data\":{\"id\":\"1\"}}]")
	var messageSegment []MessageSegment
	err := json.Unmarshal(data, &messageSegment)
	if err != nil {
		panic(err)
	}
	print(messageSegment)
}
