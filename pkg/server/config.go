// Go in Action
// @jeffotoni
// 2019-03-10

package server

import (
	"os"
	"strconv"
)

var MEMORY_S = os.Getenv("MEMORY")
var MEMORY int
var DEBUG_S = os.Getenv("DEBUG")
var DEBUG_REQ_S = os.Getenv("DEBUG_REQ")
var DEBUG bool
var DEBUG_REQ int

// postgres
// redis
// rabbitqm
// Default postgres
var SERVICE = os.Getenv("SERVICE")

func init() {

	// services accept
	// Default postgres
	if len(SERVICE) <= 0 {
		SERVICE = "postgres"
	}

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

	// show
	if DEBUG {
		println("\033[0;31mServer Run mode Debug: " + DEBUG_REQ_S + "/req ...\033[0;0m")
	} else {
		println("\033[0;33mServer Run...\033[0;0m")
	}
}
