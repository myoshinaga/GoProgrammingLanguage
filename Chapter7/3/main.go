package main

import (
	"bytes"
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

func main() {
	values := []int{1, 3, 10, 14, 9, 4}
	Sort(values)
}

func (t *tree) String() string {
	var values []int
	values = appendValues(values, t)
	pBuf := &bytes.Buffer{}
	for _, v := range values {
		fmt.Fprintf(pBuf, "%d ", v)
	}
	return pBuf.String()
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	fmt.Println(root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
