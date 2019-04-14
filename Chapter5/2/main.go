package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for element, count := range countElement(make(map[string]int), doc) {
		fmt.Printf("%s: %d\n", element, count)
	}
}

func countElement(m map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		m[n.Data]++
	}
	if n.FirstChild != nil {
		m = countElement(m, n.FirstChild)
	}
	if n.NextSibling != nil {
		m = countElement(m, n.NextSibling)
	}

	return m
}
