// Exercise 8.7: Write a concurrent program that creates a local mirror of a web site, fetching each reachable page and writing it to a directory on the local disk. Only pages withing the original domain (for instance, golang.org) should be fetched. URLs within mirrored pages should be altered as needed so taht they refer to the mirrored page, not the original.
package main

import (
	"flag"
	"fmt"
	"gopl/ch5/examples/links"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var depth = flag.Int("depth", 1, "URLs crawl depth")

// url with depth counter
type curl struct {
	c   int
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
	save(curl.url)
	return list
}

func save(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	u := resp.Request.URL
	filename := filepath.Join(u.Host, u.Path)
	if filepath.Ext(u.Path) == "" {
		filename = filepath.Join(u.Host, u.Path, "index.html")
	}
	if err := os.MkdirAll(filepath.Dir(filename), 0777); err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	flag.Parse()
	worklist := make(chan []curl)  // list of URLs, may have duplicates
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
					foundCurls = append(foundCurls, curl{link.c + 1, url})
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
