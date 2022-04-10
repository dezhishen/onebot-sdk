package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const configPath = "./configs"

const configName = "bot"

var all_config = make(map[string]*viper.Viper)

func init() {
	loadConfig(configName)
}

func loadConfig(name string) *viper.Viper {
	exists := pathExists(configPath)
	if !exists {
		return nil
	}
	c := viper.New()
	c.AddConfigPath("./configs")
	c.SetConfigName(name)
	c.SetConfigType("yaml")
	if err := c.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("读取配置文件错误:%v", err))
	}
	all_config[name] = c
	return c
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func GetViper() *viper.Viper {
	if c, ok := all_config[configName]; ok {
		return c
	}
	return nil
}
