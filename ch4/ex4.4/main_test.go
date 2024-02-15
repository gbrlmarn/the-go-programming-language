package main

import (
	"testing"
)

func BenchmarkRotate(b *testing.B) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i := 0; i < b.N; i++ {
		rotate(a, i%len(a))
	}
}
func BenchmarkRotate2(b *testing.B) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i := 0; i < b.N; i++ {
		rotate2(a, i%len(a))
	}
}
