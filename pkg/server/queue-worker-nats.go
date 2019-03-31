// Go in Action
// @jeffotoni
// 2019-03-10

// server.go
package server

import (
	"log"
	"regexp"
	"sync"

	"github.com/jeffotoni/gologs/pkg/gmail"
	"github.com/jeffotoni/gologs/pkg/nats"
)

var natsjobs = make(chan string)

var natsresults = make(chan string)

func init() {
	natsloadnatsworker()
}

func NatsProducer(jsonStr string) {
	//time.Sleep(time.Millisecond * 20)
	if len(jsonStr) <= 0 {
		return
	}
	natsjobs <- jsonStr
}

func natsworker(id int, natsjobs <-chan string, natsresults chan<- string) {
	for j := range natsjobs {
		natsresults <- j
	}
}

func natsloadnatsworker() {
	for w := 1; w <= 3000; w++ {
		go natsworker(w, natsjobs, natsresults)
	}

	NatsConsumer()
}

func NatsConsumer() {

	// Here's the worker goroutine. It repeatedly receives
	// from `natsjobs` with `j, okay := <-natsjobs`.
	// We use this to notify on `done` when we've worked
	// all our natsjobs, but never all natsjobs
	go func() {
		for {
			select {
			case j := <-natsresults:

				wg := sync.WaitGroup{}
				wg.Add(1)
				if nats.Publish(j) {
					wg.Done()
					// only enabled
					if len(gmail.GmailUser) > 0 &&
						len(gmail.GmailPassword) > 0 &&
						len(gmail.EmailNotify) > 0 {
						// rule critical
						matched, err := regexp.MatchString("#critical#", j)
						if err != nil {
							log.Println("Error regexp critical", err)
						}
						if matched {
							notifyEmailDefault()
						}
					}
				}
				wg.Wait()
			}
		}
	}()
}
