package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
	"vatprc-queue/config"
	"vatprc-queue/redis"
)

func CreateToken(scope string) string {
	newToken := strings.ReplaceAll(uuid.New().String(), "-", "")
	exp := config.File.Section("app").Key("token_ttl_minute").MustInt(60)
	redis.Client.Set(context.Background(), getTokenKey(newToken), scope, time.Duration(exp)*time.Minute)
	return newToken
}

func DeleteToken(token string) {
	redis.Client.Del(context.Background(), getTokenKey(token))
}

func HasToken(token string, airport string) bool {
	// TODO: Check airport
	exist, err := redis.Client.Exists(context.Background(), getTokenKey(token)).Result()
	if err != nil {
		return false
	}
	return exist == 1
}

func getTokenKey(token string) string {
	return fmt.Sprint("token_", strings.ToLower(token))
}
