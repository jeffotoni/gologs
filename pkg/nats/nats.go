package nats

import (
	"log"
	"sync"

	"github.com/jeffotoni/gologs/pkg/postgres"
	nats "github.com/nats-io/go-nats"
)

const (
	queue   = ""
	Subject = "gologs"
	limit   = 10000000 // 10millions
)

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

func SubscribeAsync() {

	//start := time.Now()
	chanpg := make(chan string, limit)

	// Create server connection
	nc, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)
	// Subscribe to subject
	log.Printf("Subscribing to subject 'gologs'\n")
	defer nc.Close()

	go func() {
		// var count int
		// Use a WaitGroup to wait for a message to arrive
		wg := sync.WaitGroup{}
		wg.Add(limit)

		// Subscribe
		if _, err := nc.Subscribe(Subject, func(msg *nats.Msg) {
			log.Printf("Received message %s\n", string(msg.Data))
			chanpg <- string(msg.Data)
			//redis.SaveRedis(count, string(msg.Data))
			// count++
			wg.Done()
		}); err != nil {
			log.Fatal(err)
		}

		// Wait for a message to come in
		wg.Wait()

		// Close the connection
		nc.Close()

	}()
	// close(chanpg)
	//go func() {

	for {
		select {
		case cmsgJson := <-chanpg:
			postgres.Insert5Log(cmsgJson)
		}
	}

}
