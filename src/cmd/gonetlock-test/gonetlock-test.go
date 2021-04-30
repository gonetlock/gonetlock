package main

import (
  "fmt"
  "github.com/gonetlock/gonetlock/cmd/internal/gonetlocknet"
  "github.com/gonetlock/gonetlock/gonetlockprotobuf"
  protobuf "github.com/gonetlock/gonetlock/gonetlockprotobuf/protobuf"
  "os"
)

const (
  ConnectionHost = "127.0.0.1"
  ConnectionPort = "1234"
  ConnectionType = "tcp"
)

func main() {
  connection, err := gonetlocknet.Dial(ConnectionType, ConnectionHost + ":" + ConnectionPort)
  if err != nil {
    fmt.Println("Error connecting:", err.Error())
    os.Exit(1)
  }

  defer func() {
    _ = connection.Close()
    fmt.Println("Closed connection")
  }()

  clientAttachRequest := protobuf.ClientAttachRequest {
    ClientId: "player-1",
  }

  messageBuffer, err := gonetlockprotobuf.PackMessage("123", &clientAttachRequest)
  if err != nil {
    return
  }

  _ = connection.WriteMessage(*messageBuffer)
}
