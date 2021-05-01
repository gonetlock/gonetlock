package main

import "github.com/gonetlock/gonetlock/cmd/internal/gonetlockserver"

const (
  ConnectionHost = "127.0.0.1"
  ConnectionPort = "1234"
)

func main() {
  gonetlockserver.Listen(ConnectionHost, ConnectionPort)
}
