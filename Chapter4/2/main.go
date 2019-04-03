package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var f = flag.Int("sha", 256, "select 256 or 384 or 512")

func main() {
	flag.Parse()
	fmt.Print("Please input value\n")
	r := bufio.NewReader(os.Stdin)
	b, err := r.ReadBytes('\n')
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	data := b[0 : len(b)-1]

	switch *f {
	case 256:
		fmt.Printf("sha256 of %x is [%x]\n", data, sha256.Sum256(data))
	case 384:
		fmt.Printf("sha384 of %x is [%x]\n", data, sha512.Sum384(data))
	case 512:
		fmt.Printf("sha512 of %x is [%x]\n", data, sha512.Sum512(data))
	}
}
