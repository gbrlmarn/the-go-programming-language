package popcount_test

import (
    "testing"

    "gopl/ch2/ex2.3"
)

func BenchmarkPopCount(b *testing.B) {
    for i := 0; i < b.N; i++ {
        popcount.PopCount(0x123456789ABCDEF)
    }
}

// cpu: Intel(R) Core(TM) i5-7300U CPU @ 2.60GHz
// BenchmarkPopCount-4   	258538374	         4.632 ns/op
