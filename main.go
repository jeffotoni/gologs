// Go in action
// @jeffotoni
// 2019-03-11

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	pg "github.com/jeffotoni/gologs/pkg/psql"
	"github.com/jeffotoni/gologs/pkg/server"
)

var db *sql.DB

func main() {

	DBINFO := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		pg.DB_HOST, pg.DB_PORT, pg.DB_USER, pg.DB_PASSWORD, pg.DB_NAME, pg.DB_SSL)

	db, err := sql.Open(pg.DB_SORCE, DBINFO)
	if err != nil {
		log.Println(err.Error())
		return
	}

	done := make(chan struct{})

	// Rpc open
	go server.Rpc(db)

	// Tcp open
	// go server.Tcp()

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
		println("\033[0;31mWait for 1 second to finish processing...\033[0;0m")
		time.Sleep(1 * time.Second)
		close(done)
		os.Exit(0)
	}()

	<-done
}
