package RedisDB

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
)

var c redis.Conn

func getConn() redis.Conn {
	log.Trace("Func getConn() started.")
	if c == nil {
		log.Info("Creating new connection to Redis.")
		pool := NewPool()
		c = pool.Get()
		log.Info("Got connection to Redis.")
	}
	log.Trace("Func getConn() finished.")
	return c
}

func NewPool() *redis.Pool {
	log.Trace("Getting pool from Redis...")
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				log.Fatal(err.Error())
			}
			log.Info("Got pool from Redis.")
			return c, err
		},
	}
}

func Set(key string, value int) error {
	log.Trace("Trying to Set() data...")
	_, err := getConn().Do("SET", key, value)
	if err != nil {
		log.Warn("Error of setting data.")
		return err
	}
	log.Info("Data is set.")
	log.Trace("Func Set() finished.")
	return nil
}

func Get(key string) int {
	log.Trace("Trying to Get() data...")
	count, err := redis.Int(getConn().Do("GET", key))
	if err == redis.ErrNil {
		log.Warn(key + " does not exist\n")
		return 0
	} else if err != nil {
		log.Error(err.Error())
		return 0
	}
	log.Info("Data is got.")
	log.Trace("Func Get() finished.")
	return count
}

func GetAll() {
	log.Debug("Trying to GetAll() data...")
	keys, err := redis.Strings(getConn().Do("KEYS", "*"))
	if err != nil {
		log.Error(err.Error())
	}
	log.Info("All data is got.")
	for _, key := range keys {
		fmt.Println(key, "=", Get(key))
	}
}

func Exists(key string) bool {
	log.Trace("Trying to check if data Exists()...")
	exists, err := redis.Bool(getConn().Do("EXISTS", key))
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Info("Data is checked.")
	log.Trace("Func Exists() finished.")
	return exists
}
