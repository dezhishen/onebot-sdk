package config

type OnebotConfig struct {
	Type        string `json:"type" yaml:"type"`         // http, websocket
	Endpoint    string `json:"endpoint" yaml:"endpoint"` // http://
	AccessToken string `json:"access_token" yaml:"access-token"`
}
