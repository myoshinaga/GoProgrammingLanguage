package main

import (
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[0:] {
		fmt.Println(arg)
	}
}