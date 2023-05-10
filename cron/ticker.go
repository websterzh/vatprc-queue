package cron

import (
	"time"
	"vatprc-queue/config"
	"vatprc-queue/sockets"
)

func InitTicker() {
	updateIntervalSecond := config.File.Section("app").Key("update_interval_second").MustInt(5)
	ticker := time.NewTicker(time.Duration(updateIntervalSecond) * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				sockets.Tick()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
