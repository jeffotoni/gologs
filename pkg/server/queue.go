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
	"github.com/jeffotoni/gologs/repo"
)

var jobs = make(chan string)

var done = make(chan bool)

var count, count2 int

func Publish(okay string) {
	//time.Sleep(time.Millisecond * 20)
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
			// time.Sleep(time.Second * 2)
			j, okay := <-jobs
			if okay {
				if repo.Map(count, j) {
					//if repo.Insert5Log(j) {
					//if true {
					// Just for debug
					// And test
					if DEBUG {
						count++
						if count == 1 {
							log.Println("start save Map")
						}
						if count == DEBUG_REQ {
							log.Println("fim save Map Qtn:", count)
							// start save db ..
							log.Println("start save Postgres")
							go repo.SavePg()
							count = 0
						}
					} else {
						// if prod
						count++
						count2++
						if count2 == MEMORY {
							go repo.SavePg()
							count2 = 0
						}
					}

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
							go notifyEmailDefault()
						}
					}

				} else {
					log.Println("received job, error processing service send Map: \n", j)
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

// you can parameterize
// this function as needed
func notifyEmailDefault() {

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
