package gonetlockserver

import (
  "fmt"
  "net"
  "sync"
)

type connectionManager struct {
  connections map[string]*connection
  lock        sync.Mutex
}

var containerConnectionManagerLock = &sync.Mutex{}
var containerConnectionManager *connectionManager

func getConnectionManager() *connectionManager {
  if containerConnectionManager == nil {
    // only lock on initialization
    containerConnectionManagerLock.Lock()

    // check again after lock
    if containerConnectionManager == nil {
      containerConnectionManager = new(connectionManager)
      containerConnectionManager.connections = make(map[string]*connection)
    }

    containerConnectionManagerLock.Unlock()
  }

  return containerConnectionManager
}

func (connectionManager *connectionManager) newConnection(netConnection *net.Conn) *connection {
  // Create connection object
  connection := newConnection(netConnection)

  connectionManager.lock.Lock()
  // Check if this connection already exists. close and delete it if it does.
  existingConnection := connectionManager.connections[connection.Id]
  if existingConnection != nil {
    connectionManager.closeAndRemoveConnection(existingConnection)
  }

  // Add connection to map
  connectionManager.connections[connection.Id] = connection

  // Create default client and add it to the connection
  clientManager := getClientManager()
  client := clientManager.newClient(connection.Id)
  connection.setDefaultClient(client)
  client.addConnection(connection)

  fmt.Println("New connection: ", connection.Id)

  connectionManager.lock.Unlock()

  return connection
}

func (connectionManager *connectionManager) closeAllConnections() {
  connectionManager.lock.Lock()

  fmt.Println("Closing all connections")

  for _, connection := range connectionManager.connections {
    connectionManager.closeAndRemoveConnection(connection)
  }

  connectionManager.lock.Unlock()
}

func (connectionManager *connectionManager) closeConnection(connection *connection) {
  connectionManager.lock.Lock()

  fmt.Println("Closing connection: ", connection.Id)
  connectionManager.closeAndRemoveConnection(connection)
  fmt.Println("connection closed")

  connectionManager.lock.Unlock()
}

func (connectionManager *connectionManager) closeAndRemoveConnection(connection *connection) {
  // close the tcp connection
  connection.close()

  // remove the connection from all clients
  for _, client := range connection.clients {
    client.removeConnection(connection)
  }

  // remove the connection from our connections list
  delete(connectionManager.connections, connection.Id)
}
