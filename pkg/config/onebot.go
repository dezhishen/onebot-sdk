package config

import "fmt"

func init() {

}

func GetHttpUrl() string {
	return fmt.Sprintf("%v://%v:%v", GetHttpSchema(), GetHttpHost(), GetHttpPort())
}

func GetHttpSchema() string {
	return "http"
}
func GetHttpHost() string {
	return "127.0.0.1"
}

func GetHttpPort() uint {
	return 5700
}

func GetWsHost() string {
	return "127.0.0.1"
}

func GetWsPort() uint {
	return 6700
}
