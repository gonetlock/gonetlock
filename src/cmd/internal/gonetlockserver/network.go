package gonetlockserver

import (
  "fmt"
  "net"
  "os"
)

func Listen(host string, port string) {
  listener, err := net.Listen("tcp", host + ":" + port)
  if err != nil {
    fmt.Println("Error listening:", err.Error())
    os.Exit(1)
  }

  defer func() {
    // close listener
    _ = listener.Close()

    // close connections
    getConnectionManager().closeAllConnections()
  }()

  fmt.Println("Listening on " + host + ":" + port)

  for {
    // listen for an incoming connection
    connection, err := listener.Accept()
    if err != nil {
      fmt.Println("Error accepting: ", err.Error())
      os.Exit(1)
    }

    // handle connections in a new goroutine.
    go handleRequest(&connection)
  }
}
