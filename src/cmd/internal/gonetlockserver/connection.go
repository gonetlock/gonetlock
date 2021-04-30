package gonetlockserver

import (
  "github.com/gonetlock/gonetlock/cmd/internal/gonetlocknet"
  "net"
)

type Connection struct {
  Connection    *gonetlocknet.Conn
  DefaultClient *client
  clients       map[string]*client
  Id            string
}

// to be used by connectionManager only
func newConnection(netConnection *net.Conn) *Connection {
  connection := new(Connection)

  connection.Connection = gonetlocknet.NewConn(*netConnection)

  connection.clients = make(map[string]*client)
  connection.Id = (*netConnection).RemoteAddr().String() + "-" + (*netConnection).LocalAddr().String()

  return connection
}

// to be used by connectionManager only
func (connection *Connection) Close() {
  // close the Connection
  _ = connection.Connection.Close()
}

// to be used by connectionManager only
func (connection *Connection) AddClient(client *client) {
  connection.clients[client.Id] = client
}

// to be used by connectionManager only
func (connection *Connection) RemoveClient(client *client) {
  delete(connection.clients, client.Id)
}
