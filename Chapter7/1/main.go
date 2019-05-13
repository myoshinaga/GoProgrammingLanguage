package main

import (
	"bufio"
	"fmt"
)

type WordCounter int
type LineCounter int

func main() {
	var w WordCounter
	w.Write([]byte("programming language go\n"))
	fmt.Println(w)

	var l LineCounter
	l.Write([]byte("abc\ndef\nghi\naaa bbb\n"))
	fmt.Println(l)
}

func (c *WordCounter) Write(p []byte) (int, error) {
	n := len(p)
	for {
		advance, token, _ := bufio.ScanWords(p, true)
		if token != nil {
			*c++
		}
		p = p[advance:]
		if (len(p)) == 0 {
			return n, nil
		}
	}
}

func (c *LineCounter) Write(p []byte) (int, error) {
	n := len(p)
	for {
		advance, token, _ := bufio.ScanLines(p, true)
		if token != nil {
			*c++
		}
		p = p[advance:]
		if (len(p)) == 0 {
			return n, nil
		}
	}
}
