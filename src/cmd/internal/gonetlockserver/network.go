package gonetlockserver

import (
  "log"
  "net"
)

func Listen(host string, port string) {
  listener, err := net.Listen("tcp", host + ":" + port)
  if err != nil {
    log.Fatal("Error listening:", err.Error())
  }

  defer func() {
    // close listener
    _ = listener.Close()

    // close connections
    getConnectionManager().closeAllConnections()
  }()

  log.Println("Listening on " + host + ":" + port)

  for {
    // listen for an incoming connection
    connection, err := listener.Accept()
    if err != nil {
      log.Println("Error accepting: ", err.Error())
    } else {
      // handle connections in a new goroutine.
      go handleRequests(&connection)
    }
  }
}
