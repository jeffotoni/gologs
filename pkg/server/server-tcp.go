// Go in action
// @jeffotoni
// 2019-03-11

package server

import (
  "log"
  "net"
  "os"
  "strings"
)

//  8001 and 62554
const (
  CONN_HOST = ""
  CONN_PORT = "22335"
  CONN_TYPE = "tcp"
)

func Tcp() {
  // Listen for incoming connections.
  l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
  if err != nil {
    log.Println("Error listening:", err.Error())
    os.Exit(0)
  }
  // Close the listener when the application closes.
  defer l.Close()

  // loop
  for {
    // Listen for an incoming connection.
    conn, err := l.Accept()
    if err != nil {
      log.Println("Error accepting: ", err.Error())
      os.Exit(0)
    }
    // Handle connections in a new goroutine.
    go handleRequest(conn)
  }
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
  // Make a buffer to hold incoming data.
  buf := make([]byte, 1024)
  // Read the incoming connection into the buffer.
  _, err := conn.Read(buf)
  if err != nil {
    log.Println("Error reading:", err.Error())
    return
  }

  bufclean := string(buf)
  bufclean = strings.Trim(bufclean, "\u0000")

  //log.Println("reqLen: ", reqLen)
  //log.Println("msg: ", fub)

  // Goroutine Queue
  go Publish(bufclean)

  // Send a response back to person contacting us.
  conn.Write([]byte("ok"))
  // Close the connection when you're done with it.
  conn.Close()
}
