// Exercise 3.8: Rendering fractals at high zoom levels demands great arithmetic precision. Implement the same fractal using four different reresentations of numbers: complex64, complex128, big.Float, and big.Rat. (The latter two types are found in the math/big package. Float uses arbitrary bu tbounded-precision floating-point; Rat uses unbbounded-precision rational numbers.) How do they compare in performance and memory usage? At what zoom levels do rendering artifacts become visible?
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/big"
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
        y := float64(py)/height*(ymax-ymin) + ymin
        for px := 0; px < width; px++ {
            x := float64(px)/width*(xmax-xmin) + xmin
            z := complex(x, y)
            // Image point (px, py) represents complex value z.
            img.Set(px, py, mandlebrot128(z))
        }
    }
    png.Encode(os.Stdout, img) // NOTE: ignoring errors
}
func mandlebrot64(z complex128) color.Color {
    const iterations = 200
    const contrast = 15

    var v complex64 
    for n := uint8(0); n < iterations; n++ {
        v = v*v + complex64(z)
        if cmplx.Abs(complex128(v)) > 2 {
            return color.Gray{255 - contrast*n}
        }
    }
    return color.Black
}
func mandlebrot128(z complex128) color.Color {
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
func mandlebrotbigF(z complex128) color.Color {
    const iterations = 200
    const contrast = 15
    zR := (&big.Float{}).SetFloat64(real(z))
    zI := (&big.Float{}).SetFloat64(imag(z))
    vR, vI := &big.Float{}, &big.Float{}
    for n := uint8(0); n < iterations; n++ {
        // (r+i)^2 = r^2  + 2ri + i^2
        vR2, vI2 := &big.Float{}, &big.Float{}
        vR2.Mul(vR, vR).Sub(vR2, (&big.Float{}).Mul(vI, vI)).Add(vR2, zR)
        vI2.Mul(vR, vI).Mul(vI2, big.NewFloat(2)).Add(vI2, zI)
        vR, vI = vR2, vI2
        sqSum := &big.Float{}
        sqSum.Mul(vR, vR).Add(sqSum, (&big.Float{}).Mul(vI, vI))
        if sqSum.Cmp(big.NewFloat(4)) == 1 {
            return color.Gray{255 - contrast*n}
        }
    }
    return color.Black
}
func mandelbrotbigR(z complex128) color.Color {
    const iterations = 100
    const contrast = 15
    zR := (&big.Rat{}).SetFloat64(real(z))
    zI := (&big.Rat{}).SetFloat64(imag(z))
    vR, vI := &big.Rat{}, &big.Rat{}
    for n := uint8(0); n < iterations; n++ {
        // (r+i)^2 = r^2 + 2ri + i^2
        vR2, vI2 := &big.Rat{}, &big.Rat{}
        vR2.Mul(vR, vR).Sub(vR2, (&big.Rat{}).Mul(vI, vI)).Add(vR2, zR)
        vI2.Mul(vR, vI).Mul(vI2, big.NewRat(2, 1)).Add(vI2, zI)
        vR, vI = vR2, vI2
        sqSum := &big.Rat{}
        sqSum.Mul(vR, vR).Add(sqSum, (&big.Rat{}).Mul(vI, vI))
        if sqSum.Cmp(big.NewRat(4, 1)) == 1 {
            return color.Gray{255 - contrast * n}
        } 
    }
    return color.Black
}
// goos: linux
// goarch: amd64
// pkg: the-go-programming-language/ch3/ex3.8
// cpu: Intel(R) Core(TM) i5-7300U CPU @ 2.60GHz
// BenchmarkMandelbrot64-4     	86675203	        12.05 ns/op
// BenchmarkMandelbrot128-4    	136737777	         9.005 ns/op
// BenchmarkMandelbrotbigF-4   	 2609875	       443.1 ns/op
// BenchmarkMandelbrotR-4      	  665094	      1698 ns/op
// PASS
// ok  	the-go-programming-language/ch3/ex3.8	5.955s
