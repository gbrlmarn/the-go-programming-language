// Exercise 3.7: Another simple fractal uses Newton's method to find complex solutions to a function such as z^4-1 = 0. Shade each starting point by the number of iterations required to get close to one of the four roots. color each point by the root it approaches.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
    var f func(complex128) color.Color
    if len(os.Args) == 2 {
        if os.Args[1] == "newton1" {
            f = newton1
        } else if os.Args[1] == "newton2" {
            f = newton2
        }
    } else {
        f = newton1
    }
    const (
        xmin, ymin, xmax, ymax = -2, -2, +2, +2
        width, height = 1024, 1024
    )

    img := image.NewRGBA(image.Rect(0, 0, width, height))
    for py := 0; py < height; py++ {
        y := float64(py)/height * (ymax - ymin) + ymin
        for px := 0; px < width; px++ {
            x := float64(px)/width * (xmax - xmin) + xmin
            z := complex(x, y)
            img.Set(px, py, f(z))
        }
    }
    png.Encode(os.Stdout, img)
}
func newton1(z complex128) color.Color {
    const iters = 200 
    const contrast = 15
    var r, g, b uint8
    for i := uint8(0); i < iters; i++ {
        r = 255 - contrast*i
        g = 255 - r
        b = 255 - i
        z -= (z - 1/(z*z*z))/4
        if cmplx.Abs(z*z*z*z - 1) < 1e-6 {
            return color.RGBA{r, g, b, 255}
        }
    }
    return color.Black
}
// NOTE: This is extra...
func newton2(z complex128) color.Color {
    const iters = 200 
    const contrast = 15
    var r, g, b uint8
    for i := uint8(0); i < iters; i++ {
        r = 255 - contrast*i
        g = 255 - r
        b = 255 - i
        z = root(z)
        if cmplx.Abs(f(z)) < 1e-6 {
            return color.RGBA{r, g, b, 255}
        }
    }
    return color.Black
}
func f(z complex128) complex128 {
    f := cmplx.Pow(z,8+0i) + 15*(cmplx.Pow(z,4+0i)) - 16
    return f
}
func root(z complex128) complex128 {
    fz := (cmplx.Pow(z, (8 + 0i)) + (15*(cmplx.Pow(z, (4 + 0i))) - 16))
    fdz := (8*(cmplx.Pow(z, (7 + 0i)))) + (60*(cmplx.Pow(z, (3+0i))))
    z -= fz / fdz
    return z
}

