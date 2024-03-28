package redis

import (
	"ChatDemo/global"
	"fmt"

	"github.com/go-redis/redis"
)

func InitRedis() *redis.Client {
	config := global.Config.RedisConfig
	RDB := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.Db,
	})

	return RDB
}

var RDB = InitRedis()
