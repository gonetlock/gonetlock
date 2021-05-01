package gonetlockserver

type client struct {
  connections map[string]*connection
  Id          string
}

// to be used by clientManager only
func newClient(id string) *client {
  client := new(client)
  client.Id = id
  client.connections = make(map[string]*connection)

  return client
}

// to be used by clientManager only
func (client *client) addConnection(connection *connection) {
  if client.hasConnection(connection) {
    return
  }

  client.connections[connection.Id] = connection
}

// to be used by clientManager only
func (client *client) removeConnection(connection *connection) {
  if !client.hasConnection(connection) {
    return
  }

  delete(client.connections, connection.Id)
}

// to be used by clientManager only
func (client *client) hasConnection(connection *connection) bool {
  _, present := client.connections[connection.Id]

  return present
}
