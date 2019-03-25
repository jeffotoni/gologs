// Back-End in Go server
// @jeffotoni
// 2019-01-04

package redis

import (
	"os"
	"strconv"
)

/////// DATA BASE
var (
	REDIS_DB_S     = os.Getenv("REDIS_DB")
	REDIS_DB       = 0
	REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
	REDIS_HOST     = os.Getenv("REDIS_HOST")
	REDIS_PORT     = os.Getenv("REDIS_PORT")
)

func init() {

	if len(REDIS_DB_S) <= 0 {
		REDIS_DB = 0
	} else {
		db, _ := strconv.Atoi(REDIS_DB_S)
		REDIS_DB = db
	}

	if len(REDIS_HOST) <= 0 {
		REDIS_HOST = "localhost"
	}

	if len(REDIS_PORT) <= 0 {
		REDIS_PORT = "6379"
	}
}
