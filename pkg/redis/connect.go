// Back-End in Go server
// @jeffotoni
// 2019-01-04

package redis

import (
	"log"

	"github.com/go-redis/redis"
)

func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()

	if pong != "PONG" || err != nil {
		log.Println("Redis error: ", err)
	}

	return client
	//fmt.Println(pong, err)
	// Output: PONG <nil>
}
