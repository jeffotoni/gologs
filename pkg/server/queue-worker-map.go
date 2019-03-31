// Go in Action
// @jeffotoni
// 2019-03-10

// server.go
package server

import (
	"log"
	"regexp"
	"time"

	"github.com/jeffotoni/gologs/pkg/gmail"
	"github.com/jeffotoni/gologs/pkg/maps"
)

var mapjobs = make(chan string)

var mapresults = make(chan string)

var mapcount int

func init() {
	maploadmapworker()
}

func MapProducer(jsonStr string) {
	//time.Sleep(time.Millisecond * 20)
	if len(jsonStr) <= 0 {
		return
	}
	mapjobs <- jsonStr
}

func mapworker(id int, mapjobs <-chan string, mapresults chan<- string) {
	for j := range mapjobs {
		mapresults <- j
	}
}

func maploadmapworker() {
	for w := 1; w <= 3000; w++ {
		go mapworker(w, mapjobs, mapresults)
	}

	MapConsumer()
}

func MapConsumer() {

	// Here's the worker goroutine. It repeatedly receives
	// from `mapjobs` with `j, okay := <-mapjobs`.
	// We use this to notify on `done` when we've worked
	// all our mapjobs, but never all mapjobs
	go func() {
		for {
			select {
			case j := <-mapresults:
				if maps.Save(mapcount, j) {
					mapcount++
					time.Sleep(time.Millisecond * 3)
					// maps.SavePg()
					maps.SaveRedis()

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
