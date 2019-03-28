// Go in Action
// @jeffotoni
// 2019-03-10

// server.go
package server

import (
	"fmt"
	"log"
	"regexp"

	"github.com/jeffotoni/gologs/pkg/gmail"
	"github.com/jeffotoni/gologs/repo/postgres"
	// "github.com/jeffotoni/gologs/pkg/rabbitqm"
	// "github.com/jeffotoni/gologs/pkg/redis"
	// "github.com/jeffotoni/gologs/pkg/mongo"
)

var wjobs = make(chan string)

var results = make(chan string)

func init() {
	loadWorker()
}

func WProducer(jsonStr string) {
	//time.Sleep(time.Millisecond * 20)
	if len(jsonStr) <= 0 {
		return
	}
	wjobs <- jsonStr
}

func worker(id int, wjobs <-chan string, results chan<- string) {
	for j := range wjobs {
		results <- j
	}
}

func loadWorker() {
	for w := 1; w <= 3000; w++ {
		go worker(w, wjobs, results)
	}

	WConsumer()
}

func WConsumer() {

	// Here's the worker goroutine. It repeatedly receives
	// from `wjobs` with `j, okay := <-wjobs`.
	// We use this to notify on `done` when we've worked
	// all our wjobs, but never all wjobs
	go func() {
		for {
			select {
			case j := <-results:

				// wg := sync.WaitGroup{}
				// wg.Add(1)
				//if nats.Publish(j) {
				if postgres.Insert5Log(j) {
					// wg.Done()
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
							go wnotifyEmailDefault()
						}
					}
				}

				// wg.Wait()
			}
		}
	}()
}

// you can parameterize
// this function as needed
func wnotifyEmailDefault() {

	to := []string{gmail.EmailNotify}
	subject := gmail.SubjectNotify
	project := gmail.Project
	message := gmail.Message

	// if the message comes with some critical rules, send emails
	if gmail.Send(to, subject, project, message) {
		log.Println("Mail sent successfully!")
	} else {
		fmt.Println("error sending email!")
	}
}
