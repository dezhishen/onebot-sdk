package model

import "encoding/json"

type MessageElementPoke struct {
	PokeType string `json:"type"`
	Id       string `json:"id"`
	Name     string `json:"name"`
}

func (msg *MessageElementPoke) Type() string {
	return "poke"
}

func (msg *MessageElementPoke) ToGRPC() *MessageElementPokeGRPC {
	return &MessageElementPokeGRPC{
		PokeType: msg.PokeType,
		Id:       msg.Id,
		Name:     msg.Name,
	}
}

func (msg *MessageElementPokeGRPC) ToStruct() *MessageElementPoke {
	return &MessageElementPoke{
		PokeType: msg.PokeType,
		Id:       msg.Id,
		Name:     msg.Name,
	}
}

func init() {
	messageElementUnmarshalJSONMap["poke"] = func(data []byte) (MessageElement, error) {
		var result MessageElementPoke
		err := json.Unmarshal(data, &result)
		return &result, err
	}
}
