package services

import (
	"context"
	_ "embed"
	"fmt"
	goredis "github.com/redis/go-redis/v9"
	"log"
	"math"
	"vatprc-queue/redis"
)

type QueueExtra struct {
	Callsign  string `json:"callsign,omitempty"`
	Departure string `json:"departure,omitempty"`
	Arrival   string `json:"arrival,omitempty"`
}

type QueueResult struct {
	Status int         `json:"status"`
	Cid    interface{} `json:"cid"`
	Extra  *QueueExtra `json:"extra"`
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
					Callsign:  vatsimPilot.Callsign,
					Departure: vatsimPilot.FlightPlan.Departure,
					Arrival:   vatsimPilot.FlightPlan.Arrival,
				}
			}

		}
		result[i] = QueueResult{
			Status: int(math.Floor(aircraft.Score)),
			Cid:    aircraft.Member,
			Extra:  extra,
		}
	}
	return result
}

func UpdateOrder(airport string, cid string, beforeCid string) error {
	log.Println("reorder ", cid, " from ", airport)
	updateOrderScript.Run(context.Background(), redis.Client, []string{airport}, cid, beforeCid)
	redis.Client.ZAdd(context.Background(), "_live_check", goredis.Z{Score: 0, Member: fmt.Sprint(cid, ":", airport)})
	return nil
}

//go:embed scripts/update_order.lua
var updateOrderScriptSrc string
var updateOrderScript = goredis.NewScript(updateOrderScriptSrc)

func UpdateStatus(airport string, cid string, newStatus int) error {
	log.Println("update ", cid, " from ", airport)
	updateStatusScript.Run(context.Background(), redis.Client, []string{airport}, cid, newStatus)
	redis.Client.ZAdd(context.Background(), "_live_check", goredis.Z{Score: 0, Member: fmt.Sprint(cid, ":", airport)})
	return nil
}

//go:embed scripts/update_status.lua
var updateStatusScriptSrc string
var updateStatusScript = goredis.NewScript(updateStatusScriptSrc)

func RemoveFromQueue(airport string, cid string) {
	log.Println("removing", cid, "from", airport)
	redis.Client.ZRem(context.Background(), airport, cid)
}
