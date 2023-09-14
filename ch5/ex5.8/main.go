// Exercise 5.8: Modify forEachNode so that the pre and post functions return a boolean result indicating whether to continue the traversal. Use it to write a function ElementByID with the following signature that finds the first HTML element with the specified id attribute. The function should stop the traversal as soon as a match is found.
// func ElementByID(doc *html.Node, id string) *html.Node

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	id := flag.String("id", "start", "Please provide what id name you are looking...")
	flag.Parse()
	
	for _, url := range os.Args[1:] {
		outline(url, *id)
	}
}

func outline(url, id string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}
	node := ElementByID(doc, id)
	if node != nil {
		fmt.Println(node.Type, node.Data)
	}
	return nil
}

// We only need pre
func forEachNode(n *html.Node, pre func(n *html.Node) bool) (bool, *html.Node) {
	if pre != nil {
		if pre(n) {
			return true, n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		found, node := forEachNode(c, pre)
		if found {
			return found, node
		}
	}
	return false, nil
}

func ElementByID(doc *html.Node, id string) *html.Node {
	pre := func(doc *html.Node) bool {
		for _, a := range doc.Attr {
			if a.Key == "id" && a.Val == id {
				return true
			}
		}
		return false
	}
	
	found, res := forEachNode(doc, pre)
	if found {
		return res
	}
	return nil
}
