package popcount_test

import (
    "testing"

    "gopl/ch2/ex2.5"
)

func BenchmarkPopCount(b *testing.B) {
    for i := 0; i < b.N; i++ {
        popcount.PopCount(0x123456789ABCDEF)
    }
}
