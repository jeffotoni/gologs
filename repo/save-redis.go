// Back-End in Go server
// @jeffotoni
// 2019-01-04

package repo

import (
	"log"
	"strconv"

	"github.com/jeffotoni/gologs/pkg/redis"
)

func SaveRedis(key_int int, value string) {
	client := redis.NewClient()
	key := strconv.Itoa(key_int)
	err := client.Set(key, value, 0).Err()
	if err != nil {
		log.Println("redis:: ", err)
		return
	}
}
