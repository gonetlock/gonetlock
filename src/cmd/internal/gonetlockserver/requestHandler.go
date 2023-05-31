package gonetlockserver

import (
  "fmt"
  "github.com/gonetlock/gonetlock/gonetlockprotobuf"
  "net"
  "time"
)

func handleRequests(netConnection *net.Conn) {
  connectionManager := getConnectionManager()
  connection := connectionManager.newConnection(netConnection)

  defer func() {
    connectionManager.closeConnection(connection)
    fmt.Println("Deleted connection: ", connection.Id)
  }()

  for {
    // listen for commands
    messageBinary, err := connection.Connection.ReadMessage()
    if err != nil {
      if err.Error() == "EOF" {
        fmt.Println("connection error: ", connection.Id, " ", err.Error())

        return
      } else {
        fmt.Println("Error reading line: ", err.Error())
      }
    } else {
      key, messageType, any, err := gonetlockprotobuf.UnpackMessage(&messageBinary)
      if err != nil {
        fmt.Println("Unpacking err: ", err.Error())
      }

      message := newMessage(connection, key, messageType, any)

      go handleMessage(message)
    }
  }
}

func handleMessage(message *request) {
  fmt.Println("connection ", message.Connection.Id, " - Key: ", message.Key, " - MessageType: ", message.MessageType, " Message: ", *message.Message)

  time.Sleep(2 * time.Second)

  _ = message.Connection.Connection.WriteMessage(nil)
}
