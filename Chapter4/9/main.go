package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("sample.txt")
	if err != nil {
		fmt.Print(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	m := make(map[string]int)

	for scanner.Scan() {
		m[scanner.Text()]++
	}

	for s, i := range m {
		fmt.Printf("[%d]-[%s]\n", i, s)
	}
}
