// Exercise 5.17: Write a variadic function ElementsByTagName that, given an HTML node tree and zero or more names, returns all the elements that match one of those names. Here are two examples calls
// func ElementsByTagName(doc *html.Node, name ...string) []*html.Node
// images := ElementsByTagName(doc, "img")
// headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}
	images := ElementsByTagName(doc, "img")
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	fmt.Println(len(images))
	fmt.Println(len(headings))
	
	return nil
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
	nodes := make([]*html.Node, 0)
	pre := func(doc *html.Node) {
		for _, n := range name {
			if doc.Type == html.ElementNode && doc.Data == n && doc.FirstChild != nil {
				{		nodes = append(nodes, doc.FirstChild)
				}
			}
		}
	}
	forEachNode(doc, pre)
	return nodes
}
