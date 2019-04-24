package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get("https://golang.org")
	if err != nil {
		fmt.Printf("ElementsByTagName: %v\n", err)
		os.Exit(1)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Printf("ElementsByTagName: %v\n", err)
		os.Exit(1)
	}

	tagname := []string{"body", "head"}
	ret := ElementsByTagName(doc, tagname...)
	for _, r := range ret {
		fmt.Printf("%s\n", r.Data)
	}
}

func forEachNode(n *html.Node, pre func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre)
	}
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var nodes []*html.Node
	elementsByTagName := func(doc *html.Node) {
		if doc.Type == html.ElementNode {
			for _, n := range name {
				if doc.Data == n {
					nodes = append(nodes, doc)
				}
			}
		}
	}
	forEachNode(doc, elementsByTagName)

	return nodes
}
