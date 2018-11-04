package circular

import (
	"fmt"
	"errors"
)

type Buffer struct {
	capacity int
	nextRead, nextWrite int
	isFull bool
	buffer []byte
}

func NewBuffer(size int) *Buffer {
	buffer := make([]byte, size)
	return &Buffer{
		capacity: size,
		nextRead: 0,
		nextWrite: 0,
		isFull: false,
		buffer: buffer,
	}
}

func (buffer *Buffer) size() int {
	return (buffer.nextWrite - buffer.nextRead + buffer.capacity) % buffer.capacity
}

func (buffer *Buffer) isEmpty() bool {
	return buffer.size() == 0 && !buffer.isFull
}

func (buffer *Buffer) ReadByte() (byte, error) {
	if buffer.isEmpty() {
		return ' ', errors.New("Buffer is empty:  Could not read from it.")
	}
	b := buffer.buffer[buffer.nextRead]
	buffer.nextRead = (buffer.nextRead + 1) % buffer.capacity
	buffer.isFull = false
	return b, nil
}

func (buffer *Buffer) WriteByte(b byte) error {
	if buffer.isFull {
		return errors.New("Buffer is full:  Could not write to it.")
	}
	buffer.Overwrite(b)
	return nil
}

func (buffer *Buffer) Overwrite(b byte) {
	buffer.buffer[buffer.nextWrite] = b
	if buffer.isFull {
		buffer.nextRead = (buffer.nextRead + 1) % buffer.capacity
	}
	buffer.nextWrite = (buffer.nextWrite + 1) % buffer.capacity
	buffer.isFull = buffer.nextRead == buffer.nextWrite
}

func (buffer *Buffer) Reset() {
	buffer.nextRead = 0
	buffer.nextWrite = 0
	buffer.isFull = false
	buffer.buffer = make([]byte, buffer.capacity)
}

func (buffer *Buffer) String() string {
	return fmt.Sprintf("capacity:  %v\nnextRead:  %v\nnextWrite: %v\nbuffer:   %v\n", buffer.capacity, buffer.nextRead, buffer.nextWrite, buffer.buffer)
}