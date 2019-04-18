package main

import (
	"bufio"
	"fmt"
	"strings"
)

func replace(s string) string {
	return strings.ToUpper(s)
}

func expand(s string, f func(string) string) (ret string) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		text := scanner.Text()
		if text[0] == '$' {
			ret += f(text[1:])
		} else {
			ret += text
		}
		ret += " "
	}

	return
}

func main() {
	fmt.Print(expand("$PROGRAMMING $LANGUAGE $GO is too difficult\n", replace))
}
