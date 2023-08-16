// Exercise 5.2: Write a function to populatea mapping for element names--p, div, span and so on--to the number of elements with that name in an HTML document tree
package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
    doc, err := html.Parse(os.Stdin)
    if err != nil {
        log.Fatal(err)
    }
    var m = make(map[string]int)
    m = tagcnt(m, doc)
    for k, v := range m {
        fmt.Printf("%9s: %v\n", k, v)
    }
}

func tagcnt(m map[string]int, n *html.Node) map[string]int {
    if n.Type == html.ElementNode {
        m[n.Data] += 1
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        tagcnt(m, c)
    }
    return m
}
