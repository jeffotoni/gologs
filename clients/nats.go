// Go in Action
// @jeffotoni
// 2019-03-25

package main

import (
	"log"
	"runtime"
	"time"

	"github.com/jeffotoni/gologs/repo/postgres"
	nats "github.com/nats-io/go-nats"
)

const (
	queue   = ""
	subject = "gologs"
)

func main() {

	// Create server connection
	nc, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)

	// Subscribe to subject
	log.Printf("Subscribing to subject 'gologs'\n")
	//for i := 0; i < 500000; i++ {
	// Simple Sync Subscriber
	// sub, err := nc.SubscribeSync("gologs")
	// m, err := sub.NextMsg(200)
	// if err == nil {
	// 	// 	// Handle the message
	// 	log.Printf("Subscribed message in Worker 1: %s\n", m.Data)
	// 	postgres.Insert5Log(string(m.Data))
	// }
	// sub.Unsubscribe()
	// nc.Close()

	//}

	nc.Subscribe(subject, func(msg *nats.Msg) {
		// Handle the message
		// here insert db...
		log.Printf("Received message %s\n", string(msg.Data))
		postgres.Insert5Log(string(msg.Data))
		time.Sleep(time.Millisecond * 10)
	})

	runtime.Goexit()

	// Keep the connection alive
	// Subscribe to subject
	// nc.QueueSubscribe(subject, queue, func(msg *nats.Msg) {
	// 	// 	time.Sleep(time.Millisecond * 100)
	// 	log.Printf("Received message %s\n", string(msg.Data))
	// 	// here insert db...
	// 	postgres.Insert5Log(string(msg.Data))
	// 	//eventStore := pb.EventStore{}
	// 	//err := proto.Unmarshal(msg.Data, &eventStore)
	// 	// if err == nil {
	// 	// 	// Handle the message
	// 	// 	log.Printf("Subscribed message in Worker 1: %+v\n", eventStore)
	// 	// }
	// })

	// Keep the connection alive
	// runtime.Goexit()

	// nc, err := nats.Connect(nats.DefaultURL)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// defer nc.Close()

	// Use a WaitGroup to wait for a message to arrive

	// for {
	// 	time.Sleep(time.Millisecond * 100)
	// 	wg := sync.WaitGroup{}
	// 	wg.Add(1)

	// 	// Subscribe
	// 	if _, err := nc.Subscribe("gologs", func(msg *nats.Msg) {
	// 		postgres.Insert5Log(string(msg.Data))
	// 		log.Printf("Received message %s\n", string(msg.Data))
	// 		wg.Done()
	// 	}); err != nil {
	// 		log.Println(err)
	// 		return
	// 	}

	// 	// Wait for a message to come in
	// 	wg.Wait()
	// }

	// Close the connection
	// nc.Close()
}
