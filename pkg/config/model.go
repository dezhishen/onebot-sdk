package config

type OnebotApiConfig struct {
	Type        string `json:"type" yaml:"type"`         // http, websocket
	Endpoint    string `json:"endpoint" yaml:"endpoint"` // http://
	Secret      string `json:"secret" yaml:"secret"`
	AccessToken string `json:"access_token" yaml:"access-token"`
	Timeout     int64  `json:"timeout" yaml:"timeout"` // ms
}

type OnebotEventConfig struct {
	Type        string `json:"type" yaml:"type"` // http-reverse,websocket,websocket-reverse
	Addr        string `json:"addr" yaml:"addr"` // http://\
	Secret      string `json:"secret" yaml:"secret"`
	AccessToken string `json:"access_token" yaml:"access-token"`
}

type OnebotConfig struct {
	Api   *OnebotApiConfig   `json:"api" yaml:"api"`
	Event *OnebotEventConfig `json:"event" yaml:"event"`
}
