package main

import (
	"encoding/json"
	"net/http"

	"./Logging"
	"./RedisDB"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	Logging.LogInit()
	r := mux.NewRouter()
	r.HandleFunc("/api/count_users", getCounter).Methods("GET")
	log.Debug("Starting listen and serve...")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getCounter(w http.ResponseWriter, r *http.Request) {
	log.Trace("Start function getCounter()")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	log.Info("URL query: " + r.URL.RawQuery)
	key := r.URL.Query().Get("pagename")
	visit := r.URL.Query().Get("visit")

	counter := 0

	if RedisDB.Exists(key) {
		counter = RedisDB.Get(key)
	}
	if visit == "true" && key != "" {
		log.Debug("If visit = true, then counter++ and set data to DB.")
		counter++
		err := RedisDB.Set(key, counter)
		if err != nil {
			log.Error(err.Error())
		}
	}

	json.NewEncoder(w).Encode(counter)
	log.Info("Count is got.")
	log.Trace("Function getCounter() is finished.")
}
