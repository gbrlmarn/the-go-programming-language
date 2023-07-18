// Exercise 1.9: Modify fetch to also print the HTTP status code, found in resp.Status.
package main

import (
    "os"
    "io"
    "fmt"
    "net/http"
)

func main() {
    for _, url := range os.Args[1:] {
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }
        _, err = io.Copy(os.Stdout, resp.Body)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: copy %s %v\n", url, err)
        }
        fmt.Printf("\n%s\n", resp.Status)
    }
}
