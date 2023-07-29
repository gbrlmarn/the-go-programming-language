// Exercise 3.9: Write a web server that renders fractals and writes the image data to the client. Allow the client to specify the x, y and zoom values as parameters to the HTTP request.
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func handler(w http.ResponseWriter, r *http.Request) {
    x, err := strconv.ParseFloat(r.URL.Query().Get("x"), 64)
    if err != nil {
        x = 0
    }
    y, err := strconv.ParseFloat(r.URL.Query().Get("y"), 64)
    if err != nil {
        y = 0
    }
    zoom, err := strconv.ParseFloat(r.URL.Query().Get("zoom"), 64)
    if err != nil {
        zoom = 1
    }
    fractal(w, x, y, zoom)
}
func fractal(w io.Writer, x, y, zoom float64) {
    var xmin, ymin, xmax, ymax = x - (2/zoom),
    y - (2/zoom), x + (2/zoom), y + (2/zoom)
    const (
        width, height          = 1024, 1024
    )

    img := image.NewRGBA(image.Rect(0, 0, width, height))
    for py := 0; py < height; py++ {
        y := float64(py)/height*(ymax-ymin) + ymin
        for px := 0; px < width; px++ {
            x := float64(px)/width*(xmax-xmin) + xmin
            z := complex(x, y)
            // Image point (px, py) represents complex value z.
            img.Set(px, py, mandlebrot(z))
        }
    }
    png.Encode(w, img) // NOTE: ignoring errors
}

func mandlebrot(z complex128) color.Color {
    const iterations = 200
    const contrast = 15

    var v complex128
    for n := uint8(0); n < iterations; n++ {
        v = v*v + z
        if cmplx.Abs(v) > 2 {
            return color.Gray{255 - contrast*n}
        }
    }
    return color.Black
}
// NOTE: https://localhost:8000/?x=<val>&y=<val>&zoom=<val>
