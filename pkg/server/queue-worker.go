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
	//"github.com/jeffotoni/gologs/repo/postgres"
	// "github.com/jeffotoni/gologs/pkg/rabbitqm"
	// "github.com/jeffotoni/gologs/pkg/redis"
	// "github.com/jeffotoni/gologs/pkg/mongo"
	"github.com/jeffotoni/gologs/pkg/nats"
)

const limit = 10000000 // 10millions

var jobs = make(chan string, limit)

var results = make(chan string, limit)

var done = make(chan bool)

var count int

var count2 int

func Producer(okay string) {
	//time.Sleep(time.Millisecond * 20)
	if len(okay) <= 0 {
		return
	}
	jobs <- okay
}

func worker(id int, jobs <-chan string, results chan<- string) {
	for j := range jobs {
		results <- j
	}
}

func Consumer2() {
	for w := 1; w <= 6000; w++ {
		go worker(w, jobs, results)
	}
}

func Consumer() {

	// Here's the worker goroutine. It repeatedly receives
	// from `jobs` with `j, okay := <-jobs`.
	// We use this to notify on `done` when we've worked
	// all our jobs, but never all jobs
	go func() {
		for {
			select {
			case j := <-jobs:

				// wg := sync.WaitGroup{}
				// wg.Add(1)

				if nats.Publish(j) {

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
							go notifyEmailDefault()
						}
					}
				}

				// wg.Wait()
			}

			// time.Sleep(time.Second * 2)
			//j, okay := <-jobs
			// if okay {
			// 	//if true {
			// 	// if repo.Map(count, j) {
			// 	// if postgres.Insert5Log(j) {
			// 	// if redis.SaveRedis(count, j) {
			// 	// if rabbitqm.SendV2(count, j) {
			// 	// if mongo.InsertOne(count, j) {
			// 	if nats.Publish(j) {
			// 		//if true {
			// 		// Just for debug
			// 		// And test
			// 		// if DEBUG {
			// 		// count++
			// 		// count2++

			// 		// if count2 == 20000 {
			// 		// 	time.Sleep(time.Second * 5)
			// 		// 	count2 = 0
			// 		// }
			// 		// 	if count == 1 {
			// 		// 		log.Println("start save Map")
			// 		// 	}
			// 		// 	if count == DEBUG_REQ {
			// 		// 		log.Println("fim save Map Qtn:", count)
			// 		// 		// start save db ..
			// 		// 		log.Println("start save Postgres")
			// 		// 		go repo.SavePg()
			// 		// 		count = 0
			// 		// 	}
			// 		// } else {
			// 		// 	//repo.Insert5Log(j)
			// 		// 	// if prod
			// 		// 	count++
			// 		// 	count2++
			// 		// 	if count2 == MEMORY {
			// 		// 		log.Println("start save Redis!")
			// 		// 		go repo.SaveRedis()
			// 		// 		count2 = 0
			// 		// 		time.Sleep(time.Millisecond * 3000)
			// 		// 	}
			// 		// }

			// 		if len(gmail.GmailUser) > 0 &&
			// 			len(gmail.GmailPassword) > 0 &&
			// 			len(gmail.EmailNotify) > 0 {
			// 			// log.Println("save postgres")
			// 			// rule critical
			// 			matched, err := regexp.MatchString("#critical#", j)
			// 			if err != nil {
			// 				log.Println("Error regexp critical", err)
			// 			}
			// 			if matched {
			// 				go notifyEmailDefault()
			// 			}
			// 		}

			// 	} else {
			// 		log.Println("received job, error processing service send Map: \n", j)
			// 	}
			// 	// here send Postgres or ElasticSearch or SQS or S3.
			// 	// depending on the message sending email
			// } else {
			// 	// never
			// 	log.Println("received all jobs")
			// 	done <- true
			// 	return
			// }
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
