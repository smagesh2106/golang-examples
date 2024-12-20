package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, _ := html.Parse(os.Stdin)
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			links = visit(links, c)
		}
	}
	return links
}
