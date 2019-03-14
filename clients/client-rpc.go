// Go in Action
// @jeffotoni
// 2019-03-10

// client-rpc.go
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
	"strconv"
	"time"
)

type Args struct {
	Json string
}

func main() {

	host := flag.String("host", "127.0.0.1", "")
	port := flag.String("port", "22334", "")
	request := flag.String("req", "10000", "")

	flag.Parse()
	TCPHOST := *host + ":" + *port
	client, err := net.Dial("tcp", TCPHOST)
	if err != nil {
		log.Fatal("dialing client:", err)
		return
	}

	req, _ := strconv.Atoi(*request)
	if req <= 0 {
		log.Fatal("Requests must be greater than 0")
		return
	}

	start := time.Now()
	fmt.Println("\033[0;32mRun Tests...\033[0;0m")
	fmt.Println("\033[0;33mRequests: ", req)
	fmt.Println("Port:     ", *port)
	fmt.Printf("\033[0;0m")

	var reply string
	args := &Args{}
	c := jsonrpc.NewClient(client)

	for i := 0; i < req; i++ {

		if i == 10000 {
			time.Sleep(time.Second * 10)
		}

		// Synchronous call
		args = &Args{`{"key":"jeff_` + strconv.Itoa(i) + `","level":"info", "project":"my-project-here"}`}
		err = c.Call("Receive.Json", args, &reply)
		if err != nil {
			log.Fatal("capture json error:", err)
		}
		// fmt.Printf("Result: %s\n", reply)
	}

	end := time.Now()
	diff := end.Sub(start)
	fmt.Println("Time:    ", diff)
}
