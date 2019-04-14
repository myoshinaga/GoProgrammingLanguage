package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	printTextNode(os.Stdin, os.Stdout)
}

func printTextNode(r io.Reader, w io.Writer) {
	token := html.NewTokenizer(os.Stdin)
	tags := make([]string, 40)
	t := token.Next()
	for ; t != html.ErrorToken; t = token.Next() {
		switch t {
		case html.StartTagToken:
			tagName, _ := token.TagName()
			tags = append(tags, string(tagName))
		case html.TextToken:
			tag := tags[len(tags)-1]
			if tag == "script" || tag == "style" {
				continue
			}
			text := token.Text()
			if len(strings.TrimSpace(string(text))) == 0 {
				continue
			}
			w.Write([]byte(fmt.Sprintf("<%s>", tag)))
			w.Write(text)
			if text[len(text)-1] != '\n' {
				io.WriteString(w, "\n")
			}
		case html.EndTagToken:
			tags = tags[:len(tags)-1]
		}
	}
}
