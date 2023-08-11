// Exercise 4.12: The popular web comic xkcd has a JSON interface. For example, a request to https://xkcd.com/571/info.0.json produces a detalied description of comic 571, one of many favorites. Download each URL (once!) and build an offline index. Write a tool xkcd that, using this index, prints the URL and transcript of each comic that matches a search term provieded on the command line.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
    XKCD = "https://xkcd.com/"
    API = "/info.0.json"
    usage = `go run main index_number`
)

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintln(os.Stderr, usage)
        os.Exit(1)
    }
    s := XKCD + os.Args[1] + API
    resp, err := http.Get(s)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    c := make(map[string]string)
    json.Unmarshal(body, &c) // ignore err
    story := c["transcript"]
    story = clean(story)
    fmt.Println(s)
    fmt.Println(story)
}

func clean(story string) string {
    story = strings.Replace(story, "[[", "", -1)
    story = strings.Replace(story, "]]", "", -1)
    story = strings.Replace(story, "{{", "", -1)
    story = strings.Replace(story, "}}", "", -1)
    return story
}
