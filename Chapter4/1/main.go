package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func popCount(b byte) int {
	count := 0
	for ; b != 0; count++ {
		b &= b - 1
	}
	return count
}

func diffCount(b1, b2 [32]byte) int {
	count := 0
	for i := 0; i < len(b1); i++ {
		s := b1[i] ^ b2[i]
		count += popCount(s)
	}
	return count
}

func main() {
	s1 := sha256.Sum256([]byte("x"))
	s2 := sha256.Sum256([]byte("X"))
	fmt.Printf("diffCount is %d\n", diffCount(s1, s2))
}
