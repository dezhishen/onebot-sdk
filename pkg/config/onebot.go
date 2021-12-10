package config

import (
	"fmt"
	"strconv"
)

func getConfigStringWithDefault(group, key, defaultValue string) string {
	c, ok := all_config[configName]
	if !ok {
		return defaultValue
	}
	d := c.GetStringMapString(group)
	if d == nil {
		return defaultValue
	}
	if v, ok := d[key]; ok {
		return v
	}
	return defaultValue
}

func getConfigIntWithDefault(group, key string, defaultValue int) int {
	value := getConfigStringWithDefault(group, key, strconv.Itoa(defaultValue))
	v, _ := strconv.ParseInt(value, 10, 64)
	return int(v)
}

func GetHttpUrl() string {
	return fmt.Sprintf("%v://%v:%v", GetHttpSchema(), GetHttpHost(), GetHttpPort())
}

func GetHttpSchema() string {
	return getConfigStringWithDefault("http", "schema", "http")
}
func GetHttpHost() string {
	return getConfigStringWithDefault("http", "host", "127.0.0.1")
}

func GetHttpPort() uint {
	return uint(getConfigIntWithDefault("http", "host", 5700))
}

func GetWsHost() string {
	return getConfigStringWithDefault("websocket", "host", "127.0.0.1")
}

func GetWsPort() uint {
	return uint(getConfigIntWithDefault("websocket", "host", 6700))
}
