// Go in Action
// @jeffotoni
// 2019-03-25

package main

import (
	"log"
	"sync"

	nats "github.com/nats-io/go-nats"
)

func main() {

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Println(err)
		return
	}
	defer nc.Close()

	// Use a WaitGroup to wait for a message to arrive
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Subscribe
	if _, err := nc.Subscribe("gologs", func(m *nats.Msg) {
		println(string(m.Data))
		wg.Done()
	}); err != nil {
		log.Println(err)
		return
	}

	// Wait for a message to come in
	wg.Wait()

	// Close the connection
	nc.Close()
}
