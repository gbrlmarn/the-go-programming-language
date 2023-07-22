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
