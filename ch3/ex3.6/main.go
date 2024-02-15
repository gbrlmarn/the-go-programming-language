// Exercise 3.6: Supersampling is a technique to reduce the effect of pixelation by computing the color value at several points withing each pixel and taking the average. The simplest method is to divide each pixel into four "subpixels." Implement it.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			spclr := make([]color.Color, 0)
			for spy := 0; spy < 2; spy++ {
				for spx := 0; spx < 2; spx++ {
					y := float64(py+spy)/height*(ymax-ymin) + ymin
					x := float64(px+spx)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					spclr = append(spclr, mandelbrot(z))
				}
			}
			// Image point (px, py) represents complex value z.
			img.Set(px, py, avg(spclr))
		}
	}
	png.Encode(os.Stdout, img)
}
func avg(spclr []color.Color) color.Color {
	var r, g, b, a uint8
	n := len(spclr)
	for _, v := range spclr {
		_r, _g, _b, _a := v.RGBA()
		r += uint8(_r / uint32(n))
		g += uint8(_g / uint32(n))
		b += uint8(_b / uint32(n))
		a += uint8(_a / uint32(n))
	}
	return color.RGBA{r, g, b, a}
}
func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for i := uint8(0); i < iterations; i++ {
		v = v*v + z
		r := uint8(255 - contrast*i)
		g := uint8(255 - r)
		b := uint8(255 - i)
		if cmplx.Abs(v) > 2 {
			return color.RGBA{r, g, b, 255}
		}
	}
	return color.Black
}
