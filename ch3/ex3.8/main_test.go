package main

import (
	"image/color"
	"testing"
)

func benchmarkMandelbrot(b *testing.B, f func(complex128) color.Color) {
    for i := 0; i < b.N; i++ {
        f(complex(float64(i), float64(i)))
    }
}
func BenchmarkMandelbrot64(b *testing.B) {
    benchmarkMandelbrot(b, mandlebrot64)
}
func BenchmarkMandelbrot128(b *testing.B) {
    benchmarkMandelbrot(b, mandlebrot128)
}
func BenchmarkMandelbrotbigF(b *testing.B) {
    benchmarkMandelbrot(b, mandlebrotbigF)
}
func BenchmarkMandelbrotR(b *testing.B) {
    benchmarkMandelbrot(b, mandelbrotbigR)
}
