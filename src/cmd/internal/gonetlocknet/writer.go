package gonetlocknet

import (
  "encoding/binary"
  "io"
  "sync"
)

type Writer struct {
  Writer    io.Writer
  byteOrder binary.ByteOrder
  buffer    []byte
  mutex     sync.Mutex
}

// NewUint32DelimitedWriter creates a Writer.
func NewUint32DelimitedWriter(writer io.Writer) Writer {
  return Writer{writer, binary.BigEndian, make([]byte, 4), sync.Mutex{}}
}

func (writer *Writer) WriteMessage(message []byte) error {
  writer.mutex.Lock()

  defer func() {
    writer.mutex.Unlock()
  }()

  messageLength := len(message)

  size := messageLength + 4
  if size > len(writer.buffer) {
    writer.buffer = make([]byte, size)
  }

  // put message length into the buffer
  writer.byteOrder.PutUint32(writer.buffer, uint32(messageLength))

  // copy the message into the buffer
  copy(writer.buffer[4:], message)

  _, err := writer.Writer.Write(writer.buffer[:size])
  return err
}

func (writer *Writer) Close() error {
  if closer, ok := writer.Writer.(io.Closer); ok {
    return closer.Close()
  }

  return nil
}
