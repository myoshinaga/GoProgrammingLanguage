package main

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type reader struct {
	s string
}

func (r *reader) Read(p []byte) (n int, err error) {
	n = copy(p, r.s)
	fmt.Println(r.s)
	r.s = r.s[n:]
	if len(r.s) == 0 {
		err = io.EOF
	}
	return
}

func NewReader(s string) io.Reader {
	return &reader{s}
}

func main() {
	s := "<html><head></head><body><div></div></body></html>"
	_, err := html.Parse(NewReader(s))
	fmt.Printf("%v", err)
}
