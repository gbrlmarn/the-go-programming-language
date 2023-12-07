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

// go test -bench=. the-go-programming-language/ch2/examples/popcount
// cpu: Intel(R) Core(TM) i5-7300U CPU @ 2.60GHz
// BenchmarkPopCount-4   	1000000000	         0.2878 ns/op

// goos: darwin
// goarch: arm64
// pkg: the-go-programming-language/ch2/examples/popcount
// BenchmarkPopCount-8   	1000000000	         0.3178 ns/op
// PASS
// ok  	the-go-programming-language/ch2/examples/popcount	0.645s
