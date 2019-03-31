// Go in Action
// @jeffotoni
// 2019-03-10

// server.go
package server

import (
	"log"
	"regexp"

	"github.com/jeffotoni/gologs/pkg/gmail"
	"github.com/jeffotoni/gologs/pkg/mongo"
)

var mgjobs = make(chan string)

var mgresults = make(chan string)

var mgbcount int

func init() {
	mgloadMgworker()
}

func MgProducer(jsonStr string) {
	//time.Sleep(time.Millisecond * 20)
	if len(jsonStr) <= 0 {
		return
	}
	mgjobs <- jsonStr
}

func Mgworker(id int, mgjobs <-chan string, mgresults chan<- string) {
	for j := range mgjobs {
		mgresults <- j
	}
}

func mgloadMgworker() {
	for w := 1; w <= 3000; w++ {
		go Mgworker(w, mgjobs, mgresults)
	}

	MgConsumer()
}

func MgConsumer() {

	// Here's the worker goroutine. It repeatedly receives
	// from `mgjobs` with `j, okay := <-mgjobs`.
	// We use this to notify on `done` when we've worked
	// all our mgjobs, but never all mgjobs
	go func() {
		for {
			select {
			case j := <-mgresults:
				if mongo.InsertOne(mgbcount, j) {
					mgbcount++
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
