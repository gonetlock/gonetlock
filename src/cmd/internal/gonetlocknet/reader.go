package gonetlocknet

import (
  "encoding/binary"
  "io"
  "sync"
)

type Reader struct {
  Reader    io.Reader
  byteOrder binary.ByteOrder
  lenBuf    []byte
  buffer    []byte
  maxSize   int
  mutex     sync.Mutex
}

// NewUint32DelimitedReader creates a Reader.
// If maxSize is negative, it will be treated as unlimited.
// If maxSize is 0, it will default to 1MiB.
func NewUint32DelimitedReader(reader io.Reader, maxSize int) Reader {
  if maxSize == 0 {
    // 1MiB deafault, use -1 for unlimited
    maxSize = 1 * 1024 * 1024
  }

  return Reader{reader, binary.BigEndian, make([]byte, 4), nil, maxSize, sync.Mutex{}}
}

func (reader *Reader) ReadMessage() ([]byte, error) {
  reader.mutex.Lock()

  defer func() {
    reader.mutex.Unlock()
  }()

  _, err := io.ReadFull(reader.Reader, reader.lenBuf)
  if err != nil {
    return make([]byte, 0), err
  }

  length32 := reader.byteOrder.Uint32(reader.lenBuf)
  length := int(length32)
  if length < 0 || (reader.maxSize > 0 && length > reader.maxSize) {
    return make([]byte, 0), io.ErrShortBuffer
  }

  if length > len(reader.buffer) {
    reader.buffer = make([]byte, length)
  }

  _, err = io.ReadFull(reader.Reader, reader.buffer[:length])
  if err != nil {
    return make([]byte, 0), err
  }

  message := make([]byte, length)
  copy(message, reader.buffer[:length])

  return message, nil
}

func (reader *Reader) Close() error {
  if closer, ok := reader.Reader.(io.Closer); ok {
    return closer.Close()
  }

  return nil
}
