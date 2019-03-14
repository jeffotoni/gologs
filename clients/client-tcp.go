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
  "strconv"
  "strings"
  "time"
)

func main() {

  host := flag.String("host", "localhost", "")
  port := flag.String("port", "22335", "")
  request := flag.String("req", "10000", "")

  flag.Parse()
  TCPHOST := *host + ":" + *port

  req, _ := strconv.Atoi(*request)
  if req <= 0 {
    log.Println("Requests must be greater than 0")
    return
  }

  start := time.Now()
  fmt.Println("\033[0;32mRun Tests...\033[0;0m")
  fmt.Println("\033[0;33mRequests: ", req)
  fmt.Println("Port:     ", *port)
  fmt.Printf("\033[0;0m")

  for i := 0; i < req; i++ {

    // connect to this socket
    conn, err := net.Dial("tcp", TCPHOST)
    if err != nil {
      log.Println("Net Dial Client: [ ", i, " ] ", err)
      time.Sleep(time.Second * 10)
      continue
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

  end := time.Now()
  diff := end.Sub(start)
  fmt.Println("Time:    ", diff)
}
