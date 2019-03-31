// Back-End in Go server
// @jeffotoni
// 2019-01-04

package rabbitqm

import (
	"os"
	"strconv"
)

/////// DATA BASE
var (
	RABBI_USER     = os.Getenv("RABBI_USER")
	RABBI_PASSWORD = os.Getenv("RABBI_PASSWORD")
	RABBI_HOST     = os.Getenv("RABBI_HOST")
	RABBI_PORT     = os.Getenv("RABBI_PORT")

	RABBI_PROTOCOLO     = os.Getenv("RABBI_PROTOCOLO")
	RABBI_EXCHANGE_NAME = os.Getenv("RABBI_EXCHANGE_NAME")
	RABBI_EXCHANGE_TYPE = os.Getenv("RABBI_EXCHANGE_TYPE")
	RABBI_RELIABLE_S    = os.Getenv("RABBI_RELIABLE")
	RABBI_RELIABLE      = true
	RABBI_DIAL          = ""
)

// var (
// 	uri          = flag.String("uri", "amqp://guest:guest@localhost:5672/", "AMQP URI")
// 	exchangeName = flag.String("exchange", "test-exchange", "Durable AMQP exchange name")
// 	exchangeType = flag.String("exchange-type", "direct", "Exchange type - direct|fanout|topic|x-custom")
// 	//routingKey   = flag.String("key", "test-key", "AMQP routing key")
// 	//body     = flag.String("body", "foobar", "Body of message")
// 	reliable = flag.Bool("reliable", true, "Wait for the publisher confirmation before exiting")
// )

func init() {

	if len(RABBI_PROTOCOLO) <= 0 {
		RABBI_PROTOCOLO = "amqp://"
	}

	if len(RABBI_USER) <= 0 {
		RABBI_USER = "guest"
	}

	if len(RABBI_PASSWORD) <= 0 {
		RABBI_PASSWORD = "guest"
	}

	if len(RABBI_HOST) <= 0 {
		RABBI_HOST = "localhost"
	}

	if len(RABBI_PORT) <= 0 {
		RABBI_PORT = "5672"
	}

	if len(RABBI_EXCHANGE_NAME) <= 0 {
		RABBI_EXCHANGE_NAME = "test-exchange"
	}

	if len(RABBI_EXCHANGE_TYPE) <= 0 {
		RABBI_EXCHANGE_TYPE = "exchange-type"
	}

	if len(RABBI_RELIABLE_S) <= 0 {
		RABBI_RELIABLE, _ = strconv.ParseBool(RABBI_RELIABLE_S)
	}

	RABBI_DIAL = RABBI_PROTOCOLO + RABBI_USER + ":" + RABBI_PASSWORD + "@" + RABBI_HOST + ":" + RABBI_PORT + "/"
}
