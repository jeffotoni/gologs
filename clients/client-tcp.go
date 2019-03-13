// Go in action
// @jeffotoni
// 2019-03-11

// client-tcp.go
package main

import (
  "bufio"
  "flag"
  "fmt"
  "log"
  "net"
  "strings"
)

func main() {

  host := flag.String("host", "", "")
  port := flag.String("port", "22335", "")

  flag.Parse()
  TCPHOST := *host + ":" + *port

  for i := 0; i < 10000; i++ {

    // connect to this socket
    conn, err := net.Dial("tcp", TCPHOST)
    if err != nil {
      log.Fatal("net Dial Client:", err)
    }

    // println("Text to send: ")
    jsonmsg := `{"versão": "1.1", "host": "exemplo.org", "short_message": "one msg here...", "nível": 5, "some_info": "foo"}`
    // send to socket
    fmt.Fprintf(conn, jsonmsg)

    // listen for reply
    message, _ := bufio.NewReader(conn).ReadString('\n')
    message = strings.Trim(message, " ")

    if strings.ToLower(message) == "ok" {
      // println("\nSave")
    } else {
      // println("\nError server tcp")
    }
  }
}
