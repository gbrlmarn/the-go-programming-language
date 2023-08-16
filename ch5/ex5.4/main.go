// Exercise 5.4: Extend the visit function so that it extracts other kinds of links from the document, such as images, scripts, and style sheets.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var stuff = make(map[string][]string)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for k, v := range visit(stuff, doc) {
        fmt.Printf("%s: %s\n", k, v)
	}
}

// visit appends to m each link, img, script and
// style sheet found in n and returns the result.
func visit(m map[string][]string, n *html.Node) map[string][]string {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a":
			for _, a := range n.Attr {
				if a.Key == "href" {
					m[n.Data] = append(m[n.Data], a.Val+"\n")
				}
			}
		case "img":
            for _, img := range n.Attr {
                if img.Key == "src" {
                    m[n.Data] = append(m[n.Data], img.Val+"\n")
                } 
            }
		case "script":
            for _, script := range n.Attr {
                if script.Key == "src" {
                    m[n.Data] = append(m[n.Data], script.Val+"\n")
                }
            }
		case "link":
            for _, link := range n.Attr {
                if link.Key == "href" {
                    m[n.Data] = append(m[n.Data], link.Val+"\n")
                }
            }
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		m = visit(m, c)
	}
	return m
}
