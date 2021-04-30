package gonetlockserver

import googleproto "google.golang.org/protobuf/proto"

type Request struct {
  Connection *Connection
  Key string
  MessageType string
  Message *googleproto.Message
}

func NewMessage (connection *Connection, key string, messageType string, any *googleproto.Message) *Request {
  message := new(Request)
  message.Connection = connection
  message.Key = key
  message.MessageType = messageType
  message.Message = any

  return message
}
