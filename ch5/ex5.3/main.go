// Exercise 5.3: Write a function to print the contents of all text nodes in an HTML document tree. Do not descend into <script> or <style> elements, since their contents are not visible in a web browser
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
	contents(doc)
}

func contents(n *html.Node) {
	if n.Type == html.ElementNode &&
		(n.Data == "script" || n.Data == "style") {
		return
	}
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}
	if n.FirstChild != nil {
		contents(n.FirstChild)
	}
	if n.NextSibling != nil {
		contents(n.NextSibling)
	}
}
