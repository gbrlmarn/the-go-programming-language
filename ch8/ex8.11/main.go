// Exercise 8.11: Following the approach of mirroredQuery in Section 8.4.4, implement a variant of fetch that request several URLs concurrently. As soon as the first response arrives, cancel the other requests.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

type file struct {
	name string
	n    int64
	err  error
}

var cancel = make(chan struct{})

func main() {
	responses := make(chan file, len(os.Args[1:]))
	for _, url := range os.Args[1:] {
		go func() { responses <- fetch(url) }()
	}
	file := <-responses
	close(cancel)
	fmt.Printf("%s\t%d\t%v\n", file.name, file.n, file.err)
}

// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url string) file {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return file{"", 0, err}
	}
	req.Cancel = cancel
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return file{"", 0, err}
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return file{"", 0, err}
	}
	n, err := io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return file{local, n, err}
}
