// Exercise 1.8: Modify fetch to add the prefix http:// to each argument URL ifit is missing. You might want to use strings.HasPrefix.
package main

import (
    "fmt"
    "io"
    "os"
    "strings"
    "net/http"
)

func main() {
    for _, url := range os.Args[1:] {
        if !strings.HasPrefix(url, "http://") {
        url = "http://" + url
        }
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }
        _, err = io.Copy(os.Stdout, resp.Body)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: copy %s %v\n", url, err)
            os.Exit(1)
        }
    }
}
