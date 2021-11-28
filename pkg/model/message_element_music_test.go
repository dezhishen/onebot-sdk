package model

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestMusic(t *testing.T) {
	data := []byte("[{\"type\":\"music\",\"data\":{\"type\":\"custom\",\"url\":\"http://baidu.com\",\"audio\":\"http://baidu.com/1.mp3\",\"title\":\"音乐标题\",\"content\":\"发送时可选，内容描述\",\"image\":\"发送时可选，图片 URL\"}}]")
	var msg []MessageSegment
	err := json.Unmarshal(data, &msg)
	if err != nil {
		panic(err)
	}
	if msg[0].Data.Type() != "music" {
		panic(errors.New("类型错误"))
	}
	_, ok := msg[0].Data.(MusicMessage)
	if !ok {
		panic(errors.New("类型错误"))
	}
}
