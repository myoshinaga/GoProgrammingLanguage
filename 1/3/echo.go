package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	echo()
	echoOptimized()
}

func echo() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("echo:  %.2fs\n", time.Since(start).Seconds())
}

func echoOptimized() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[0:], " "))
	fmt.Printf("Optimized: %.2fs\n", time.Since(start).Seconds())
}
