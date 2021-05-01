package gonetlockserver

import (
  "github.com/gonetlock/gonetlock/cmd/internal/gonetlocknet"
  "net"
)

type connection struct {
  Connection    *gonetlocknet.Conn
  DefaultClient *client
  clients       map[string]*client
  Id            string
}

// to be used by connectionManager only
func newConnection(netConnection *net.Conn) *connection {
  connection := new(connection)

  connection.Connection = gonetlocknet.NewConn(*netConnection)

  connection.clients = make(map[string]*client)
  connection.Id = (*netConnection).RemoteAddr().String() + "-" + (*netConnection).LocalAddr().String()

  return connection
}

// to be used by connectionManager only
func (connection *connection) close() {
  // close the connection
  _ = connection.Connection.Close()
}

// to be used by connectionManager only
func (connection *connection) addClient(client *client) {
  if connection.hasClient(client) {
    return
  }

  connection.clients[client.Id] = client
}

// to be used by connectionManager only
func (connection *connection) removeClient(client *client) {
  if !connection.hasClient(client) {
    return
  }

  delete(connection.clients, client.Id)
}

// to be used by connectionManager only
func (connection *connection) hasClient(client *client) bool {
  _, present := connection.clients[client.Id]

  return present
}

// to be used by connectionManager only
func (connection *connection) setDefaultClient(client *client) {
  connection.DefaultClient = client

  connection.addClient(client)
}
