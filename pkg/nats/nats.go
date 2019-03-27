package nats

import (
	"log"

	nats "github.com/nats-io/go-nats"
)

const Subject = "gologs"

func Publish(jsonStr string) bool {

	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)
	if err := nc.Publish(Subject, []byte(jsonStr)); err != nil {
		log.Println(err)
	}
	defer nc.Close()
	nc.Flush()
	return true
}

// func Subscribe() {

// }
