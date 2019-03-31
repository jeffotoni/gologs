// Go in Action
// @jeffotoni
// 2019-03-10

// server.go
package server

import (
	"log"
	"regexp"

	"github.com/jeffotoni/gologs/pkg/gmail"
	"github.com/jeffotoni/gologs/pkg/rabbitqm"
)

var rbjobs = make(chan string)

var rbresults = make(chan string)

var rabcount int

func init() {
	rbloadrbworker()
}

func RbProducer(jsonStr string) {
	//time.Sleep(time.Millisecond * 20)
	if len(jsonStr) <= 0 {
		return
	}
	rbjobs <- jsonStr
}

func rbworker(id int, rbjobs <-chan string, rbresults chan<- string) {
	for j := range rbjobs {
		rbresults <- j
	}
}

func rbloadrbworker() {
	for w := 1; w <= 3000; w++ {
		go rbworker(w, rbjobs, rbresults)
	}

	RbConsumer()
}

func RbConsumer() {

	// Here's the worker goroutine. It repeatedly receives
	// from `rbjobs` with `j, okay := <-rbjobs`.
	// We use this to notify on `done` when we've worked
	// all our rbjobs, but never all rbjobs
	go func() {
		for {
			select {
			case j := <-rbresults:
				if rabbitqm.PublishQueue(rabcount, j) {
					rabcount++
					// only enabled
					if len(gmail.GmailUser) > 0 &&
						len(gmail.GmailPassword) > 0 &&
						len(gmail.EmailNotify) > 0 {
						// log.Println("save postgres")
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
			}
		}
	}()
}
