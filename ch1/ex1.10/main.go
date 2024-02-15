// Exercise 1.10: Find a web site that produces a large amount of data. Investigate caching by running fetchall twice in succession to see whether the reported time changes much. Do you get the same content each time? Modify fetchhall to print its output to a file so it can be examined
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	f, err := os.Create("ex1.10.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex1.10: %s %v", f.Name(), err)
	}
	for range os.Args[1:] {
		_, err := f.WriteString(<-ch)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex1.10: write %s %v", f.Name(), err)
		}
	}
	finish := fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())
	_, err = f.WriteString(finish)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex1.10: write %s %v", f.Name(), err)
	}
	f.Close()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
}
