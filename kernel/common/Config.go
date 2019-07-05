package common

import (
	"fmt"
	"os"

	"webserver/kernel/log"

	"github.com/BurntSushi/toml"
)

var conf *Config

// LoadConfig 載入 config
func LoadConfig() *Config {
	if conf != nil {
		return conf
	}

	// 載入config
	cf := "config/" + os.Getenv("ENV") + ".toml"
	if _, err := toml.DecodeFile(cf, &conf); err != nil {
		log.Print("error", fmt.Sprintf("toml.DecodeFile: %s\nData: %v", err.Error(), cf))
	}
	conf.ENV = os.Getenv("ENV")

	return conf
}
