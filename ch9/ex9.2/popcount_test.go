package popcount_test

import (
    "testing"

    "the-go-programming-language/ch2/examples/popcount"
)

func BenchmarkPopCount(b *testing.B) {
    for i := 0; i < b.N; i++ {
        popcount.PopCount(0x123456789ABCDEF)
    }
}

// goos: darwin
// goarch: arm64
// pkg: the-go-programming-language/ch9/ex9.2
// BenchmarkPopCount-8   	1000000000	         0.3181 ns/op
// PASS
// ok  	the-go-programming-language/ch9/ex9.2	0.729s
