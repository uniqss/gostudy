package config_viper

import "github.com/spf13/viper"

type Config struct {
	RedisHost    string
	RedisPort   int
	RedisDb     int
	RedisPassword string
}

func NewDefaultConfig() *Config {
	return &Config{
		RedisHost:    "127.0.0.1",
		RedisPort:   6379,
		RedisDb:     0,
		RedisPassword: "",
	}
}

func (cfg *Config) initConfig() {
	if v := viper.GetString("redis.host"); v != "" {
		cfg.RedisHost = v
	}
	if v := viper.GetInt("redis.port"); v != 0 {
		cfg.RedisPort = v
	}
	if v := viper.GetInt("redis.db"); v != 0 {
		cfg.RedisDb = v
	}
	if v := viper.GetString("redis.password"); v != "" {
		cfg.RedisPassword = v
	}
}
