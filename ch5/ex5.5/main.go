// Exercise 5.5: Implement countWordsAndImages.
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
    for _, url := range os.Args[1:] {
        words, images, err := CountWordsAndImages(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
            continue
        }
        fmt.Printf("url: %s\nwords: %v\timages: %v\n", url, words, images)
    }
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
    if n.Type == html.ElementNode && n.Data == "img" {
        images += 1;
    } else if n.Type == html.TextNode {
        words += wct(n.Data)
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        w, i := countWordsAndImages(c)
        words += w
        images += i
    }
    return
}

// count words from string input
func wct (in string) int {
    var w int
    input := bufio.NewScanner(strings.NewReader(in))
    input.Split(bufio.ScanWords) 
    for input.Scan() {
        w += 1
    }
    return w
}

