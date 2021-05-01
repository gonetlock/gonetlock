package gonetlockserver

import (
  "fmt"
  "sync"
)

type clientManager struct {
  clients map[string]*client
  lock        sync.Mutex
}

var containerClientManagerLock = &sync.Mutex{}
var containerClientManager *clientManager

func getClientManager() *clientManager {
  if containerClientManager == nil {
    // only lock on initialization
    containerClientManagerLock.Lock()

    // check again after lock
    if containerClientManager == nil {
      containerClientManager = new(clientManager)
      containerClientManager.clients = make(map[string]*client)
    }

    containerClientManagerLock.Unlock()
  }

  return containerClientManager
}

func (clientManager *clientManager) newClient(id string) *client {
  clientManager.lock.Lock()

  // Check if this client already exists. Return it if it does.
  existingClient := clientManager.clients[id]
  if existingClient != nil {
    return existingClient
  }

  // Create client object
  client := newClient(id)

  // Add client to map
  clientManager.clients[client.Id] = client

  fmt.Println("New client: ", client.Id)

  clientManager.lock.Unlock()

  return client
}

func (clientManager *clientManager) removeClient(client *client) {
  clientManager.lock.Lock()

  fmt.Println("Removing client: ", client.Id)
  clientManager.removeClientWithConnections(client)
  fmt.Println("client removed")

  clientManager.lock.Unlock()
}

func (clientManager *clientManager) removeClientWithConnections(client *client) {
  // remove the client from all connections
  for _, connection := range client.connections {
    connection.removeClient(client)
  }

  // remove the client from our clients list
  delete(clientManager.clients, client.Id)
}
