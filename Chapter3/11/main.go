package main

import (
	"fmt"
	"strings"
)

func main() {
	c := commaFloat("123456789.1234")
	fmt.Println(c)
}

func commaFloat(s string) string {
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		return comma(s[:dot]) + s[dot:]
	}
	return comma(s)
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
