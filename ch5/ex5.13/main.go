// Exercise 5.13: Modify crawl to make local copies of the pages it finds, creating directories as necessary. Don't make copies of pages that come from a different domain. For example, if the original page comes from golang.org, save all files from there, but exclude ones from vimeo.com.
package main

import (
	"fmt"
	"log"
	"net/url"
	"net/http"
	"os"

	"gopl/ch5/examples/links"
)

const (
	URLSDIR = "urls"
)

func breadthFirst(f func(item string) []string, worklist []string) {
    seen := make(map[string]bool)
    for len(worklist) > 0 {
        items := worklist
        worklist = nil
        for _, item := range items {
            if !seen[item] {
                seen[item] = true
                worklist = append(worklist, f(item)...)
            }
        }
    }
}

func crawl(u string) []string {
	fmt.Println(u)
    list, err := links.Extract(u)
    if err != nil {
        log.Print(err)
    }

	os.Mkdir(URLSDIR, 0777) 
	for _, v := range list {
		//fmt.Println(url.PathEscape(v[8:]))
		u, err := url.Parse(v)
		if err != nil {
			log.Fatal(err)
		}
		hpath := URLSDIR + "/" + u.Host
		if _, err := os.Stat(hpath); os.IsNotExist(err) {
			os.Mkdir(hpath, 0777)
		}
		resp, err := http.Get(v)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		tmp := url.PathEscape(v) // html file name 
		f, err := os.Create(hpath + "/" + tmp + ".html")
		if err != nil {
			log.Fatal(err)
		}
		f.ReadFrom(resp.Body)
		fmt.Printf("Content of:\t %s.html writed\n", tmp)
	}
	
    return list
}

func main() {
    // Crawl the web breadth-first,
    // starting from the command-line arguments.
    breadthFirst(crawl, os.Args[1:])
}
