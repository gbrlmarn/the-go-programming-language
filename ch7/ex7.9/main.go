// Exercise 7.9: Use the html/template package (§4.6) to replace printTracks with a function that displays the tracks as an HTML table. Use the solution to the previous exercise to arrange that each click on a column head makes an HTTP request to sort the table.
package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type customSort struct {
	t []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }

func by(fld string) func(x, y *Track) bool {
	switch fld {
	case "Title":
		return func(x, y *Track) bool {
			return x.Title < y.Title
		}
	case "Artist":
		return func(x, y *Track) bool {
			return x.Artist < y.Artist
		}
	case "Album":
		return func(x, y *Track) bool {
			return x.Album < y.Album
		}
	case "Year":
		return func(x, y *Track) bool {
			return x.Year < y.Year
		}
	case "Length":
		return func(x, y *Track) bool {
			return x.Length < y.Length
		}
	default:
		return func(x, y *Track) bool {
			return false
		}
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/Title":
		sort.Sort(customSort{tracks, by("Title")})
	case "/Artist":
		sort.Sort(customSort{tracks, by("Artist")})
	case "/Album":
		sort.Sort(customSort{tracks, by("Album")})
	case "/Year":
		sort.Sort(customSort{tracks, by("Year")})
	case "/Length":
		sort.Sort(customSort{tracks, by("Length")})
	}
	
	tmpl := template.Must(template.ParseFiles("index.html"))
	if err := tmpl.Execute(w, &tracks); err != nil {
		log.Println(err)
	}
}

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

