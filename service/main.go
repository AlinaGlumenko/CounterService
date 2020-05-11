package main

import (
	"encoding/json"
	"log"
	"net/http"

	"./RedisDB"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/count_users", getCounter).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getCounter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	key := r.URL.Query().Get("pagename")
	visit := r.URL.Query().Get("visit")

	counter := 0

	pool := RedisDB.NewPool()
	conn := pool.Get()
	if RedisDB.Exists(conn, key) {
		counter = RedisDB.Get(conn, key)
	}
	if visit == "true" && key != "" {
		counter++
		err := RedisDB.Set(conn, key, counter)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	json.NewEncoder(w).Encode(counter)
}
