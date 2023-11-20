// Exercise 8.6: Add depth-limiting to the concurrent crawler. That is, if the user sets -depth=3, then only URLs reachable by at most three links will be fetched.
package main

import (
	"flag"
	"fmt"
	"log"
	
	"the-go-programming-language/ch5/examples/links"
)

var depth = flag.Int("depth", 1, "URLs crawl depth")

// url with depth counter
type curl struct {
	c int
	url string
}

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(curl curl) []string {
	fmt.Println(curl.c, curl.url)
	if curl.c >= *depth {
		return nil
	}
	tokens <- struct{}{} // aquire token
	list, err := links.Extract(curl.url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	flag.Parse()
	worklist := make(chan []curl) // list of URLs, may have duplicates
	unseenLinks := make(chan curl) // de-duplicated URLs

	// Add command-line arguments to worklist
	go func() {
		var curls []curl
		for _, url := range flag.Args() {
			curls = append(curls, curl{0, url})
		}
		worklist <- curls
	}()

	// Create 20 crawlers goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				var foundCurls []curl
				for _, url := range foundLinks {
					foundCurls = append(foundCurls, curl{link.c+1, url})
				}
				go func() {
					worklist <- foundCurls
				}()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link.url] {
				seen[link.url] = true
				unseenLinks <- link
			}
		}
	}
}
