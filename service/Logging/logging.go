package Logging

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func LogInit() {
	file, err := os.OpenFile("./logs/info.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}
