package main

import "github.com/dezhishen/onebot-sdk/pkg/model"

func main() {
	jsonStr := "{\"type\":\"node\",\"data\":{\"user_id\":\"10001000\",\"nickname\":\"某人\",\"content\":[{\"type\":\"face\",\"data\":{\"id\":\"123\"}},{\"type\":\"text\",\"data\":{\"text\":\"哈喽～\"}}]}}"
	model.GenMessageElement(jsonStr, "./pkg/model")
}
