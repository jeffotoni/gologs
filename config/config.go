// Go in Action
// @jeffotoni
// 2019-03-10

package config

import (
	"os"
	"strconv"
	"strings"
)

var MEMORY_S = os.Getenv("MEMORY")
var MEMORY int
var DEBUG_S = os.Getenv("DEBUG")
var DEBUG_REQ_S = os.Getenv("DEBUG_REQ")
var DEBUG bool
var DEBUG_REQ int

const (
	POSTGRES = "postgres"
	REDIS    = "redis"
	RABBITQM = "rabbitqm"
	MONGO    = "mongo"
	NATS     = "nats"
	MAPS     = "maps"
)

// postgres
// redis
// rabbitqm
// mongo
// nats
// Default postgres
var SERVICE = os.Getenv("SERVICE")

// NATS PERSISTENT default redis
var NATS_PERSISTENT = os.Getenv("NATS_PERSISTENT")

func init() {

	// services accept
	// Default postgres
	if len(SERVICE) <= 0 {
		SERVICE = "postgres"
	}

	// remove, tolower
	SERVICE = strings.Trim(strings.ToLower(SERVICE), " ")

	if len(DEBUG_S) <= 0 {
		DEBUG = false
	} else {
		DEBUG, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	}

	if len(DEBUG_REQ_S) > 0 {
		DEBUG_REQ, _ = strconv.Atoi(DEBUG_REQ_S)
	} else {
		DEBUG_REQ = 0
	}

	if len(MEMORY_S) > 0 {
		MEMORY, _ = strconv.Atoi(MEMORY_S)
	} else {
		MEMORY = 20000
	}

	// default
	if len(NATS_PERSISTENT) <= 0 {
		NATS_PERSISTENT = "redis"
	}

	// show
	if DEBUG {
		println("\033[0;31mServer Run mode Debug: " + DEBUG_REQ_S + "/req ...\033[0;0m")
	} else {
		println("\033[0;33mServer Run...\033[0;0m")
	}
}
