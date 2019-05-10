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

	//"github.com/getsentry/raven-go"
	"github.com/jeffotoni/gologs/config"
	"github.com/jeffotoni/gologs/pkg/nats"
	"github.com/jeffotoni/gologs/pkg/server"
)

// func init() {
// 	if len(os.Getenv("SET_RAVEN")) > 0 {
// 		raven.SetDSN(os.Getenv("SET_RAVEN"))
// 	}
// }

func main() {

	done := make(chan struct{})

	// only for nats
	if config.SERVICE == "nats" {
		go nats.SubscribeAsync(config.NATS_PERSISTENT)
		//go nats.SubscribeAsync("postgres")
	}

	// Rpc open
	go server.Rpc()

	// Tcp open
	go server.Tcp()

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		log.Printf("caught sig: %+v", sig)
		println("\033[0;31mWait for 1 second to finish processing...\033[0;0m")
		time.Sleep(1 * time.Second)
		close(done)
		os.Exit(0)
	}()

	<-done
}
