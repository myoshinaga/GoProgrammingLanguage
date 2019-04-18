package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func findByID(n *html.Node, id string) bool {
	if n.Type != html.ElementNode {
		return false
	}
	for _, a := range n.Attr {
		if a.Key == "id" && a.Val == id {
			return true
		}

	}
	return false
}

func ElementByID(n *html.Node, id string) *html.Node {
	return forEachNode(n, id, findByID, nil)
}

func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) *html.Node {
	nodes := make([]*html.Node, 0)
	nodes = append(nodes, n)
	for len(nodes) > 0 {
		n = nodes[0]
		nodes = nodes[1:]
		if pre != nil {
			if pre(n, id) {
				return n
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			nodes = append(nodes, c)
		}
		if post != nil {
			if post(n, id) {
				return n
			}
		}
	}
	return nil
}

func main() {
	resp, err := http.Get("https://golang.org")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", ElementByID(doc, "gopher"))
}
