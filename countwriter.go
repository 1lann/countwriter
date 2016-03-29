// Package countwriter is a package for a writer that counts the number of
// bytes written to it.
package countwriter

import (
	"io"
	"sync/atomic"
)

// A CountWriter is a writer that passes data through to the underlying writer
// and counts the number of bytes written to it.
type CountWriter struct {
	w     io.Writer
	count uint64
}

func NewWriter(w io.Writer) *CountWriter {
	return &CountWriter{w: w, count: 0}
}

func (c *CountWriter) Count() uint64 {
	return atomic.LoadUint64(&c.count)
}

func (c *CountWriter) Write(p []byte) (int, error) {
	num, err := c.w.Write(p)
	atomic.AddUint64(&c.count, uint64(num))
	return num, err
}
