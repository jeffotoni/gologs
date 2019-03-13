// Go in action
// @jeffotoni
// 2019-03-11

// client-tcp.go
package main

import (
  "bufio"
  "fmt"
  "net"
  "strings"
)

func main() {
  for i := 0; i < 10000; i++ {
    // connect to this socket
    conn, _ := net.Dial("tcp", "localhost:22335")
    // println("Text to send: ")
    jsonmsg := `{"versão": "1.1", "host": "exemplo.org", "short_message": "one msg", "nível": 5, "_some_info": "foo"}`
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
