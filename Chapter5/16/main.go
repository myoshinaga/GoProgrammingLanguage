package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(join("\t", "test1", "test2", "test3"))
}

func join(sep string, a ...string) string {
	switch len(a) {
	case 0:
		return ""
	case 1:
		return a[0]
	}

	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(a[0])
	for _, s := range a[1:] {
		b.WriteString(sep)
		b.WriteString(s)
	}
	return b.String()
}
