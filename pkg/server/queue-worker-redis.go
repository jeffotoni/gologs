// Go in Action
// @jeffotoni
// 2019-03-10

// server.go
package server

import (
	"log"
	"regexp"

	"github.com/jeffotoni/gologs/pkg/gmail"
	"github.com/jeffotoni/gologs/pkg/redis"
)

var rdjobs = make(chan string)

var rdresults = make(chan string)

var rdcount int

func init() {
	rdloadrdworker()
}

func RdProducer(jsonStr string) {
	//time.Sleep(time.Millisecond * 20)
	if len(jsonStr) <= 0 {
		return
	}
	rdjobs <- jsonStr
}

func rdworker(id int, rdjobs <-chan string, rdresults chan<- string) {
	for j := range rdjobs {
		rdresults <- j
	}
}

func rdloadrdworker() {
	for w := 1; w <= 3000; w++ {
		go rdworker(w, rdjobs, rdresults)
	}

	RdConsumer()
}

func RdConsumer() {

	// Here's the worker goroutine. It repeatedly receives
	// from `rdjobs` with `j, okay := <-rdjobs`.
	// We use this to notify on `done` when we've worked
	// all our rdjobs, but never all rdjobs
	go func() {
		for {
			select {
			case j := <-rdresults:

				if redis.SaveRedis(rdcount, j) {
					rdcount++
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
