// Go in action
// @jeffotoni
// 2019-03-11

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jeffotoni/gologs/pkg/server"
)

func main() {

	done := make(chan struct{})

	println("Server Run...")

	// Rpc open
	go server.Rpc()

	// Tcp open
	go server.Tcp()

	// Receives job
	// from queue
	// and executes
	go server.Consumer()

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		log.Printf("caught sig: %+v", sig)
		log.Println("Wait for 1 second to finish processing...")
		time.Sleep(1 * time.Second)
		close(done)
		os.Exit(0)
	}()

	<-done
}
