package popcount_test

import (
    "testing"

    "gopl/ch2/examples/popcount"
)

func BenchmarkPopCount(b *testing.B) {
    for i := 0; i < b.N; i++ {
        popcount.PopCount(0x123456789ABCDEF)
    }
}

// goos: darwin
// goarch: arm64
// pkg: gopl/ch9/ex9.2
// BenchmarkPopCount-8   	1000000000	         0.3181 ns/op
// PASS
// ok  	gopl/ch9/ex9.2	0.729s
