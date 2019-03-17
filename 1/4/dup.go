package main

import (
	"bufio"
	"fmt"
	"os"
)

// 「<重複回数> <重複テキスト> <該当ファイル>」の形式で表示
func main() {
	counts := make(map[string]int)
	fileNames := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileNames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileNames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, fileNames[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileNames map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		counts[text]++
		if !contains(fileNames[text], f.Name()) {
			fileNames[text] = append(fileNames[text], f.Name())
		}
	}
}

func contains(array []string, str string) bool {
	for _, a := range array {
		if str == a {
			return true
		}
	}
	return false
}
