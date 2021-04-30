package gonetlocknet

import (
  "io"
  "net"
)

type Conn struct {
  Reader
  Writer
  conn io.ReadWriteCloser
}

// NewConn returns a new Conn using conn for I/O.
func NewConn(conn io.ReadWriteCloser) *Conn {
  return &Conn{
    Reader: NewUint32DelimitedReader(conn, 0),
    Writer: NewUint32DelimitedWriter(conn),
    conn:   conn,
  }
}

func (c *Conn) Close() error {
  return c.conn.Close()
}

// Dial connects to the given address on the given network using net.Dial
// and then returns a new Conn for the connection.
func Dial(network, addr string) (*Conn, error) {
  c, err := net.Dial(network, addr)
  if err != nil {
    return nil, err
  }
  return NewConn(c), nil
}
