package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	word, image, err := CountWordsAndImages("https://github.com/")
	if err != nil {
		fmt.Printf("CountWordsAndimages error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Word count:[%d], Image count:[%d]\n", word, image)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	node := make([]*html.Node, 0)
	node = append(node, n)
	for len(node) > 0 {
		n = node[len(node)-1]
		node = node[:len(node)-1]
		switch n.Type {
		case html.TextNode:
			words += countWord(n.Data)
		case html.ElementNode:
			if n.Data == "img" {
				images++
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			node = append(node, c)
		}
	}

	return
}

func countWord(s string) (n int) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		n++
	}
	return
}
