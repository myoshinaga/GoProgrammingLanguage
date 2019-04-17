package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var depth int

func startElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		depth++
		for _, a := range n.Attr {
			fmt.Printf(" %s='%s'", a.Key, a.Val)
		}
		if n.FirstChild == nil {
			fmt.Printf(" />\n")
		} else {
			fmt.Printf(">\n")
		}

	case html.TextNode:
		fmt.Printf("%*s%s\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		depth--
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}

	case html.TextNode:
		depth--
	}
}

func main() {
	resp, err := http.Get("https://golang.org")
	if err != nil {
		os.Exit(1)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	forEachNode(doc, startElement, endElement)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
