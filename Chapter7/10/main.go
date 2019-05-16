package main

import (
	"fmt"
	"sort"
)

func equal(i, j int, s sort.Interface) bool {
	return !s.Less(i, j) && !s.Less(j, i)
}

func IsPalindrome(s sort.Interface) bool {
	lastIndex := s.Len() - 1
	for i := 0; i < s.Len()>>1; i++ {
		if !equal(i, lastIndex-i, s) {
			return false
		}
	}
	return true
}

func main() {
	values1 := []int{1, 2, 3, 4, 3, 2, 1}
	fmt.Printf("%v: %v\n", values1, IsPalindrome(sort.IntSlice(values1)))

	values2 := []string{"a", "c", "b", "a"}
	fmt.Printf("%v: %v\n", values2, IsPalindrome(sort.StringSlice(values2)))
}
