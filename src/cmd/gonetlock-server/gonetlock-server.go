package main

import (
  "fmt"
  "github.com/gonetlock/gonetlock/cmd/internal/gonetlockserver"
  "github.com/gonetlock/gonetlock/gonetlockprotobuf"
  "net"
  "os"
  "time"
)

const (
  ConnectionHost = "127.0.0.1"
  ConnectionPort = "1234"
  ConnectionType = "tcp"
)

var connectionManager = gonetlockserver.NewConnectionManager()

func main() {
  listener, err := net.Listen(ConnectionType, ConnectionHost+":"+ConnectionPort)
  if err != nil {
    fmt.Println("Error listening:", err.Error())
    os.Exit(1)
  }

  defer func() {
    // close listener
    _ = listener.Close()

    // close connections
    connectionManager.CloseAllConnections()
  }()

  fmt.Println("Listening on " + ConnectionHost + ":" + ConnectionPort)

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

func handleRequest(netConnection *net.Conn) {
  connection := connectionManager.NewConnection(netConnection)

  defer func() {
    connectionManager.CloseConnection(connection)
    fmt.Println("Deleted connection: ", connection.Id)
  }()

  for {
    // listen for commands
    messageBinary, err := connection.Connection.ReadMessage()
    if err != nil {
      if err.Error() == "EOF" {
        fmt.Println("Connection error: ", connection.Id, " ", err.Error())
        return
      } else {
        fmt.Println("Error reading line: ", err.Error())
      }
    } else {
      key, messageType, any, err := gonetlockprotobuf.UnpackMessage(&messageBinary)
      if err != nil {
        fmt.Println("Unpacking err: ", err.Error())
      }

      message := gonetlockserver.NewMessage(connection, key, messageType, any)

      go handleMessage(message)
    }
  }
}

func handleMessage(message *gonetlockserver.Request) {
  fmt.Println("Connection ", message.Connection.Id, " - Key: ", message.Key, " - MessageType: ", message.MessageType, " Message: ", *message.Message)

  time.Sleep(2 * time.Second)

  _ = message.Connection.Connection.WriteMessage(nil)
}
