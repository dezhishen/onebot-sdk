package model

import "encoding/json"

type PokeMessage struct {
	PokeType string `json:"type"`
	Id       string `json:"id"`
	Name     string `json:"name"`
}

func (msg PokeMessage) Type() string {
	return "poke"
}

func init() {
	unmarshalJSONMap["poke"] = func(data []byte) (MessageElement, error) {
		var result PokeMessage
		err := json.Unmarshal(data, &result)
		return result, err
	}
}
