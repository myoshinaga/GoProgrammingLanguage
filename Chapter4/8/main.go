package main

import (
	"fmt"
	"unicode"
)

func countChar(str string) {
	rs := []rune(str)
	m := make(map[rune]int)
	for _, r := range rs {
		if unicode.IsLetter(r) {
			m[r]++
		}
	}

	for r, i := range m {
		fmt.Printf("%#U[%d]\t", r, i)
	}
}

func main() {
	countChar("あaaあいう߷えݶおkakiく茉あああaaaaaaaaa")
}
