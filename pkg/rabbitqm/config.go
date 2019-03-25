// Back-End in Go server
// @jeffotoni
// 2019-01-04

package rabbitqm

import (
	"os"
)

/////// DATA BASE
var (
	RABBI_USER     = os.Getenv("RABBI_USER")
	RABBI_PASSWORD = os.Getenv("RABBI_PASSWORD")
	RABBI_HOST     = os.Getenv("RABBI_HOST")
	RABBI_PORT     = os.Getenv("RABBI_PORT")
)

func init() {

	if len(RABBI_USER) <= 0 {
		RABBI_USER = "guest"
	}

	if len(RABBI_HOST) <= 0 {
		RABBI_HOST = "localhost"
	}

	if len(RABBI_PORT) <= 0 {
		RABBI_PORT = "5672"
	}
}
