package cron

import (
	"time"
	"vatprc-queue/config"
	"vatprc-queue/gin/services"
)

func InitFsdFetcher() {
	fsdCheckEnabled := config.File.Section("app").Key("enable_fsd_check").MustBool(true)
	if !fsdCheckEnabled {
		return
	}
	fsdFetchInterval := config.File.Section("app").Key("fsd_fetch_interval_second").MustInt(60)
	ticker := time.NewTicker(time.Duration(fsdFetchInterval) * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				vatsimResponse, err := services.FetchOnlineDataFromVatsim()
				if err != nil {
					continue
				}
				_ = services.AddOnlineDataToRedis(vatsimResponse)
				services.CleanTimeoutPilotFromQueue()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
