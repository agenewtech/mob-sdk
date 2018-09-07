package comm

import (
	"bytes"
	"sync"
)

var bufferPool = &sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 1024))
	},
}

func PutBuffer(buf *bytes.Buffer) {
	if buf != nil {
		bufferPool.Put(buf)
	}
}

func GetBuffer() *bytes.Buffer {
	b := bufferPool.Get().(*bytes.Buffer)
	b.Reset()
	return b
}
