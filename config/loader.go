package config

import (
	"gopkg.in/ini.v1"
	"log"
)

var File *ini.File

func LoadConfig() {
	var err error
	File, err = ini.Load("config/config.ini")
	if err != nil {
		log.Fatalf("%v -> %v\n", "loading config file", err)
	}
}
