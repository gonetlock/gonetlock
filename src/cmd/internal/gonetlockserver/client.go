package gonetlockserver

type client struct {
  connections map[string]*Connection
  Id          string
}

// to be used by connectionManager only
func (client *client) AddConnection(connection *Connection) {
  client.connections[connection.Id] = connection
}

// to be used by connectionManager only
func (client *client) RemoveConnection(connection *Connection) {
  delete(client.connections, connection.Id)
}
