// Go in Action
// @jeffotoni
// 2019-03-25

package main

import (
	"flag"
	"log"
	"strconv"

	nats "github.com/nats-io/go-nats"
)

func main() {

	p := flag.String("producer", "1", "example: -producer 1")
	flag.Parse()
	produc := *p

	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)
	log.Println("Producer Msg Json.. 500000")
	for i := 0; i < 500000; i++ {
		msgJson := `{"versão": "1.1","host": "exemplo.org","key":"producer_` + produc + `_` + strconv.Itoa(i) + `","level":"info","project":"my-project-here","short_message":"one msg here...","nível": 5,"some_info":"foo jeff"}`
		if err := nc.Publish("gologs", []byte(msgJson)); err != nil {
			log.Println(err)
		}
		// Close connection
		defer nc.Close()
		nc.Flush()
	}
}
