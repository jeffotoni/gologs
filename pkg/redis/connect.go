// Back-End in Go server
// @jeffotoni
// 2019-01-04

package redis

import (
	"log"
	"strconv"

	"github.com/go-redis/redis"
)

var client *redis.Client

func init() {
	client = NewClient()
}

func NewClient() *redis.Client {
	c := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// pong, err := client.Ping().Result()
	// if pong != "PONG" || err != nil {
	// 	log.Println("Redis error: ", err)
	// 	return nil
	// }

	return c
	//fmt.Println(pong, err)
	// Output: PONG <nil>
}

func SaveRedis(key_int int, value string) bool {
	key := strconv.Itoa(key_int)
	err := client.Set(key, value, 0).Err()
	if err != nil {
		log.Println("redis:: ", err)
		return false
	}

	retur true
}
