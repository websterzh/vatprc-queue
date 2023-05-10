package services

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/goccy/go-json"
	goredis "github.com/redis/go-redis/v9"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
	"vatprc-queue/config"
	"vatprc-queue/redis"
)

type VatsimController struct {
	Cid      int    `json:"cid"`
	Callsign string `json:"callsign"`
}

type VatsimPilot struct {
	Cid        int              `json:"cid"`
	Callsign   string           `json:"callsign"`
	FlightPlan VatsimFlightPlan `json:"flight_plan"`
}

type VatsimFlightPlan struct {
	Departure string `json:"departure"`
	Arrival   string `json:"arrival"`
}

type VatsimResponse struct {
	Pilots      []VatsimPilot      `json:"pilots"`
	Controllers []VatsimController `json:"controllers"`
}

var VatsimPilots map[string]VatsimPilot

func FetchOnlineDataFromVatsim() (*VatsimResponse, error) {
	url := "https://data.vatsim.net/v3/vatsim-data.json"
	timeout := config.File.Section("app").Key("fsd_fetch_timeout_second").MustInt(5)
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		log.Println("fetch VATSIM data error:", err.Error())
		return nil, err
	}

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("read VATSIM data error:", err.Error())
		return nil, err
	}

	var result VatsimResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		log.Println("parse VATSIM data error:", err.Error())
		return nil, err
	}

	VatsimPilots = make(map[string]VatsimPilot)
	for _, pilot := range result.Pilots {
		VatsimPilots[fmt.Sprint(pilot.Cid)] = pilot
	}

	log.Println("fetched VATSIM data")
	return &result, nil
}

//go:embed scripts/import_online_data.lua
var importOnlineDataScriptSrc string
var importOnlineDataScript = goredis.NewScript(importOnlineDataScriptSrc)

func AddOnlineDataToRedis(data *VatsimResponse) error {
	if data == nil {
		return nil
	}
	toAdd := make([]interface{}, 0)
	for _, pilot := range data.Pilots {
		if pilot.FlightPlan.Departure == "" {
			continue
		}
		toAdd = append(toAdd, -1, fmt.Sprint(pilot.Cid, ":", pilot.FlightPlan.Departure))
	}
	importOnlineDataScript.Run(context.Background(), redis.Client, []string{}, toAdd...)
	return nil
}

//go:embed scripts/clean_timeout.lua
var cleanTimeoutScriptSrc string
var cleanTimeoutScript = goredis.NewScript(cleanTimeoutScriptSrc)

func CleanTimeoutPilotFromQueue() {
	threshold := config.File.Section("app").Key("remove_fail_fsd_check").MustInt(30)
	remove, err := cleanTimeoutScript.Run(context.Background(), redis.Client, []string{}, threshold).StringSlice()
	if err != nil {
		return
	}
	for _, pilotAirport := range remove {
		splitResult := strings.Split(pilotAirport, ":")
		if len(splitResult) < 2 {
			continue
		}
		RemoveFromQueue(splitResult[1], splitResult[0])
	}
}
