package model

import "encoding/json"

type MessageElementPoke struct {
	PokeType string `json:"type"`
	Id       string `json:"id"`
	Name     string `json:"name"`
}

func (msg MessageElementPoke) Type() string {
	return "poke"
}

func init() {
	unmarshalJSONMap["poke"] = func(data []byte) (MessageElement, error) {
		var result MessageElementPoke
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
