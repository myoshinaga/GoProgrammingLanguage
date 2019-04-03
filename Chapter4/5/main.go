package main

import "fmt"

func removeDup(strings []string) []string {
	i := 0
	for _, s := range strings {
		if strings[i] == s {
			continue
		}
		i++
		strings[i] = s
	}
	return strings[:i+1]
}

func main() {
	s := []string{"aaa", "aaa", "bbb", "bbb", "ccc", "ddd", "ddd", "ddd"}

	fmt.Printf("removeDup %v\n", removeDup(s))
}
