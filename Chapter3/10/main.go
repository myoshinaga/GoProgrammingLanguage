package main

import (
	"bytes"
	"fmt"
)

func main() {
	c := comma("123456789")
	fmt.Println(c)
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	remainder := n % 3
	var cIndex int
	if remainder == 0 {
		cIndex = 2
	} else {
		cIndex = remainder - 1
	}
	lastIndex := n - 1
	var buf bytes.Buffer
	for i, v := range s {
		buf.WriteRune(v)
		if i == cIndex {
			buf.WriteByte(',')
			if cIndex+3 != lastIndex {
				cIndex += 3
			}
		}
	}

	return buf.String()
}
