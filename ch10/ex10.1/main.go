// Exercise 10.1: Extend the jpeg program so that it converts any supported input format to any output format, using image.Decode to detect the input format and a flat to select the output format.
package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

func main() {
	format := flag.String("format", "jpeg", "Image encoding format")
	flag.Parse()
	if err := toFormat(*format, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toFormat(format string, in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format = ", kind)
	switch format {
	case "jpeg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png":
		return png.Encode(out, img)
	case "gif":
		return gif.Encode(out, img, &gif.Options{NumColors: 256})
	default:
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	}
}

// 116K Dec 20 17:54 mandelbrot.gif
// 95K  Dec 20 17:54 mandelbrot.jpeg
// 88K  Dec 20 17:54 mandelbrot.png
