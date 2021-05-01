package gonetlockserver

import googleproto "google.golang.org/protobuf/proto"

type request struct {
  Connection *connection
  Key string
  MessageType string
  Message *googleproto.Message
}

func newMessage(connection *connection, key string, messageType string, any *googleproto.Message) *request {
  message := new(request)
  message.Connection = connection
  message.Key = key
  message.MessageType = messageType
  message.Message = any

  return message
}
