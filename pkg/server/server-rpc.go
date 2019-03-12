// Go in Action
// @jeffotoni
// 2019-03-10

// server.go
package server

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

var PORT_RPC = ":22334"

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

func Rpc() {
	re := new(Receive)
	server := rpc.NewServer()
	server.Register(re)
	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	listener, e := net.Listen("tcp", PORT_RPC)
	if e != nil {
		log.Println("listen error:", e)
		return
	}

	// to listen
	for {
		if conn, err := listener.Accept(); err != nil {
			log.Println("accept error: " + err.Error())
			return
		} else {
			log.Printf("New connection established in rpc server\n")
			go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}
}