package services

import (
	"context"
	_ "embed"
	"fmt"
	goredis "github.com/redis/go-redis/v9"
	"log"
	"math"
	"vatprc-queue/config"
	"vatprc-queue/redis"
)

type QueueExtra struct {
	Cid       string `json:"cid,omitempty"`
	Departure string `json:"departure,omitempty"`
	Arrival   string `json:"arrival,omitempty"`
}

type QueueResult struct {
	Status   int         `json:"status"`
	Callsign interface{} `json:"callsign"`
	Extra    *QueueExtra `json:"extra"`
}

func GetQueueResult(airport string, withExtra bool) []QueueResult {
	redisResult, err := redis.Client.ZRangeWithScores(context.Background(), airport, 0, -1).Result()
	if err != nil {
		return make([]QueueResult, 0)
	}
	result := make([]QueueResult, len(redisResult))
	for i, aircraft := range redisResult {
		var extra *QueueExtra = nil
		if withExtra {
			vatsimPilot, ok := VatsimPilots[fmt.Sprint(aircraft.Member)]
			if ok {
				extra = &QueueExtra{
					Cid:       fmt.Sprint(vatsimPilot.Cid),
					Departure: vatsimPilot.FlightPlan.Departure,
					Arrival:   vatsimPilot.FlightPlan.Arrival,
				}
			}

		}
		result[i] = QueueResult{
			Status:   int(math.Floor(aircraft.Score)),
			Callsign: aircraft.Member,
			Extra:    extra,
		}
	}

	if config.File.Section("app").Key("debug").MustBool(false) && len(result) == 0 {
		result = append(result, QueueResult{Status: 1, Callsign: "DEBUG1", Extra: &QueueExtra{Cid: "1", Departure: "TEST", Arrival: "TEST"}})
		result = append(result, QueueResult{Status: 1, Callsign: "DEBUG2", Extra: &QueueExtra{Cid: "2", Departure: "TEST", Arrival: "TEST"}})
		result = append(result, QueueResult{Status: 1, Callsign: "DEBUG3", Extra: &QueueExtra{Cid: "3", Departure: "TEST", Arrival: "TEST"}})
	}
	return result
}

func UpdateOrder(airport string, callsign string, beforeCallsign string) error {
	log.Println("reorder ", callsign, " from ", airport)
	updateOrderScript.Run(context.Background(), redis.Client, []string{airport}, callsign, beforeCallsign)
	redis.Client.ZAdd(context.Background(), "_live_check", goredis.Z{Score: 0, Member: fmt.Sprint(callsign, ":", airport)})
	return nil
}

//go:embed scripts/update_order.lua
var updateOrderScriptSrc string
var updateOrderScript = goredis.NewScript(updateOrderScriptSrc)

func UpdateStatus(airport string, callsign string, newStatus int) error {
	log.Println("update", callsign, "from", airport)
	updateStatusScript.Run(context.Background(), redis.Client, []string{airport}, callsign, newStatus)
	redis.Client.ZAdd(context.Background(), "_live_check", goredis.Z{Score: 0, Member: fmt.Sprint(callsign, ":", airport)})
	return nil
}

//go:embed scripts/update_status.lua
var updateStatusScriptSrc string
var updateStatusScript = goredis.NewScript(updateStatusScriptSrc)

func RemoveFromQueue(airport string, callsign string) {
	log.Println("removing", callsign, "from", airport)
	redis.Client.ZRem(context.Background(), airport, callsign)
}
