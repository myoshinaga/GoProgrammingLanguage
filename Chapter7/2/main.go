package main

import (
	"bytes"
	"fmt"
	"io"
)

type NewWriter struct {
	w io.Writer
	n int64
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var nw NewWriter
	nw.w = w
	return &nw, &nw.n
}

func (c *NewWriter) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	c.n += int64(n)
	return n, err
}

func main() {
	buf := &bytes.Buffer{}
	c, n := CountingWriter(buf)
	c.Write([]byte("1234567890"))
	fmt.Printf("%d", *n)
}
