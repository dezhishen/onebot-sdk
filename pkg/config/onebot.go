package config

import "fmt"

func init() {

}

func GetUrl() string {
	return fmt.Sprintf("%v://%v:%v", GetSchema(), GetIp(), GetPort())
}

func GetSchema() string {
	return "http"
}
func GetIp() string {
	return "127.0.0.1"
}

func GetPort() uint8 {
	return 80
}
