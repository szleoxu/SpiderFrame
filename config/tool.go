package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var fileName = "config/config.ini"

func GetConfigFile() *ini.File {
	cfg, err := ini.Load(fileName)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	return cfg
}

func GetValue(section string, key string) string {
	cfg := GetConfigFile()
	return cfg.Section(section).Key(key).String()
}

func SetValue(section string, key string, value string) {
	cfg := GetConfigFile()
	cfg.Section(section).Key(key).SetValue(value)
	cfg.SaveTo(fileName)
}

func IsExist(section string, key string) bool {
	cfg := GetConfigFile()
	result := cfg.Section(section).Key(key).String()
	if result == "" {
		return false
	} else {
		return true
	}
}
