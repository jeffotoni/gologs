// Go in Action
// @jeffotoni
// 2019-03-10

// server.go
package server

import (
	"log"
	"regexp"

	"github.com/jeffotoni/gologs/pkg/gmail"
	"github.com/jeffotoni/gologs/pkg/postgres"
)

var pgjobs = make(chan string)

var pgresults = make(chan string)

func init() {
	pgloadpgworker()
}

func PgProducer(jsonStr string) {
	//time.Sleep(time.Millisecond * 20)
	if len(jsonStr) <= 0 {
		return
	}
	pgjobs <- jsonStr
}

func pgworker(id int, pgjobs <-chan string, pgresults chan<- string) {
	for j := range pgjobs {
		pgresults <- j
	}
}

func pgloadpgworker() {
	for w := 1; w <= 3000; w++ {
		go pgworker(w, pgjobs, pgresults)
	}

	PgConsumer()
}

func PgConsumer() {

	// Here's the worker goroutine. It repeatedly receives
	// from `pgjobs` with `j, okay := <-pgjobs`.
	// We use this to notify on `done` when we've worked
	// all our pgjobs, but never all pgjobs
	go func() {
		for {
			select {
			case j := <-pgresults:

				// wg := sync.WaitGroup{}
				// wg.Add(1)
				//if nats.Publish(j) {
				if postgres.Insert5Log(j) {
					// wg.Done()
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
				// wg.Wait()
			}
		}
	}()
}
