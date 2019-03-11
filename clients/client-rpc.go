// Go in Action
// @jeffotoni
// 2019-03-10

// client-rpc.go
package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
	"strconv"
)

type Args struct {
	Json string
}

func main() {

	client, err := net.Dial("tcp", "127.0.0.1:22334")
	if err != nil {
		log.Fatal("dialing client:", err)
	}

	var reply string
	args := &Args{}
	c := jsonrpc.NewClient(client)

	for i := 0; i < 10000; i++ {
		// Synchronous call
		args = &Args{`{"key":"jeff_` + strconv.Itoa(i) + `","level":"info"}`}
		err = c.Call("Receive.Json", args, &reply)
		if err != nil {
			log.Fatal("capture json error:", err)
		}
		fmt.Printf("Result: %s\n", reply)
	}
}
