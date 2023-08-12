// Exercise 4.13: The JSON-based web service of the Open Movie Database les you search https://omdbapi.com/ for a movie by name and download its poster image. Write a tool poster that downloads the poster for the movie named on the command line.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
)

// found this api here: 'https://codepen.io/pixelnik/pen/pgWQBZ'
// didn't manage to makke 'https://omdbapi.com/' to work
const (
    MovieAPI = "https://api.themoviedb.org/3/search/movie?api_key=15d2ea6d0dc1d476efbca3eba2b9bbfb&query="
    PosterAPI = "http://image.tmdb.org/t/p/w500/"
    usage = `go run main.go <movieTitle>`
)

type Movies struct {
    Results []Movie `json:"results"`
}
type Movie struct {
    Title  string `json:"original_title"`
    Date   string `json:"release_date"`
    Poster string `json:"poster_path"`
}

func main() {
    if len(os.Args) < 2 {
        log.Fatal(usage)
    }
	movies := getResults(os.Args[1])
    if len(movies.Results) > 0 {
        resp, err := http.Get(PosterAPI+movies.Results[0].Poster)
        if err != nil {
            log.Fatal(err)
        }
        defer resp.Body.Close()
        body, err := io.ReadAll(resp.Body)
        if err != nil {
            log.Fatal(err)
        }
        decodeToJpeg(body, movies.Results[0].Title)
    } else {
        fmt.Println("Nothing found...")
        return
    }
}

func getResults(title string) Movies {
	resp, err := http.Get(MovieAPI + title)
	if err != nil {
        log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
        log.Fatal(err)
	}
    var movies Movies 
    err = json.Unmarshal(body, &movies)
    if err != nil {
        log.Fatal(err)
    }
    return movies 
}

func decodeToJpeg(imgByte []byte, imgName string) {
    img, _, err := image.Decode(bytes.NewReader(imgByte))
    if err != nil {
        log.Fatalln(err)
    }
    out, _ := os.Create(imgName+".jpeg")
    defer out.Close()

    var opts jpeg.Options
    opts.Quality = 100 // best quality, just for you
    err = jpeg.Encode(out, img, &opts)
    if err != nil {
        log.Println(err)
    }
}
