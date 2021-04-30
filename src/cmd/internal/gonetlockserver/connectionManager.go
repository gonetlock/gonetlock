package gonetlockserver

import (
  "fmt"
  "net"
  "sync"
)

type connectionManager struct {
  connections map[string]*Connection
  lock        sync.Mutex
}

func NewConnectionManager() *connectionManager {
  connectionManger := new(connectionManager)
  connectionManger.connections = make(map[string]*Connection)

  return connectionManger
}

func (connectionManager *connectionManager) NewConnection(netConnection *net.Conn) *Connection {
  // Create Connection object
  connection := newConnection(netConnection)

  connectionManager.lock.Lock()
  // Check if this Connection already exists. Close and delete it if it does.
  existingConnection := connectionManager.connections[connection.Id]
  if existingConnection != nil {
    connectionManager.closeAndRemoveConnection(existingConnection)
  }

  // Add Connection to map
  connectionManager.connections[connection.Id] = connection

  fmt.Println("New Connection: ", connection.Id)

  connectionManager.lock.Unlock()

  return connection
}

func (connectionManager *connectionManager) CloseAllConnections() {
  connectionManager.lock.Lock()

  fmt.Println("Closing all connections")

  for _, connection := range connectionManager.connections {
    connectionManager.closeAndRemoveConnection(connection)
  }

  connectionManager.lock.Unlock()
}

func (connectionManager *connectionManager) CloseConnection(connection *Connection) {
  connectionManager.lock.Lock()

  fmt.Println("Closing Connection: ", connection.Id)
  connectionManager.closeAndRemoveConnection(connection)
  fmt.Println("Connection closed")

  connectionManager.lock.Unlock()
}

func (connectionManager *connectionManager) closeAndRemoveConnection(connection *Connection) {
  // close the tcp Connection
  connection.Close()

  // remove the Connection from all clients
  for _, client := range connection.clients {
    client.RemoveConnection(connection)
  }

  // remove the Connection from our connections list
  delete(connectionManager.connections, connection.Id)
}
