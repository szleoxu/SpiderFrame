package config

import (
	"fmt"
	"github.com/unknwon/goconfig"
	"os"
)

func GetConfigFile() *goconfig.ConfigFile {
	cfg, configErr := goconfig.LoadConfigFile("config/config.ini")
	if configErr != nil {
		fmt.Println(configErr)
		os.Exit(1)
	}
	return cfg
}


func GetValue(section string,key string) string {
	cfg := GetConfigFile()
	value, _ := cfg.GetValue(section, key)
	return value
}
