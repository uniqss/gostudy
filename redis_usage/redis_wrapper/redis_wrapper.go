package redis_wrapper

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"redis_usage/config_viper"
	"strconv"
)

func factory(name string) *redis.Client {
	var cfg = config_viper.GetCfg()
	fmt.Printf("RedisHost:%s RedisPort:%d RedisPassword:%s\r\n", cfg.RedisHost, cfg.RedisPort, cfg.RedisPassword)

	redisOptions := &redis.Options{}
	redisOptions.Addr = cfg.RedisHost + ":" + strconv.Itoa(cfg.RedisPort)

	if v := cfg.RedisPassword; v != "" {
		redisOptions.Password = v
	}

	newClient := redis.NewClient(redisOptions)

	_, err := newClient.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("newClient.Ping failed. err:", err)
		return nil
	}
	return newClient
}
func GetRedisConnClient(name string) *redis.Client {
	return factory(name)
}
