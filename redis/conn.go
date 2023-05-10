package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"vatprc-queue/config"
)

var Client *redis.Client

func InitRedis() {
	addr := config.File.Section("redis").Key("address").MustString("localhost")
	port := config.File.Section("redis").Key("port").MustInt(6379)
	password := config.File.Section("redis").Key("password").MustString("")
	database := config.File.Section("redis").Key("database").MustInt(0)
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", addr, port),
		Password: password,
		DB:       database,
	})
}
