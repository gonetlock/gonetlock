package gonetlockprotobuf

import (
  "github.com/golang/protobuf/proto"
  "github.com/golang/protobuf/ptypes"
  gonetlockprotobuf "github.com/gonetlock/gonetlock/gonetlockprotobuf/protobuf"
  googleproto "google.golang.org/protobuf/proto"
)

func PackMessage(key string, anyMessage proto.Message) (*[]byte, error) {
  any, err := ptypes.MarshalAny(anyMessage)
  if err != nil {
    return nil, err
  }

  message := gonetlockprotobuf.Message{
  MessageKey:  key,
  Message:     any,
  }

  messageBufer, err := proto.Marshal(&message)

  return &messageBufer, err
}

func UnpackMessage(buffer *[]byte) (string, string, *googleproto.Message, error) {
  message := gonetlockprotobuf.Message{}

  err := proto.Unmarshal(*buffer, &message)
  if err != nil {
    return "", "", nil, err
  }

  messageName := string(message.Message.MessageName())
  any, err := message.Message.UnmarshalNew()
  if err != nil {
    return message.MessageKey, messageName, nil, err
  }

  return message.MessageKey, messageName, &any, nil
}
