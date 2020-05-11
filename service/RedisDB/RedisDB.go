package RedisDB

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

func NewPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				log.Fatal(err.Error())
			}
			return c, err
		},
	}
}

func Set(c redis.Conn, key string, value int) error {
	_, err := c.Do("SET", key, value)
	if err != nil {
		return err
	}
	return nil
}

func Get(c redis.Conn, key string) int {
	count, err := redis.Int(c.Do("GET", key))
	if err == redis.ErrNil {
		fmt.Printf("%s does not exist\n", key)
		return 0
	} else if err != nil {
		log.Fatal(err.Error())
	}
	return count
}

func GetAll(c redis.Conn) {
	keys, err := redis.Strings(c.Do("KEYS", "*"))
	if err != nil {
		fmt.Println(err)
	}
	for _, key := range keys {
		fmt.Println(key, "=", Get(c, key))
	}
}

func Exists(c redis.Conn, key string) bool {
	exists, err := redis.Bool(c.Do("EXISTS", key))
	if err != nil {
		log.Fatal(err.Error())
	}
	return exists
}
