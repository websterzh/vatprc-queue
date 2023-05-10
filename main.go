package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"vatprc-queue/config"
	"vatprc-queue/cron"
	"vatprc-queue/gin/router"
	"vatprc-queue/redis"
	"vatprc-queue/sockets"
)

func init() {
	config.LoadConfig()
	redis.InitRedis()
	cron.InitTicker()
	cron.InitFsdFetcher()
}

func main() {
	port := config.File.Section("app").Key("port").MustInt(8080)
	debug := config.File.Section("app").Key("port").MustBool(false)
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	app := router.InitRouter()
	app.Run(fmt.Sprintf(":%d", port))
}

var pool *sockets.ClientPool
