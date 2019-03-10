// Go in Action
// @jeffotoni
// 2019-03-10

// server.go
package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"

	"github.com/jeffotoni/gologs/repo"
)

var jobs = make(chan string)

var done = make(chan bool)

type Args struct {
	Json string
}

type Receive struct{}

func (t *Receive) Json(args *Args, reply *string) error {

	if len(args.Json) <= 0 {
		*reply = `{"status":"error", "msg":"json field is required"}`
		return nil
	}
	//*reply = `{"status":"ok", "msg":"Receive json"}`
	*reply = "ok"
	// log.Println("Server Receive: ", args.Json)
	// add msg
	go Publish(args.Json)
	return nil
}

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
				fmt.Println("received job, process service: ", j)
				repo.InsertLog(j)
				// here send Postgres or ElasticSearch or SQS or S3.
				// depending on the message sending email

			} else {
				// never
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()
}

func main() {
	re := new(Receive)
	server := rpc.NewServer()
	server.Register(re)
	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	listener, e := net.Listen("tcp", ":22334")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	// exec
	// service
	Consumer()

	for {
		if conn, err := listener.Accept(); err != nil {
			log.Fatal("accept error: " + err.Error())
		} else {
			log.Printf("New connection established rpc server\n")
			go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}
}
