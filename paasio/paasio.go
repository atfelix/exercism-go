package paasio

import (
	"io"
	"sync"
)

type syncCounter struct {
	sync.Mutex
	count int64
	numOps int
}

func (counter *syncCounter) countBytes(numBytes int) {
	counter.Lock()
	defer counter.Unlock()
	counter.count += int64(numBytes)
	counter.numOps++
}

func (counter *syncCounter) totalCount() (int64, int) {
	counter.Lock()
	defer counter.Unlock()
	return counter.count, counter.numOps
}

type writeCounter struct {
	syncCounter
	io.Writer
}

func (wCounter *writeCounter) WriteCount() (int64, int) {
	return wCounter.totalCount()
}

func (wCounter *writeCounter) Write(p []byte) (int, error) {
	numWrites, err := wCounter.Writer.Write(p)
	wCounter.countBytes(numWrites)
	return numWrites, err
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{ Writer: writer }
}

type readCounter struct {
	syncCounter
	io.Reader
}

func (rCounter *readCounter) ReadCount() (int64, int) {
	return rCounter.totalCount()
}

func (rCounter *readCounter) Read(p []byte) (int, error) {
	numReads, err := rCounter.Reader.Read(p)
	rCounter.countBytes(numReads)
	return numReads, err
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{ Reader: reader }
}

type readWriteCounter struct {
	syncCounter
	writeCounter
	readCounter
}

func NewReadWriteCounter(readWriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		writeCounter: writeCounter{ Writer: readWriter },
		readCounter: readCounter{ Reader: readWriter },
	}
}