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
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a":
			find, val := findAttrValue(n.Attr, "href")
			if find {
				links = append(links, val)
			}
		case "img", "script":
			find, val := findAttrValue(n.Attr, "src")
			if find {
				links = append(links, val)
			}
		case "link":
			find, _ := findAttrValue(n.Attr, "rel")
			if find {
				hrefFind, val := findAttrValue(n.Attr, "href")
				if hrefFind {
					links = append(links, val)
				}
			}
		}
	}
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}

	return links
}

func findAttrValue(attr []html.Attribute, key string) (bool, string) {
	for _, a := range attr {
		if a.Key == key {
			return true, a.Val
		}
	}
	return false, ""
}
