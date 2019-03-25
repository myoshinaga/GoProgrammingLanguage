package main

import (
	"fmt"
	"sort"
)

type Rune []rune

func main() {
	b := isAnagram("あなぐらむ", "ぐあむなら")
	fmt.Printf("isAnagram=%t\n", b)
}

func isAnagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	r1 := Rune(s1)
	sort.Sort(r1)
	r2 := Rune(s2)
	sort.Sort(r2)

	for i, r := range r1 {
		if r != r2[i] {
			return false
		}
	}

	return true
}

func (r Rune) Len() int {
	return len(r)
}

func (r Rune) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r Rune) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
