package config

type httpConfig struct {
	Host   string `json:"host"`
	Schema string `json:"schema"`
	Port   int    `json:"port"`
}

type webSocketConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type oneBotConfig struct {
	Http      *httpConfig      `json:"http"`
	Websocket *webSocketConfig `json:"websocket"`
}
