// Go in Action
// @jeffotoni
// 2019-03-10

// server.go
package server

import (
	"log"
	"time"

	"github.com/jeffotoni/gologs/repo"
)

var jobs = make(chan string)

var done = make(chan bool)

func Publish(okay string) {
	time.Sleep(time.Millisecond * 200)
	if len(okay) <= 0 {
		return
	}
	jobs <- okay
}

func Consumer() {
	// Here's the worker goroutine. It repeatedly receives
	// from `jobs` with `j, okay := <-jobs`.
	// We use this to notify on `done` when we've worked
	// all our jobs, but never all jobs
	go func() {
		for {
			time.Sleep(time.Second * 5)
			j, okay := <-jobs
			if okay {
				if repo.InsertLog(j) {
					log.Println("save postgres")
				} else {

					log.Println("received job, error processing service send postgres: \n", j)
				}
				// here send Postgres or ElasticSearch or SQS or S3.
				// depending on the message sending email
			} else {
				// never
				log.Println("received all jobs")
				done <- true
				return
			}
		}
	}()
}
