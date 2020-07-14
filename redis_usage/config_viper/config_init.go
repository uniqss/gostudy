package config_viper

import (
	"fmt"
	"github.com/spf13/viper"
)

var cfg *Config

func GetCfg() *Config {
	return cfg
}

func init() {
	cfg = NewDefaultConfig()
	configFileName := "serverconfig.json"
	viper.SetConfigFile(configFileName)
	if err := viper.MergeInConfig(); err != nil {
		fmt.Println("read server config failed. use default config. server config:", configFileName)
		return
	}
	cfg.initConfig()
}
